package repository

import (
	"traning/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) GetUserByID(userID uint) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) CountCompletedExerciseLogs(userID uint) (int64, error) {
	var count int64
	err := r.db.Model(&models.ExerciseLog{}).
		Where("user_id = ? AND is_completed = true", userID).
		Count(&count).Error
	return count, err
}

func (r *UserRepository) SumCompletedWeights(userID uint) (int64, error) {
	var result struct {
		TotalWeight *int64
	}
	err := r.db.Model(&models.ExerciseTime{}).
		Select("SUM(weight) as total_weight").
		Joins("JOIN exercise_logs ON exercise_logs.id = exercise_times.exercise_log_id").
		Where("exercise_logs.user_id = ? AND exercise_times.is_completed = true", userID).
		Scan(&result).Error

	if result.TotalWeight == nil {
		return 0, err
	}
	return *result.TotalWeight, err
}

func (r *UserRepository) CountCompletedWorkouts(userID uint) (int64, error) {
	var count int64
	err := r.db.Model(&models.WorkoutLog{}).
		Where("user_id = ? AND is_completed = true", userID).
		Count(&count).Error
	return count, err
}
