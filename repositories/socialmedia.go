package repositories

import (
	"errors"
	"hacktiv8-final-project/models"

	"gorm.io/gorm"
)

type SocialMediaRepo interface {
	CreateSocialMedia(socialMedia models.SocialMedia) (*models.SocialMedia, error)
	GetSocialMedia() ([]models.SocialMedia, error)
	UpdateSocialMediaById(Id int, userId int, socialMedia models.SocialMedia) (*models.SocialMedia, error)
	DeleteSocialMediaById(Id int, userId int) error
}

type socialMediaRepo struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) SocialMediaRepo {
	return &socialMediaRepo{
		db: db,
	}
}

func (smr *socialMediaRepo) CreateSocialMedia(socialMedia models.SocialMedia) (*models.SocialMedia, error) {
	err := smr.db.Create(&socialMedia).Error
	return &socialMedia, err
}

func (smr *socialMediaRepo) GetSocialMedia() ([]models.SocialMedia, error) {
	var socialMedias []models.SocialMedia
	err := smr.db.Preload("User").Find(&socialMedias).Error
	return socialMedias, err
}

func (smr *socialMediaRepo) UpdateSocialMediaById(ID int, userID int, socialMedia models.SocialMedia) (*models.SocialMedia, error) {
	var result models.SocialMedia
	err := smr.db.Transaction(func(tx *gorm.DB) error {
		err := tx.First(&result, ID).Error
		if err != nil {
			return err
		}

		if uint(result.UserID) != uint(userID) {
			err := errors.New("you're forbidden to delete this data")
			return err
		}

		err = tx.Where("id=?", ID).Updates(&socialMedia).Find(&result).Error
		if err != nil {
			return err
		}
		return nil
	})

	return &result, err
}

func (smr *socialMediaRepo) DeleteSocialMediaById(Id int, userId int) error {
	var result models.SocialMedia
	err := smr.db.Transaction(func(tx *gorm.DB) error {
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
