package cipher

import (
	"crypto/sha1"
	"fmt"
)

type PasswordCipher interface {
	CreateHash(password string) (string, error)
}

type Cipher struct {
	passwordSalt string
}

func NewCipher(passwordSalt string) *Cipher {
	return &Cipher{passwordSalt: passwordSalt}
}

func (h *Cipher) CreateHash(password string) (string, error) {
	hash := sha1.New()

	if _, err := hash.Write([]byte(password)); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum([]byte(h.passwordSalt))), nil
}
