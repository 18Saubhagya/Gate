package routes

import (
	controller "Gate/controllers"
	"Gate/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(routes *gin.Engine) {
	routes.Use(middleware.Authenticate())
	routes.GET("users", controller.GetUsers())
	routes.GET("users/:user_id", controller.GetUser())
}
