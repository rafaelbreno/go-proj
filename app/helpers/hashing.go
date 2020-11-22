package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(str string) (string, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(str), 14)

	return string(pass), err
}

func VerifyPassword(password string, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err
}
