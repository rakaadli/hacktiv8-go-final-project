package controllers

import (
	"hacktiv8-final-project/helpers"
	"hacktiv8-final-project/params"
	"hacktiv8-final-project/services"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	UpdateUserByID(ctx *gin.Context)
	DeleteUserByID(ctx *gin.Context)
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

	//two step validate (1)
	_, errCreate := govalidator.ValidateStruct(user)
	if errCreate != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//two step validate (2)
	err = helpers.ValidateUserRegisterRequest(user)
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

func (uc *userContorller) Login(ctx *gin.Context) {
	user := params.LoginUserRequest{}
	err := ctx.ShouldBind(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = helpers.ValidateUserLoginRequest(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := uc.userService.Login(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (uc *userContorller) UpdateUserByID(ctx *gin.Context) {
	user := params.UpdateUserRequest{}
	err := ctx.ShouldBind(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = helpers.ValidateUserUpdateRequest(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	Idstr := ctx.Param("id")
	Id, err := strconv.Atoi(Idstr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := int(ctx.Keys["id"].(float64))
	if Id != userID {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": "you're forbidden to update this data",
		})
		return
	}

	res, err := uc.userService.UpdateUserByID(Id, user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)

}

func (uc *userContorller) DeleteUserByID(ctx *gin.Context) {
	Idstr := ctx.Param("id")
	Id, err := strconv.Atoi(Idstr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userId := int(ctx.Keys["id"].(float64))
	if Id != userId {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": "you're forbidden to delete this data",
		})
		return
	}

	res, err := uc.userService.DeleteUserByID(Id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
