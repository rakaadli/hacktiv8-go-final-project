package repositories

import (
	"hacktiv8-final-project/models"

	"gorm.io/gorm"
)

type CommentRepo interface {
	CreateComment(request models.Comment) (*models.Comment, error)
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
