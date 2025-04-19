package repository

import (
	"traning/models"

	"gorm.io/gorm"
)

type ExercisePostgres struct {
	db *gorm.DB
}

func NewExercisePostgres(db *gorm.DB) *ExercisePostgres {
	return &ExercisePostgres{db: db}
}

// Exercise CRUD
func (r *ExercisePostgres) CreateExercise(exercise *models.Exercise) error {
	return r.db.Create(exercise).Error
}

func (r *ExercisePostgres) UpdateExercise(exercise *models.Exercise) error {
	return r.db.Save(exercise).Error
}

func (r *ExercisePostgres) DeleteExercise(id uint) error {
	return r.db.Delete(&models.Exercise{}, id).Error
}

func (r *ExercisePostgres) GetExercises() ([]models.Exercise, error) {
	var exercises []models.Exercise
	err := r.db.Order("created_at desc").Find(&exercises).Error
	return exercises, err
}

// Exercise Logs
func (r *ExercisePostgres) CreateExerciseLog(log *models.ExerciseLog) error {
	return r.db.Create(log).Error
}

func (r *ExercisePostgres) GetExerciseLogWithTimes(id uint) (*models.ExerciseLog, error) {
	var log models.ExerciseLog
	err := r.db.Preload("Times").Preload("Exercise").First(&log, id).Error
	return &log, err
}

func (r *ExercisePostgres) UpdateExerciseTime(time *models.ExerciseTime) error {
	return r.db.Save(time).Error
}

func (r *ExercisePostgres) CompleteExerciseLog(log *models.ExerciseLog) error {
	return r.db.Save(log).Error
}

func (r *ExercisePostgres) GetPreviousLog(exerciseID, userID uint) (*models.ExerciseLog, error) {
	var log models.ExerciseLog
	err := r.db.Where("exercise_id = ? AND user_id = ? AND is_completed = ?",
		exerciseID, userID, true).
		Order("created_at desc").
		Preload("Times").
		First(&log).Error
	return &log, err
}

func (r *ExercisePostgres) GetExerciseByID(id uint) (*models.Exercise, error) {
	var exercise models.Exercise
	if err := r.db.First(&exercise, id).Error; err != nil {
		return nil, err
	}
	return &exercise, nil
}

func (r *ExercisePostgres) CreateExerciseTimes(times []models.ExerciseTime) error {
	return r.db.Create(&times).Error
}

func (r *ExercisePostgres) GetExerciseLogWithDetails(id uint, userID uint) (*models.ExerciseLog, error) {
	var log models.ExerciseLog
	err := r.db.Preload("Exercise").
		Preload("Times", func(db *gorm.DB) *gorm.DB {
			return db.Order("id ASC")
		}).
		Where("id = ? AND user_id = ?", id, userID).
		First(&log).Error

	return &log, err
}
