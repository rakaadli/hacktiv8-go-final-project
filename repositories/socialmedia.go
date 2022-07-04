package repositories

import (
	"hacktiv8-final-project/models"

	"gorm.io/gorm"
)

type SocialMediaRepo interface {
	CreateSocialMedia(socialMedia models.SocialMedia) (*models.SocialMedia, error)
}

type socialMediaRepo struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) SocialMediaRepo {
	db.AutoMigrate(&models.SocialMedia{})
	return &socialMediaRepo{
		db: db,
	}
}

func (smr *socialMediaRepo) CreateSocialMedia(socialMedia models.SocialMedia) (*models.SocialMedia, error) {
	err := smr.db.Create(&socialMedia).Error
	return &socialMedia, err
}
