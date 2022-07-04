package services

import (
	"hacktiv8-final-project/helpers"
	"hacktiv8-final-project/models"
	"hacktiv8-final-project/params"
	"hacktiv8-final-project/repositories"
)

type UserService interface {
	Register(request params.RegisterUserRequest) (*params.RegisterUserResponse, error)
}

type userService struct {
	userRepo repositories.UserRepo
}

func NewUserService(ur repositories.UserRepo) UserService {
	return &userService{
		userRepo: ur,
	}
}

func (us *userService) Register(request params.RegisterUserRequest) (*params.RegisterUserResponse, error) {
	hashedPassword, err := helpers.HashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	userModel := models.User{
		Username: request.Username,
		Password: hashedPassword,
		Email:    request.Email,
		Age:      request.Age,
	}

	user, err := us.userRepo.CreateUser(&userModel)
	if err != nil {
		return nil, err
	}

	registerResponse := params.RegisterUserResponse{
		ID:       user.ID,
		Age:      user.Age,
		Email:    user.Email,
		Username: user.Username,
	}
	return &registerResponse, nil

}
