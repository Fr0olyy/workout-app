package services

import (
	"traning/internal/repository"
	"traning/models"
)

type Profile interface {
	GetUserProfile(userID uint) (*UserProfileResponse, error)
}

type Authorization interface {
	CreateUser(user models.User) (uint, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (uint, error)
}

type Workout interface {
	GetAllWorkouts() ([]models.Workout, error)
	GetWorkout(id uint) (*models.Workout, int, error)
	CreateWorkout(name string, exerciseIDs []uint) (*models.Workout, error)
	UpdateWorkout(id uint, name string, exerciseIDs []uint) (*models.Workout, error)
	DeleteWorkout(id uint) error
}

type Exercise interface {
	CreateExercise(exercise *models.Exercise) error
	UpdateExercise(exercise *models.Exercise) error
	DeleteExercise(id uint) error
	GetAllExercises() ([]models.Exercise, error)
	CreateExerciseLog(exerciseID, userID uint) (*models.ExerciseLog, error)
	GetExerciseLog(logID uint, userID uint) (*models.ExerciseLog, []TimeWithPrev, error)
}

type Service struct {
	Authorization
	Workout
	Exercise
	Profile
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(r.Authorization),
		Workout:       NewWorkoutService(r.Workout),
		Exercise:      NewExerciseService(r.Exercise),
		Profile:       NewUserService(r.Profile),
	}
}
