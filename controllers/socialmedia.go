package controllers

import (
	"hacktiv8-final-project/helpers"
	"hacktiv8-final-project/params"
	"hacktiv8-final-project/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SocialMediaController interface {
	CreateSocialMedia(ctx *gin.Context)
	GetSocialMedia(ctx *gin.Context)
	UpdateSocialMediaById(ctx *gin.Context)
	DeleteSocialMediaById(ctx *gin.Context)
}

type socialMediaController struct {
	socialMediaService services.SocialMediaService
}

func NewSocialMediaController(sms services.SocialMediaService) SocialMediaController {
	return &socialMediaController{
		socialMediaService: sms,
	}
}

func (smc *socialMediaController) CreateSocialMedia(ctx *gin.Context) {
	var request params.CreateSocialMediaRequest
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = helpers.ValidateSocialMediaRequest(request.Name, request.SocialMediaUrl)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := int(ctx.Keys["id"].(float64))
	res, err := smc.socialMediaService.CreateSocialMedia(userID, request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

func (smc *socialMediaController) GetSocialMedia(ctx *gin.Context) {
	res, err := smc.socialMediaService.GetSocialMedia()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (smc *socialMediaController) UpdateSocialMediaById(ctx *gin.Context) {
	IdStr := ctx.Param("socialMediaId")
	Id, err := strconv.Atoi(IdStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var request params.UpdateSocialMediaRequest
	err = ctx.ShouldBind(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = helpers.ValidateSocialMediaRequest(request.Name, request.SocialMediaUrl)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userId := int(ctx.Keys["id"].(float64))
	res, err := smc.socialMediaService.UpdateSocialMediaById(Id, userId, request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)

}

func (smc *socialMediaController) DeleteSocialMediaById(ctx *gin.Context) {
	IdStr := ctx.Param("socialMediaId")
	Id, err := strconv.Atoi(IdStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userId := int(ctx.Keys["id"].(float64))
	res, err := smc.socialMediaService.DeleteSocialMediaById(Id, userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
