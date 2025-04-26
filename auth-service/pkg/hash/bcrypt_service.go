package hash

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrHashPassword    = errors.New("failed to hash password")
	ErrComparePassword = errors.New("failed to compare passwords")
	ErrInvalidPassword = errors.New("invalid password")
)

type BcryptService struct{}

func NewBcryptService() *BcryptService {
	return &BcryptService{}
}

func (bc *BcryptService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", ErrHashPassword
	}

	return string(hashedPassword), nil

}

func (bc *BcryptService) CompareHashAndPassword(hash string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	if err == bcrypt.ErrMismatchedHashAndPassword {
		return ErrInvalidPassword
	}

	if err != nil {
		return ErrComparePassword
	}

	return nil
}
