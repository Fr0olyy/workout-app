package services

import (
	"errors"
	"math"
	"traning/internal/repository"
	"traning/models"
)

type WorkoutService struct {
	repo repository.Workout
}

func NewWorkoutService(repo repository.Workout) *WorkoutService {
	return &WorkoutService{repo: repo}
}

func (s *WorkoutService) GetAllWorkouts() ([]models.Workout, error) {
	return s.repo.GetAllWorkouts()
}

func (s *WorkoutService) GetWorkout(id uint) (*models.Workout, int, error) {
	w, err := s.repo.GetWorkoutWithExercises(id)
	if err != nil {
		return nil, 0, errors.New("Workout not found")
	}
	minutes := calculateMinute(len(w.Exercises))
	return w, minutes, nil
}

func (s *WorkoutService) CreateWorkout(name string, exerciseIDs []uint) (*models.Workout, error) {
	return s.repo.CreateWorkout(name, exerciseIDs)
}

func (s *WorkoutService) UpdateWorkout(id uint, name string, exerciseIDs []uint) (*models.Workout, error) {
	return s.repo.UpdateWorkout(id, name, exerciseIDs)
}

func (s *WorkoutService) DeleteWorkout(id uint) error {
	return s.repo.DeleteWorkout(id)
}

func calculateMinute(length int) int {
	return int(math.Ceil(float64(length) * 3.7))
}
