package repositories

import (
	"hacktiv8-final-project/models"

	"gorm.io/gorm"
)

type PhotoRepo interface {
	CreatePhoto(photo models.Photo) (*models.Photo, error)
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
