package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string   `gorm:"uniqueIndex;not null"`
	Name     string   `gorm:"not null"`
	Password string   `gorm:"not null"`
	Images   []string `gorm:"type:text[]"`

	ExerciseLogs []ExerciseLog `gorm:"foreignKey:UserID"`
	WorkoutLogs  []WorkoutLog  `gorm:"foreignKey:UserID"`
}

type Workout struct {
	ID        uint           `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name        string
	Exercises   []Exercise   `gorm:"many2many:workout_exercises;"`
	WorkoutLogs []WorkoutLog `gorm:"foreignKey:WorkoutID"`
}

type Exercise struct {
	ID        uint           `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name     string `gorm:"not null"`
	Times    int    `gorm:"not null"`
	IconPath string `gorm:"column:icon_path"`

	Workouts     []Workout     `gorm:"many2many:workout_exercises;"`
	ExerciseLogs []ExerciseLog `gorm:"foreignKey:ExerciseID"`
}

type ExerciseLog struct {
	ID        uint           `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	IsCompleted bool           `gorm:"column:is_completed;default:false"`
	Times       []ExerciseTime `gorm:"foreignKey:ExerciseLogID"`

	UserID       uint `gorm:"column:user_id"`
	WorkoutLogID uint `gorm:"column:workout_log_id"`
	ExerciseID   uint `gorm:"column:exercise_id"`

	User       *User       `gorm:"foreignKey:UserID"`
	WorkoutLog *WorkoutLog `gorm:"foreignKey:WorkoutLogID"`
	Exercise   *Exercise   `gorm:"foreignKey:ExerciseID"`
}

type ExerciseTime struct {
	ID        uint           `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Weight      int  `gorm:"default:0"`
	Repeat      int  `gorm:"default:0"`
	IsCompleted bool `gorm:"column:is_completed;default:false"`

	ExerciseLogID uint         `gorm:"column:exercise_log_id"`
	ExerciseLog   *ExerciseLog `gorm:"foreignKey:ExerciseLogID"`
}

type WorkoutLog struct {
	ID        uint           `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	IsCompleted bool `gorm:"column:is_completed;default:false"`

	UserID    uint `gorm:"column:user_id"`
	WorkoutID uint `gorm:"column:workout_id"`

	User         *User         `gorm:"foreignKey:UserID"`
	Workout      *Workout      `gorm:"foreignKey:WorkoutID"`
	ExerciseLogs []ExerciseLog `gorm:"foreignKey:WorkoutLogID"`
}
