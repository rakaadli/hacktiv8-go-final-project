package controllers

import (
	"hacktiv8-final-project/helpers"
	"hacktiv8-final-project/params"
	"hacktiv8-final-project/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentController interface {
	CreateComment(ctx *gin.Context)
	GetCommentsByUserId(ctx *gin.Context)
	UpdateCommentById(ctx *gin.Context)
	DeleteCommentById(ctx *gin.Context)
}

type commentController struct {
	commentService services.CommentService
}

func NewCommentController(cs services.CommentService) CommentController {
	return &commentController{
		commentService: cs,
	}
}

func (cc *commentController) CreateComment(ctx *gin.Context) {
	var request params.CreateCommentRequest
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = helpers.ValidateCommentRequest(request.Message)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := int(ctx.Keys["id"].(float64))
	res, err := cc.commentService.CreateComment(userID, request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

func (cc *commentController) GetCommentsByUserId(ctx *gin.Context) {
	res, err := cc.commentService.GetCommentsByUserId()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (cc *commentController) UpdateCommentById(ctx *gin.Context) {
	IDStr := ctx.Param("commentid")
	Id, err := strconv.Atoi(IDStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var request params.UpdateCommentRequest
	err = ctx.ShouldBind(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = helpers.ValidateCommentRequest(request.Message)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userId := int(ctx.Keys["id"].(float64))
	res, err := cc.commentService.UpdateCommentById(Id, userId, request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (cc *commentController) DeleteCommentById(ctx *gin.Context) {
	IdStr := ctx.Param("commentid")
	Id, err := strconv.Atoi(IdStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userId := int(ctx.Keys["id"].(float64))
	res, err := cc.commentService.DeleteCommentById(Id, userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
