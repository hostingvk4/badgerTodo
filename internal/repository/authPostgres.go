package repository

import (
	"github.com/hostingvk4/badgerList/internal/models"
	"gorm.io/gorm"
	"time"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) (uint, error) {
	result := r.db.Create(&user)

	return uint(user.ID), result.Error
}

func (r *AuthPostgres) GetUser(username, password string) (models.User, error) {
	var userModel models.User
	err := r.db.Where("username = ? AND password = ?", username, password).First(&userModel).Error

	return userModel, err
}

func (r *AuthPostgres) SetRefreshToken(token models.RefreshToken) error {
	result := r.db.Create(&token)

	return result.Error
}

func (r *AuthPostgres) UpdateRefreshToken(oldRefreshToken string, refreshToken models.RefreshToken) error {
	var tokenModel models.RefreshToken
	err := r.db.Where("user_id = ? AND refresh_token = ? AND expires_at > ?", refreshToken.UserId, oldRefreshToken, time.Now()).First(&tokenModel).Error
	if err == nil {
		tokenModel.RefreshToken = refreshToken.RefreshToken
		tokenModel.ExpiresAt = refreshToken.ExpiresAt
		err = r.db.Model(&tokenModel).Updates(models.RefreshToken{RefreshToken: refreshToken.RefreshToken, ExpiresAt: refreshToken.ExpiresAt}).Error
	}
	return err
}
