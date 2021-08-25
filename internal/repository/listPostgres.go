package repository

import (
	"errors"
	"github.com/hostingvk4/badgerList/internal/models"
	"gorm.io/gorm"
)

type ListPostgres struct {
	db *gorm.DB
}

func NewListPostgres(db *gorm.DB) *ListPostgres {
	return &ListPostgres{db: db}
}

func (r *ListPostgres) Create(list models.List) (int, error) {
	result := r.db.Create(&list)

	return int(list.ID), result.Error
}

func (r *ListPostgres) GetAll(userId int) ([]models.List, error) {
	var lists []models.List
	result := r.db.Where("user_id = ?", userId).Find(&lists)

	return lists, result.Error
}
func (r *ListPostgres) GetListById(userId, listId int) (models.List, error) {
	var lists models.List
	result := r.db.Where("id = ? AND user_id = ?", listId, userId).Find(&lists)

	return lists, result.Error
}
func (r *ListPostgres) Update(userId, listId int, list models.UpdateListInput) error {
	var lists models.List
	r.db.First(&lists, "id = ? and user_id = ?", listId, userId)
	result := r.db.Model(&lists).Updates(models.List{Title: list.Title, Description: list.Description})
	if result.RowsAffected == 0 {
		return errors.New("update error")
	}
	return result.Error
}
func (r *ListPostgres) Delete(userId, listId int) error {
	var list models.List
	result := r.db.Where("id = ? AND user_id = ?", listId, userId).Delete(&list)
	if result.RowsAffected == 0 {
		return errors.New("delete error")
	}
	return result.Error
}
