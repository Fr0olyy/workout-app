package repository

import (
	"traning/models"

	"gorm.io/gorm"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) (uint, error) {
	u := models.User{
		Email:    user.Email,
		Password: user.Password,
	}

	if err := r.db.Create(&u).Find(&u).Error; err != nil {
		return 0, err
	}

	return u.ID, nil
}

func (r *AuthPostgres) GetUser(email, password string) (models.User, error) {
	var user models.User
	result := r.db.Where("email = ? AND password = ?", email, password).First(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}
