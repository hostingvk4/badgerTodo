package service

import (
	"github.com/hostingvk4/badgerList/internal/models"
	"github.com/hostingvk4/badgerList/internal/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type List interface {
	Create(list models.List) (int, error)
	GetAll(userId int) ([]models.List, error)
	GetListById(userId, listId int) (models.List, error)
	Update(userId, id int, list models.UpdateListInput) error
	Delete(userId, listId int) error
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
