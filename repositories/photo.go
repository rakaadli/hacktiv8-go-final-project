package repositories

import (
	"errors"
	"hacktiv8-final-project/models"

	"gorm.io/gorm"
)

type PhotoRepo interface {
	CreatePhoto(photo models.Photo) (*models.Photo, error)
	GetAllPhotosByUserId() ([]models.Photo, error)
	UpdatePhotoById(Id int, photo models.Photo) (*models.Photo, error)
	DeletePhotoById(Id int, userId int) error
}

type photoRepo struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) PhotoRepo {
	return &photoRepo{
		db: db,
	}
}

func (pr *photoRepo) CreatePhoto(photo models.Photo) (*models.Photo, error) {
	err := pr.db.Create(&photo).Error
	return &photo, err
}

func (pr *photoRepo) GetAllPhotosByUserId() ([]models.Photo, error) {
	var photos []models.Photo
	err := pr.db.Preload("User").Find(&photos).Error

	return photos, err

}

func (pr *photoRepo) UpdatePhotoById(Id int, photo models.Photo) (*models.Photo, error) {
	var result models.Photo
	err := pr.db.Transaction(func(tx *gorm.DB) error {
		err := tx.First(&result, Id).Error
		if err != nil {
			return err
		}

		if uint(result.UserID) != uint(photo.UserID) {
			err := errors.New("you're forbidden to update this data")
			return err
		}

		err = tx.Where("id=?", Id).Updates(&photo).Find(&result).Error
		if err != nil {
			return err
		}

		return nil
	})

	return &result, err
}

func (pr *photoRepo) DeletePhotoById(Id int, userId int) error {
	var result models.Photo
	err := pr.db.Transaction(func(tx *gorm.DB) error {
		err := tx.First(&result, Id).Error
		if err != nil {
			return err
		}

		if uint(result.UserID) != uint(userId) {
			err := errors.New("you're forbidden to delete this data")
			return err
		}
		err = tx.Where("id=?", Id).Delete(&result).Error
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
