package routes

import (
	"github.com/mhusendev/go-mysql-api/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRouter(router *gin.Engine) {
	routes := router.Group("api/v1/auth")
	routes.POST("/login", controllers.Login) // New route for login
	routes.POST("/refreshToken", controllers.RefreshToken)
}
