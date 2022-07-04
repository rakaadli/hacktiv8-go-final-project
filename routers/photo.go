package routers

import (
	"hacktiv8-final-project/controllers"
	"hacktiv8-final-project/middlewares"

	"github.com/gin-gonic/gin"
)

type photoRouter struct {
	router          *gin.Engine
	photoController controllers.PhotoController
}

func NewPhotoRouter(g *gin.Engine, pc controllers.PhotoController) *photoRouter {
	return &photoRouter{
		router:          g,
		photoController: pc,
	}
}

func (pr *photoRouter) Setup() {
	photo := pr.router.Group("/photos")
	{
		photo.Use(middlewares.Auth())
		photo.POST("/", pr.photoController.CreatePhoto)
		photo.GET("/", pr.photoController.GetPhotos)
		photo.PUT("/:photoid", pr.photoController.UpdatePhotoById)
		photo.DELETE("/:photoid", pr.photoController.DeletePhotoById)
	}
}
