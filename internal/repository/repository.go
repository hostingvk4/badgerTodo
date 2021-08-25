package repository

import (
	"github.com/hostingvk4/badgerList/internal/models"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type List interface {
	Create(list models.List) (int, error)
	GetAll(userId int) ([]models.List, error)
	GetListById(userId, listId int) (models.List, error)
	Update(userId, listId int, list models.UpdateListInput) error
	Delete(userId, listId int) error
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
