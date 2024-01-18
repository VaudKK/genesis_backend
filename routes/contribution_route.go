package routes

import (
	"genesis/controllers"

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
	router.POST("/contribution", r.controller.AddContribution())
	router.GET("/contribution/:id", r.controller.GetContributionById())
}
