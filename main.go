package main

import (
	"github.com/gin-gonic/gin"
	"menu-api/configs"
	"menu-api/routes"
)

func main() {
	router := gin.Default()

	configs.ConnectDB()

	routes.UserRoute(router)

	router.Run("localhost:8080")
}
