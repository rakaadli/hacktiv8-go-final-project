package services

import (
	"hacktiv8-final-project/models"
	"hacktiv8-final-project/params"
	"hacktiv8-final-project/repositories"
)

type PhotoService interface {
	CreatePhoto(userId int, request params.CreatePhotoRequest) (*params.CreatePhotoResponse, error)
	// GetPhotosByUserId() ([]params.GetPhotoResponse, error)
	// UpdatePhotoById(Id int, userId int, request params.UpdatePhotoRequest) (*params.UpdatePhotoResponse, error)
	// DeletePhotoById(Id int, userId int) (*params.DeletePhotoResponse, error)
}

type photoService struct {
	photoRepo repositories.PhotoRepo
}

func NewPhotoService(pr repositories.PhotoRepo) PhotoService {
	return &photoService{
		photoRepo: pr,
	}
}

func (ps *photoService) CreatePhoto(userID int, request params.CreatePhotoRequest) (*params.CreatePhotoResponse, error) {
	photo := models.Photo{
		Title:    request.Title,
		Caption:  request.Caption,
		PhotoUrl: request.PhotoUrl,
		UserID:   userID,
	}

	res, err := ps.photoRepo.CreatePhoto(photo)
	if err != nil {
		return nil, err
	}

	createPhotoResponse := params.CreatePhotoResponse{
		ID:        int(res.ID),
		Title:     res.Title,
		Caption:   res.Caption,
		PhotoUrl:  res.PhotoUrl,
		UserID:    res.UserID,
		CreatedAt: res.CreatedAt,
	}
	return &createPhotoResponse, err
}
