package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

const (
	SALT = 8
)

func HashPassword(pass string) (string, error) {
	password := []byte(pass)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, SALT)
	return string(hashedPassword), err
}

func ValidatePassword(hashed, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}
