package services

import (
	"hacktiv8-final-project/models"
	"hacktiv8-final-project/params"
	"hacktiv8-final-project/repositories"
)

type CommentService interface {
	CreateComment(userID int, request params.CreateCommentRequest) (*params.CreateCommentResponse, error)
	GetCommentsByUserId() ([]params.GetCommentResponse, error)
	UpdateCommentById(Id int, userId int, request params.UpdateCommentRequest) (*params.UpdateCommentResponse, error)
	DeleteCommentById(Id, userId int) (*params.DeleteCommentResponse, error)
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

func getCommentResponse(commentModel models.Comment) params.GetCommentResponse {
	user := params.User{
		ID:       &commentModel.User.ID,
		Username: commentModel.User.Username,
		Email:    commentModel.User.Email,
	}

	photo := params.Photo{
		ID:       commentModel.Photo.ID,
		Title:    commentModel.Photo.Title,
		Caption:  commentModel.Photo.Caption,
		PhotoUrl: commentModel.Photo.PhotoUrl,
		UserID:   commentModel.Photo.UserID,
	}

	return params.GetCommentResponse{
		ID:        commentModel.ID,
		Message:   commentModel.Message,
		PhotoID:   commentModel.PhotoID,
		UserID:    commentModel.UserID,
		CreatedAt: commentModel.CreatedAt,
		UpdatedAt: commentModel.UpdatedAt,
		User:      user,
		Photo:     photo,
	}
}

func getCommentResponses(commentModels []models.Comment) []params.GetCommentResponse {
	getCommentResponses := make([]params.GetCommentResponse, len(commentModels))
	for idx, commentModel := range commentModels {
		getCommentResponses[idx] = getCommentResponse(commentModel)
	}

	return getCommentResponses
}

func (cs *commentService) GetCommentsByUserId() ([]params.GetCommentResponse, error) {
	res, err := cs.commentRepo.GetCommentsByUserId()
	if err != nil {
		return nil, err
	}

	return getCommentResponses(res), nil
}

func (cs *commentService) UpdateCommentById(ID int, userID int, request params.UpdateCommentRequest) (*params.UpdateCommentResponse, error) {
	commentModel := models.Comment{
		Message: request.Message,
	}

	res, err := cs.commentRepo.UpdateCommentById(ID, userID, commentModel)
	if err != nil {
		return nil, err
	}

	updateCommentResponse := params.UpdateCommentResponse{
		ID:        res.ID,
		Message:   res.Message,
		PhotoUrl:  res.Photo.PhotoUrl,
		UserID:    res.UserID,
		UpdatedAt: res.UpdatedAt,
	}

	return &updateCommentResponse, nil
}

func (cs *commentService) DeleteCommentById(ID, userID int) (*params.DeleteCommentResponse, error) {
	err := cs.commentRepo.DeleteCommentById(ID, userID)
	if err != nil {
		return nil, err
	}

	deleteCommentResponse := params.DeleteCommentResponse{
		Message: "Your comment has been successfully deleted",
	}

	return &deleteCommentResponse, err
}
