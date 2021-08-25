package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required" gorm:"size:255;unique"`
	Password string `json:"password" binding:"required"`
	List     []List
}
