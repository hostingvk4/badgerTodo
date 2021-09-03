package service

import (
	"github.com/hostingvk4/badgerList/internal/models"
	"github.com/hostingvk4/badgerList/internal/repository"
)

type Authorization interface {
	CreateUser(user models.User) (uint, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (uint, error)
}

type List interface {
	Create(list models.List) (uint, error)
	GetAll(userId uint) ([]models.List, error)
	GetListById(userId, listId uint) (models.List, error)
	Update(userId, id uint, list models.List) error
	Delete(userId, listId uint) error
}

type Service struct {
	Authorization
	List
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		List:          NewListService(repos.List),
	}
}
