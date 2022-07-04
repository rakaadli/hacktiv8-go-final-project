package helpers

import (
	"errors"
	"hacktiv8-final-project/params"
	"net/mail"
)

func ValidateUserRegisterRequest(request params.RegisterUserRequest) error {
	if len(request.Username) == 0 {
		return errors.New("username is required")
	}
	if len(request.Password) == 0 {
		return errors.New("password is required")
	}
	if len(request.Password) < 6 {
		return errors.New("password is required to have 6 minimum characters")
	}

	if len(request.Email) == 0 {
		return errors.New("email is required")
	}

	_, err := mail.ParseAddress(request.Email)
	if err != nil {
		return errors.New("email format is invalid")
	}

	if request.Age == 0 {
		return errors.New("age is required")
	}

	if request.Age < 8 {
		return errors.New("to register this app your age must be greater than 8")
	}

	return nil
}

func ValidateUserLoginRequest(request params.LoginUserRequest) error {
	if len(request.Email) == 0 {
		return errors.New("email is required")
	}

	_, err := mail.ParseAddress(request.Email)
	if err != nil {
		return errors.New("email format is invalid")
	}

	if len(request.Password) == 0 {
		return errors.New("password is required")
	}

	return nil
}

func ValidateUserUpdateRequest(request params.UpdateUserRequest) error {
	if len(request.Username) == 0 {
		return errors.New("username is required")
	}

	if len(request.Email) == 0 {
		return errors.New("email is required")
	}

	_, err := mail.ParseAddress(request.Email)
	if err != nil {
		return errors.New("email format is invalid")
	}

	return nil
}
