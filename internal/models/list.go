package models

import "gorm.io/gorm"

type List struct {
	gorm.Model
	Title       string
	UserId      uint
	Description string
}
type ListDto struct {
	ID          uint   `json:"id,string,omitempty"`
	Title       string `json:"title" binding:"required"`
	UserId      uint   `json:"user_id"`
	Description string `json:"description" binding:"required"`
}
