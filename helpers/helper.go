package helpers

import (
	"fmt"
	"hacktiv8-final-project/config"

	"golang.org/x/crypto/bcrypt"
)

const (
	SALT = 8
)

func HashPassword(pass string) (string, error) {
	password := []byte(pass)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, config.SALT)
	return string(hashedPassword), err
}

func ValidatePassword(hashed, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	fmt.Println(err)
	return err == nil
}
