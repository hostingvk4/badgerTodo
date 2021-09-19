package auth

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"math/rand"
	"time"
)

type TokenAdministrator interface {
	NewJWT(userId uint, ttl time.Duration) (string, error)
	Parse(accessToken string) (uint, error)
	NewRefreshToken() (string, error)
}

type Administrator struct {
	signingKey string
}
type tokenClaims struct {
	jwt.StandardClaims
	UserId uint `json:"user_id"`
}

func NewAdministrator(signingKey string) (*Administrator, error) {
	if signingKey == "" {
		return nil, errors.New("error empty signing key")
	}

	return &Administrator{signingKey: signingKey}, nil
}

func (m *Administrator) NewJWT(userId uint, timeDuration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(timeDuration).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userId,
	})

	return token.SignedString([]byte(m.signingKey))
}

func (m *Administrator) Parse(accessToken string) (uint, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(m.signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}
func (m *Administrator) NewRefreshToken() (string, error) {
	tokenByte := make([]byte, 32)

	tokenSource := rand.NewSource(time.Now().Unix())
	tokenRandomString := rand.New(tokenSource)

	if _, err := tokenRandomString.Read(tokenByte); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", tokenByte), nil
}
