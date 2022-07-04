package services

import (
	"hacktiv8-final-project/models"
	"hacktiv8-final-project/params"
	"hacktiv8-final-project/repositories"
)

type SocialMediaService interface {
	CreateSocialMedia(userID int, request params.CreateSocialMediaRequest) (*params.CreateSocialMediaResponse, error)
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
