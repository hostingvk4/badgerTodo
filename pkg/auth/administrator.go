package auth

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"math/rand"
	"time"
)

type TokenAdministrator interface {
	NewJWT(userId string, ttl time.Duration) (string, error)
	Parse(accessToken string) (string, error)
	NewRefreshToken() (string, error)
}

type Administrator struct {
	signingKey string
}

func NewAdministrator(signingKey string) (*Administrator, error) {
	if signingKey == "" {
		return nil, errors.New("error empty signing key")
	}

	return &Administrator{signingKey: signingKey}, nil
}

func (m *Administrator) NewJWT(userId string, timeDuration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(timeDuration).Unix(),
		Subject:   userId,
	})

	return token.SignedString([]byte(m.signingKey))
}

func (m *Administrator) Parse(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(m.signingKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("error claims user token")
	}

	return claims["sub"].(string), nil
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
