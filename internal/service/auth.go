package service

import (
	"github.com/hostingvk4/badgerList/internal/models"
	"github.com/hostingvk4/badgerList/internal/repository"
	"github.com/hostingvk4/badgerList/pkg/auth"
	"github.com/hostingvk4/badgerList/pkg/cipher"
	"time"
)

type AuthService struct {
	repo               repository.Authorization
	tokenAdministrator auth.TokenAdministrator
	refreshTokenTTL    time.Duration
	tokenTTL           time.Duration
	cipher             cipher.PasswordCipher
}

func NewAuthService(
	repo repository.Authorization,
	tokenAdministrator auth.TokenAdministrator,
	refreshTokenTTL time.Duration,
	tokenTTL time.Duration,
	cipher cipher.PasswordCipher) *AuthService {
	return &AuthService{repo: repo, tokenAdministrator: tokenAdministrator, refreshTokenTTL: refreshTokenTTL, tokenTTL: tokenTTL, cipher: cipher}
}

func (s *AuthService) CreateUser(user models.User) (uint, error) {
	passwordHash, err := s.cipher.CreateHash(user.Password)
	if err != nil {
		return 0, err
	}
	user.Password = passwordHash
	id, err := s.repo.CreateUser(user)
	return id, err
}
func (s *AuthService) GenerateToken(username, password string) (Tokens, error) {
	passwordHash, err := s.cipher.CreateHash(password)
	userModel, err := s.repo.GetUser(username, passwordHash)
	if err != nil {
		return Tokens{}, err
	}

	return s.createTokens(userModel.ID)
}

func (s *AuthService) ParseToken(accessToken string) (uint, error) {
	id, err := s.tokenAdministrator.Parse(accessToken)
	return id, err
}

func (s *AuthService) createTokens(userId uint) (Tokens, error) {
	var (
		res Tokens
		err error
	)

	res.AccessToken, err = s.tokenAdministrator.NewJWT(userId, s.tokenTTL)
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
	err = s.repo.SetRefreshToken(RefreshToken)

	return res, err
}

func (s *AuthService) RefreshToken(userId uint, oldRefreshToken string) (Tokens, error) {
	var (
		res Tokens
		err error
	)
	res.AccessToken, err = s.tokenAdministrator.NewJWT(userId, s.tokenTTL)
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
	err = s.repo.UpdateRefreshToken(oldRefreshToken, RefreshToken)

	return res, err
}
