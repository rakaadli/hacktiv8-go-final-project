package repositories

import (
	"hacktiv8-final-project/models"

	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(user *models.User) (*models.User, error)
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
