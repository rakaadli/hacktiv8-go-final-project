package repositories

import (
	"errors"
	"fmt"
	"hacktiv8-final-project/models"

	"gorm.io/gorm"
)

type CommentRepo interface {
	CreateComment(request models.Comment) (*models.Comment, error)
	GetCommentsByUserId() ([]models.Comment, error)
	UpdateCommentById(ID int, userID int, request models.Comment) (*models.Comment, error)
	DeleteCommentById(ID, userID int) error
}

type commentRepo struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepo {
	db.AutoMigrate(models.Comment{})
	return &commentRepo{
		db: db,
	}
}

func (cr *commentRepo) CreateComment(comment models.Comment) (*models.Comment, error) {
	err := cr.db.Create(&comment).Error
	return &comment, err
}

func (cr *commentRepo) GetCommentsByUserId() ([]models.Comment, error) {
	var comments []models.Comment
	err := cr.db.Preload("User").Preload("Photo").Find(&comments).Error

	return comments, err
}

func (cr *commentRepo) UpdateCommentById(ID int, userID int, request models.Comment) (*models.Comment, error) {
	var result models.Comment
	err := cr.db.Transaction(func(tx *gorm.DB) error {
		err := tx.First(&result, ID).Error
		if err != nil {
			return err
		}

		if result.UserID != userID {
			err := errors.New("you're forbidden to update this data")
			return err
		}

		err = tx.Where("id=?", ID).Updates(&request).Error
		if err != nil {
			fmt.Println("2")
			return err
		}

		err = cr.db.Preload("Photo").First(&result, ID).Error
		if err != nil {
			return err
		}

		return nil
	})

	return &result, err
}

func (cr *commentRepo) DeleteCommentById(ID, userID int) error {
	var result models.Comment
	err := cr.db.Transaction(func(tx *gorm.DB) error {
		err := tx.First(&result).Error
		if err != nil {
			return err
		}

		if result.UserID != userID {
			err := errors.New("you're forbidden to delete this data")
			return err
		}

		err = tx.Where("id=?", ID).Delete(&result).Error
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
