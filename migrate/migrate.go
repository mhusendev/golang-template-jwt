package main

import (
	"github.com/mhusendev/go-mysql-api/configs"
	"github.com/mhusendev/go-mysql-api/models"
)

func init() {
	configs.ConnectToDB()
}

func main() {
	configs.DB.AutoMigrate(&models.Person{})
}
