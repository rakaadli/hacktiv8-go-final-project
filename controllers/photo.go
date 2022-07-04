package controllers

import (
	"hacktiv8-final-project/helpers"
	"hacktiv8-final-project/params"
	"hacktiv8-final-project/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PhotoController interface {
	CreatePhoto(ctx *gin.Context)
}

type photoController struct {
	photoService services.PhotoService
}

func NewPhotoController(ps services.PhotoService) PhotoController {
	return &photoController{
		photoService: ps,
	}
}

func (pc *photoController) CreatePhoto(ctx *gin.Context) {
	request := params.CreatePhotoRequest{}
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = helpers.ValidatePhotoRequest(request.Title, request.PhotoUrl)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := int(ctx.Keys["id"].(float64))
	res, err := pc.photoService.CreatePhoto(userID, request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, res)
}
