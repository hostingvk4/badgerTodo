package repository

import (
	"github.com/hostingvk4/badgerList/internal/models"
	"gorm.io/gorm"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
	result := r.db.Create(&user)

	return int(user.ID), result.Error
}

func (r *AuthPostgres) GetUser(username, password string) (models.User, error) {
	var userModel models.User
	result := r.db.Where("username = ? AND password = ?", username, password).Find(&userModel)

	return userModel, result.Error
}
