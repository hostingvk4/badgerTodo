package repository

import (
	"github.com/hostingvk4/badgerList/internal/models"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user models.User) (uint, error)
	GetUser(username, password string) (models.User, error)
}

type List interface {
	Create(list models.List) (uint, error)
	GetAll(userId uint) ([]models.List, error)
	GetListById(userId, listId uint) (models.List, error)
	Update(userId, listId uint, list models.List) error
	Delete(userId, listId uint) error
}

type Repository struct {
	Authorization
	List
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		List:          NewListPostgres(db),
	}
}
