package services

import (
	"hacktiv8-final-project/models"
	"hacktiv8-final-project/params"
	"hacktiv8-final-project/repositories"
)

type SocialMediaService interface {
	CreateSocialMedia(userID int, request params.CreateSocialMediaRequest) (*params.CreateSocialMediaResponse, error)
	GetSocialMedia() (*params.GetSocialMediaResponses, error)
	UpdateSocialMediaById(id int, userId int, request params.UpdateSocialMediaRequest) (*params.UpdateSocialMediaResponse, error)
	DeleteSocialMediaById(Id, userId int) (*params.DeleteSocialMediaResponse, error)
}

type socialMediaService struct {
	socialMediaRepo repositories.SocialMediaRepo
}

func NewSocialMediaService(smr repositories.SocialMediaRepo) SocialMediaService {
	return &socialMediaService{
		socialMediaRepo: smr,
	}
}

func (sms *socialMediaService) CreateSocialMedia(userID int, request params.CreateSocialMediaRequest) (*params.CreateSocialMediaResponse, error) {
	socialMedia := models.SocialMedia{
		Name:           request.Name,
		SocialMediaUrl: request.SocialMediaUrl,
		UserID:         userID,
	}

	res, err := sms.socialMediaRepo.CreateSocialMedia(socialMedia)
	if err != nil {
		return nil, err
	}

	response := params.CreateSocialMediaResponse{
		ID:             res.ID,
		Name:           res.Name,
		SocialMediaUrl: res.SocialMediaUrl,
		UserID:         res.UserID,
		CreatedAt:      res.CreatedAt,
	}

	return &response, nil

}

func getSocialMediaParam(socialMediaModel models.SocialMedia) params.SocialMedia {
	user := params.User{
		ID:       &socialMediaModel.UserID,
		Username: socialMediaModel.User.Username,
		Email:    socialMediaModel.User.Email,
	}
	return params.SocialMedia{
		ID:             socialMediaModel.ID,
		Name:           socialMediaModel.Name,
		UserID:         socialMediaModel.UserID,
		CreatedAt:      socialMediaModel.CreatedAt,
		UpdatedAt:      socialMediaModel.UpdatedAt,
		SocialMediaUrl: socialMediaModel.SocialMediaUrl,
		User:           user,
	}
}

func getSocialMediaParams(socialMediaModels []models.SocialMedia) []params.SocialMedia {
	socialMediaParams := make([]params.SocialMedia, len(socialMediaModels))
	for idx, socialMediaModel := range socialMediaModels {
		socialMediaParams[idx] = getSocialMediaParam(socialMediaModel)
	}

	return socialMediaParams
}

func (sms *socialMediaService) GetSocialMedia() (*params.GetSocialMediaResponses, error) {
	res, err := sms.socialMediaRepo.GetSocialMedia()
	if err != nil {
		return nil, err
	}

	getSocialMediaResponse := params.GetSocialMediaResponses{
		SocialMedias: getSocialMediaParams(res),
	}

	return &getSocialMediaResponse, nil

}

func (sms *socialMediaService) UpdateSocialMediaById(Id int, userId int, request params.UpdateSocialMediaRequest) (*params.UpdateSocialMediaResponse, error) {
	socialMedia := models.SocialMedia{
		Name:           request.Name,
		SocialMediaUrl: request.SocialMediaUrl,
	}

	res, err := sms.socialMediaRepo.UpdateSocialMediaById(Id, userId, socialMedia)
	if err != nil {
		return nil, err
	}

	updateSocialMediaResponse := params.UpdateSocialMediaResponse{
		ID:             res.ID,
		Name:           res.Name,
		SocialMediaUrl: res.SocialMediaUrl,
		UserID:         res.UserID,
		UpdatedAt:      res.UpdatedAt,
	}

	return &updateSocialMediaResponse, nil
}

func (sms *socialMediaService) DeleteSocialMediaById(Id, userId int) (*params.DeleteSocialMediaResponse, error) {
	err := sms.socialMediaRepo.DeleteSocialMediaById(Id, userId)
	if err != nil {
		return nil, err
	}

	deleteSocialMediaResponse := params.DeleteSocialMediaResponse{
		Message: "Your social media has been successfully deleted",
	}

	return &deleteSocialMediaResponse, nil
}
