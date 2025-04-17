package handlers

import (
	"net/http"
	"traning/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Statistic struct {
	Label string  `json:"label"`
	Value float64 `json:"value"`
}

type UserProfileResponse struct {
	User       models.User `json:"user"`
	Statistics []Statistic `json:"statistics"`
}

func (h *Handler) GetUserProfile(c echo.Context) error {
	userID := c.Get("user_id").(uint)
	var db *gorm.DB

	// Получаем пользователя из базы данных
	var user models.User
	if err := db.Preload("ExerciseLogs").Preload("WorkoutLogs").First(&user, userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "User not found")
		}
		return err
	}

	// Подсчет выполненных упражнений (exercise logs)
	var countExerciseTimesCompleted int64
	db.Model(&models.ExerciseLog{}).Where("user_id = ? AND is_completed = ?", userID, true).Count(&countExerciseTimesCompleted)

	// Подсчет суммарного веса (exercise times)
	var totalWeight float64
	db.Model(&models.ExerciseTime{}).
		Joins("JOIN exercise_logs ON exercise_logs.id = exercise_times.exercise_log_id").
		Where("exercise_logs.user_id = ? AND exercise_times.is_completed = ?", userID, true).
		Select("SUM(exercise_times.weight)").Scan(&totalWeight)

	// Подсчет количества завершенных тренировок
	var workouts int64
	db.Model(&models.WorkoutLog{}).Where("user_id = ? AND is_completed = ?", userID, true).Count(&workouts)

	// Формирование статистики
	statistics := []Statistic{
		{
			Label: "Minutes",
			Value: float64(countExerciseTimesCompleted) * 2.3,
		},
		{
			Label: "Workouts",
			Value: float64(workouts),
		},
		{
			Label: "Kgs",
			Value: totalWeight,
		},
	}

	// Возвращаем результат в формате JSON
	return c.JSON(http.StatusOK, UserProfileResponse{
		User:       user,
		Statistics: statistics,
	})
}
