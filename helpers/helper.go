package helpers

import (
	"fmt"
	config "hacktiv8-final-project/configs"

	"golang.org/x/crypto/bcrypt"
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
