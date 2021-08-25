package models

import "gorm.io/gorm"

type List struct {
	gorm.Model
	Title       string `json:"title" binding:"required"`
	UserId      int    `json:"user_id"`
	Description string `json:"description"`
}
type UpdateListInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
