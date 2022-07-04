package main

import (
	"hacktiv8-final-project/config"
	"hacktiv8-final-project/controllers"
	"hacktiv8-final-project/repositories"
	"hacktiv8-final-project/routers"
	"hacktiv8-final-project/services"

	"github.com/gin-gonic/gin"
)

func main() {

	db := config.ConnectDB()
	route := gin.Default()

	userRepo := repositories.NewUserRepo(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)
	userRouter := routers.NewUserRouter(route, userController)
	userRouter.Setup()

	photoRepo := repositories.NewPhotoRepository(db)
	photoService := services.NewPhotoService(photoRepo)
	photoController := controllers.NewPhotoController(photoService)
	photoRouter := routers.NewPhotoRouter(route, photoController)
	photoRouter.Setup()

	route.Run(config.APP_PORT)

}
