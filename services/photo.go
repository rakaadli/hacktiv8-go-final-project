package services

import (
	"hacktiv8-final-project/models"
	"hacktiv8-final-project/params"
	"hacktiv8-final-project/repositories"
)

type PhotoService interface {
	CreatePhoto(userId int, request params.CreatePhotoRequest) (*params.CreatePhotoResponse, error)
	GetAllPhotosByUserId() ([]params.GetPhotoResponse, error)
	UpdatePhotoById(Id int, userId int, request params.UpdatePhotoRequest) (*params.UpdatePhotoResponse, error)
	DeletePhotoById(Id int, userId int) (*params.DeletePhotoResponse, error)
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

func getPhotoResponse(photoModel models.Photo) params.GetPhotoResponse {
	userParam := params.User{
		Username: photoModel.User.Username,
		Email:    photoModel.User.Email,
	}

	return params.GetPhotoResponse{
		ID:        int(photoModel.ID),
		Title:     photoModel.Title,
		Caption:   photoModel.Caption,
		PhotoUrl:  photoModel.PhotoUrl,
		User:      userParam,
		CreatedAt: photoModel.CreatedAt,
		UpdatedAt: photoModel.UpdatedAt,
	}
}

func getPhotoResponses(photoModels []models.Photo) []params.GetPhotoResponse {
	getPhotoResponses := make([]params.GetPhotoResponse, len(photoModels))
	for photoId, photoModel := range photoModels {
		getPhotoResponses[photoId] = getPhotoResponse(photoModel)
	}

	return getPhotoResponses
}

func (ps *photoService) GetAllPhotosByUserId() ([]params.GetPhotoResponse, error) {
	result, err := ps.photoRepo.GetAllPhotosByUserId()
	if err != nil {
		return nil, err
	}

	return getPhotoResponses(result), nil
}

func (ps *photoService) UpdatePhotoById(Id int, userId int, request params.UpdatePhotoRequest) (*params.UpdatePhotoResponse, error) {
	photoModel := models.Photo{
		Title:    request.Title,
		PhotoUrl: request.PhotoUrl,
		Caption:  request.Caption,
		UserID:   userId,
	}

	res, err := ps.photoRepo.UpdatePhotoById(Id, photoModel)
	if err != nil {
		return nil, err
	}

	photoParam := params.UpdatePhotoResponse{
		ID:        int(res.ID),
		Title:     res.Title,
		Caption:   res.Caption,
		PhotoUrl:  res.PhotoUrl,
		UserID:    res.UserID,
		UpdatedAt: res.UpdatedAt,
	}

	return &photoParam, nil

}
func (ps *photoService) DeletePhotoById(Id int, userId int) (*params.DeletePhotoResponse, error) {
	err := ps.photoRepo.DeletePhotoById(Id, userId)
	if err != nil {
		return nil, err
	}

	deletePhotoResponse := params.DeletePhotoResponse{
		Message: "Your photo has been successfully deleted",
	}

	return &deletePhotoResponse, nil
}
