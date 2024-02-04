package routes

import (
	"genesis/controllers"
	"genesis/middleware"

	"github.com/gin-gonic/gin"
)

type ContributionRoute interface {
	Routes(*gin.Engine)
}

type route struct {
	controller controllers.ContributionController
}

func NewContributionRoutes(controller controllers.ContributionController) ContributionRoute {
	return &route{
		controller: controller,
	}
}

func (r *route) Routes(router *gin.Engine) {
	//all routes related to contributions
	protectedRoutes := router.Group("/contributions")

	protectedRoutes.Use(middleware.AuthenticationMiddleware())

	protectedRoutes.POST("/add", r.controller.AddContribution())
	protectedRoutes.POST("/add-many", r.controller.AddManyContributions())
	protectedRoutes.GET("/:id", r.controller.GetContributionById())

}
