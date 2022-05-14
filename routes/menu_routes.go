package routes

import (
	"github.com/gin-gonic/gin"
	"menu-api/controllers"
)

func UserRoute(router *gin.Engine) {

	router.GET("/menus", controllers.GetAllMenus())
	router.POST("/menus", controllers.CreateMenu())
}
