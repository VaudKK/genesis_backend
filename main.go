package main

import (
	"genesis/configs"
	"genesis/controllers"
	"genesis/routes"
	"genesis/service"

	"github.com/gin-gonic/gin"
)

var (
	contributionSerivce    service.ContributionService        = service.NewContributionService(configs.GetCollection(configs.DB, "contributions"))
	contributionController controllers.ContributionController = controllers.NewContributionController(contributionSerivce)
	contributionRoutes     routes.ContributionRoute           = routes.NewContributionRoutes(contributionController)
)

func main() {
	router := gin.Default()

	//routes
	contributionRoutes.Routes(router)

	router.Run("localhost:8080")
}
