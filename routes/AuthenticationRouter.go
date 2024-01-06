package routes

import (
	controller "Gate/controllers"

	"github.com/gin-gonic/gin"
)

func AuthJWTroutes(routes *gin.Engine) {
	routes.POST("users/signup", controller.SignUp())
	routes.POST("users/login", controller.Login())
}
