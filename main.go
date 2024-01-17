package main

import (
	"genesis/configs"
	"genesis/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//database
	configs.ConnectToDB()

	//routes
	routes.ContributionRoute(router)

	router.Run("localhost:8080")
}
