package repositories

import (
	"hacktiv8-final-project/models"

	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	UpdateUserByID(ID int, user models.User) (*models.User, error)
	DeleteUserByID(ID int) error
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

func (ur *userRepo) UpdateUserByID(ID int, user models.User) (*models.User, error) {
	var result models.User
	err := ur.db.Transaction(func(tx *gorm.DB) error {
		err := tx.First(&result, ID).Error
		if err != nil {
			return err
		}

		err = tx.Where("id=?", ID).Updates(&user).Find(&result).Error
		if err != nil {
			return err
		}

		return nil
	})

	return &result, err
}

func (ur *userRepo) DeleteUserByID(ID int) error {
	var user models.User
	err := ur.db.Transaction(func(tx *gorm.DB) error {
		err := tx.First(&user, ID).Error
		if err != nil {
			return err
		}

		err = tx.Where("id=?", ID).Delete(&user).Error
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
