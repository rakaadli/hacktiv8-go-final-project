package services

import (
	"hacktiv8-final-project/models"
	"hacktiv8-final-project/params"
	"hacktiv8-final-project/repositories"
)

type CommentService interface {
	CreateComment(userID int, request params.CreateCommentRequest) (*params.CreateCommentResponse, error)
}

type commentService struct {
	commentRepo repositories.CommentRepo
}

func NewCommentRepository(cr repositories.CommentRepo) CommentService {
	return &commentService{
		commentRepo: cr,
	}
}

func (cs *commentService) CreateComment(userID int, request params.CreateCommentRequest) (*params.CreateCommentResponse, error) {
	commentModel := models.Comment{
		Message: request.Message,
		PhotoID: request.PhotoID,
		UserID:  userID,
	}

	res, err := cs.commentRepo.CreateComment(commentModel)
	if err != nil {
		return nil, err
	}

	createCommentResponse := params.CreateCommentResponse{
		ID:        res.ID,
		Message:   res.Message,
		PhotoID:   res.PhotoID,
		UserID:    res.UserID,
		CreatedAt: res.CreatedAt,
	}

	return &createCommentResponse, nil

}
