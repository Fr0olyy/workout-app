package repository

import (
	"traning/models"

	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user models.User) (uint, error)
	GetUser(email, password string) (models.User, error)
}

type Workout interface {
	GetAllWorkouts() ([]models.Workout, error)
	GetWorkoutWithExercises(id uint) (*models.Workout, error)
	CreateWorkout(name string, exerciseIDs []uint) (*models.Workout, error)
	UpdateWorkout(id uint, name string, exerciseIDs []uint) (*models.Workout, error)
	DeleteWorkout(id uint) error
}

type Exercise interface {
	GetExerciseByID(id uint) (*models.Exercise, error)
	CreateExercise(exercise *models.Exercise) error
	UpdateExercise(exercise *models.Exercise) error
	DeleteExercise(id uint) error
	GetExercises() ([]models.Exercise, error)
	CreateExerciseLog(log *models.ExerciseLog) error
	GetExerciseLogWithTimes(id uint) (*models.ExerciseLog, error)
	UpdateExerciseTime(time *models.ExerciseTime) error
	CompleteExerciseLog(log *models.ExerciseLog) error
	GetPreviousLog(exerciseID, userID uint) (*models.ExerciseLog, error)
	CreateExerciseTimes(times []models.ExerciseTime) error
	GetExerciseLogWithDetails(id uint, userID uint) (*models.ExerciseLog, error)
}

type Profile interface {
	GetUserByID(userID uint) (*models.User, error)
	CountCompletedExerciseLogs(userID uint) (int64, error)
	SumCompletedWeights(userID uint) (int64, error)
	CountCompletedWorkouts(userID uint) (int64, error)
}

type Repository struct {
	Authorization
	Workout
	Exercise
	Profile
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Workout:       NewWorkoutRepository(db),
		Exercise:      NewExercisePostgres(db),
		Profile:       NewUserRepository(db),
	}
}
