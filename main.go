package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mhusendev/go-mysql-api/configs"
	"github.com/mhusendev/go-mysql-api/routes"
)

func init() {
	configs.ConnectToDB()
}

func main() {

	r := gin.Default()
	routes.AuthRouter(r)
	routes.PersonRouter(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world from server Go.",
		})
	})
	r.Run()
}
