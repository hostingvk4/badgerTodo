package service

import (
	"github.com/hostingvk4/badgerList/internal/models"
	"github.com/hostingvk4/badgerList/internal/repository"
	"github.com/hostingvk4/badgerList/pkg/auth"
	"github.com/hostingvk4/badgerList/pkg/cipher"
	"time"
)

type Authorization interface {
	CreateUser(user models.User) (uint, error)
	GenerateToken(username, password string) (Tokens, error)
	ParseToken(token string) (string, error)
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

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type ServicesConfig struct {
	Repos              *repository.Repository
	TokenAdministrator auth.TokenAdministrator
	RefreshTokenTTL    time.Duration
	Cipher             cipher.PasswordCipher
}

func NewService(servicesConfig ServicesConfig) *Service {
	return &Service{
		Authorization: NewAuthService(
			servicesConfig.Repos.Authorization,
			servicesConfig.TokenAdministrator,
			servicesConfig.RefreshTokenTTL,
			servicesConfig.Cipher,
		),
		List: NewListService(servicesConfig.Repos.List),
	}
}
