package db

import (
	"fmt"
	"log"
	"os"
	"traning/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBname   string
	SSLMode  string
}

func InitDB() (*gorm.DB, error) {

	var err error

	cfg := Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		DBname:   os.Getenv("DB_DBNAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	if cfg.Host == "" || cfg.Port == "" || cfg.Username == "" || cfg.DBname == "" || cfg.Password == "" {
		log.Fatal("Missing required database configuration")
	}
	log.Printf("DB config: host=%s port=%s db=%s user=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.DBname, cfg.Username, cfg.SSLMode)

	log.Println("Starting DB migration...")

	// Если SSLMode не задан, установите значение по умолчанию
	if cfg.SSLMode == "" {
		cfg.SSLMode = "disable" // или "require" для продакшена
	}

	db, err = gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Host, cfg.Username, cfg.Password, cfg.DBname, cfg.Port, cfg.SSLMode)), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	if err := db.AutoMigrate(
		&models.User{},
		&models.Workout{},
		&models.Exercise{},
		&models.ExerciseLog{},
		&models.ExerciseTime{},
		&models.WorkoutLog{},
	); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	return db, nil
}
