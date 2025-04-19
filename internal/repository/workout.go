package repository

import (
	"traning/models"

	"gorm.io/gorm"
)

type WorkoutRepository struct {
	db *gorm.DB
}

func NewWorkoutRepository(db *gorm.DB) *WorkoutRepository {
	return &WorkoutRepository{db}
}

func (r *WorkoutRepository) GetAllWorkouts() ([]models.Workout, error) {
	var workouts []models.Workout
	err := r.db.Preload("Exercises").Order("created_at desc").Find(&workouts).Error
	return workouts, err
}

func (r *WorkoutRepository) GetWorkoutWithExercises(id uint) (*models.Workout, error) {
	var workout models.Workout
	err := r.db.Preload("Exercises").First(&workout, id).Error
	if err != nil {
		return nil, err
	}
	return &workout, nil
}

func (r *WorkoutRepository) CreateWorkout(name string, exerciseIDs []uint) (*models.Workout, error) {
	exercises := make([]models.Exercise, len(exerciseIDs))
	for i, id := range exerciseIDs {
		exercises[i] = models.Exercise{ID: id}
	}

	workout := models.Workout{
		Name:      name,
		Exercises: exercises,
	}

	if err := r.db.Create(&workout).Error; err != nil {
		return nil, err
	}
	return &workout, nil
}

func (r *WorkoutRepository) UpdateWorkout(id uint, name string, exerciseIDs []uint) (*models.Workout, error) {
	var workout models.Workout
	if err := r.db.First(&workout, id).Error; err != nil {
		return nil, err
	}

	exercises := make([]models.Exercise, len(exerciseIDs))
	for i, id := range exerciseIDs {
		exercises[i] = models.Exercise{ID: id}
	}

	workout.Name = name
	if err := r.db.Model(&workout).Association("Exercises").Replace(exercises); err != nil {
		return nil, err
	}

	err := r.db.Save(&workout).Error
	return &workout, err
}

func (r *WorkoutRepository) DeleteWorkout(id uint) error {
	return r.db.Delete(&models.Workout{}, id).Error
}
