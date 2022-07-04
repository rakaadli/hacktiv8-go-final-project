package routers

import (
	"hacktiv8-final-project/controllers"
	"hacktiv8-final-project/middlewares"

	"github.com/gin-gonic/gin"
)

type socialMediaRouter struct {
	router                *gin.Engine
	socialMediaController controllers.SocialMediaController
}

func NewSocialMediaRouter(r *gin.Engine, smc controllers.SocialMediaController) *socialMediaRouter {
	return &socialMediaRouter{
		router:                r,
		socialMediaController: smc,
	}
}

func (smr *socialMediaRouter) Setup() {
	socialMedia := smr.router.Group("/socialmedia")
	{
		socialMedia.Use(middlewares.Auth())
		socialMedia.POST("/", smr.socialMediaController.CreateSocialMedia)
	}
}
