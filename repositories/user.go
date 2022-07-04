package repositories

import (
	"hacktiv8-final-project/models"

	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{db}
}

func (ur *userRepo) CreateUser(user *models.User) (*models.User, error) {
	return user, ur.db.Create(&user).Error
}

func (ur *userRepo) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := ur.db.Where("email=?", email).Find(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
