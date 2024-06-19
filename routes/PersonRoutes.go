package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mhusendev/go-mysql-api/controllers"
	"github.com/mhusendev/go-mysql-api/middleware"
)

func PersonRouter(router *gin.Engine) {

	routes := router.Group("api/v1/persons")
	routes.Use(middleware.JWTMiddleware())
	routes.POST("", controllers.PersonCreate)
	routes.GET("", controllers.PersonGet)
	routes.GET("/:id", controllers.PersonGetById)
	routes.PUT("/:id", controllers.PersonUpdate)
	routes.DELETE("/:id", controllers.PersonDelete)

}
