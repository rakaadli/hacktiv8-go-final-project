package controllers

import (
	"hacktiv8-final-project/params"
	"hacktiv8-final-project/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Register(ctx *gin.Context)
}

type userContorller struct {
	userService services.UserService
}

func NewUserController(us services.UserService) UserController {
	return &userContorller{
		userService: us,
	}
}

func (uc *userContorller) Register(ctx *gin.Context) {
	user := params.RegisterUserRequest{}
	err := ctx.ShouldBind(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := uc.userService.Register(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, res)
}
