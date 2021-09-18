package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/hostingvk4/badgerList/internal/models"
	"github.com/hostingvk4/badgerList/internal/repository"
	"github.com/hostingvk4/badgerList/pkg/auth"
	"strconv"
	"time"
)

const (
	salt = "aweawddsadfas23423asda"
)

type AuthService struct {
	repo               repository.Authorization
	tokenAdministrator auth.TokenAdministrator
	refreshTokenTTL    time.Duration
}

func NewAuthService(
	repo repository.Authorization,
	tokenAdministrator auth.TokenAdministrator,
	refreshTokenTTL time.Duration) *AuthService {
	return &AuthService{repo: repo, tokenAdministrator: tokenAdministrator, refreshTokenTTL: refreshTokenTTL}
}

func (s *AuthService) CreateUser(user models.User) (uint, error) {
	user.Password = generatePasswordHash(user.Password)
	id, err := s.repo.CreateUser(user)
	return id, err
}
func (s *AuthService) GenerateToken(username, password string) (Tokens, error) {
	userModel, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return Tokens{}, err
	}
	if userModel.ID == 0 {
		return Tokens{}, errors.New("customer not found")
	}

	return s.createTokens(userModel.ID)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
func (s *AuthService) ParseToken(accessToken string) (string, error) {
	id, err := s.tokenAdministrator.Parse(accessToken)
	return id, err
}

func (s *AuthService) createTokens(userId uint) (Tokens, error) {
	var (
		res Tokens
		err error
	)

	res.AccessToken, err = s.tokenAdministrator.NewJWT(strconv.Itoa(int(userId)), s.refreshTokenTTL)
	if err != nil {
		return res, err
	}

	res.RefreshToken, err = s.tokenAdministrator.NewRefreshToken()
	if err != nil {
		return res, err
	}
	RefreshToken := models.RefreshToken{
		RefreshToken: res.RefreshToken,
		UserId:       userId,
		ExpiresAt:    time.Now().Add(s.refreshTokenTTL),
	}
	err = s.repo.SetRefreshToken(userId, RefreshToken)

	return res, err
}
