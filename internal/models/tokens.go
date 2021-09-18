package models

import (
	"gorm.io/gorm"
	"time"
)

type RefreshToken struct {
	gorm.Model
	RefreshToken string
	UserId       uint
	ExpiresAt    time.Time
}
