package services

import (
	"errors"
	"traning/internal/repository"
	"traning/models"
)

type ExerciseService struct {
	repo repository.Exercise
}

func NewExerciseService(repo repository.Exercise) *ExerciseService {
	return &ExerciseService{repo: repo}
}

func (s *ExerciseService) CreateExercise(exercise *models.Exercise) error {
	if exercise.Name == "" {
		return errors.New("exercise name is required")
	}
	return s.repo.CreateExercise(exercise)
}

func (s *ExerciseService) UpdateExercise(exercise *models.Exercise) error {
	existing, err := s.repo.GetExerciseByID(exercise.ID)
	if err != nil {
		return errors.New("exercise not found")
	}
	existing.Name = exercise.Name
	existing.Times = exercise.Times
	existing.IconPath = exercise.IconPath
	return s.repo.UpdateExercise(existing)
}

func (s *ExerciseService) DeleteExercise(id uint) error {
	return s.repo.DeleteExercise(id)
}

func (s *ExerciseService) GetAllExercises() ([]models.Exercise, error) {
	return s.repo.GetExercises()
}

func (s *ExerciseService) CreateExerciseLog(exerciseID, userID uint) (*models.ExerciseLog, error) {
	exercise, err := s.repo.GetExerciseByID(exerciseID)
	if err != nil {
		return nil, errors.New("exercise not found")
	}

	log := &models.ExerciseLog{
		UserID:     userID,
		ExerciseID: exerciseID,
	}

	if err := s.repo.CreateExerciseLog(log); err != nil {
		return nil, err
	}

	times := make([]models.ExerciseTime, exercise.Times)
	for i := range times {
		times[i] = models.ExerciseTime{
			ExerciseLogID: log.ID,
			Weight:        0,
			Repeat:        0,
		}
	}

	if err := s.repo.CreateExerciseTimes(times); err != nil {
		return nil, err
	}

	return log, nil
}

func (s *ExerciseService) GetExerciseLog(logID uint, userID uint) (*models.ExerciseLog, []TimeWithPrev, error) {
	log, err := s.repo.GetExerciseLogWithDetails(logID, userID)
	if err != nil {
		return nil, nil, errors.New("exercise log not found")
	}

	prevLog, _ := s.repo.GetPreviousLog(log.ExerciseID, userID)

	timesWithPrev := addPrevValues(log, prevLog)

	return log, timesWithPrev, nil
}

type TimeWithPrev struct {
	models.ExerciseTime
	PrevWeight float64 `json:"prevWeight"`
	PrevRepeat int     `json:"prevRepeat"`
}

func addPrevValues(current *models.ExerciseLog, prev *models.ExerciseLog) []TimeWithPrev {
	result := make([]TimeWithPrev, len(current.Times))

	for i, time := range current.Times {
		twp := TimeWithPrev{ExerciseTime: time}

		if prev != nil && i < len(prev.Times) {
			twp.PrevWeight = float64(prev.Times[i].Weight)
			twp.PrevRepeat = prev.Times[i].Repeat
		}

		result[i] = twp
	}

	return result
}
