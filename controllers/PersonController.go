package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mhusendev/go-mysql-api/configs"
	"github.com/mhusendev/go-mysql-api/models"
	"golang.org/x/crypto/bcrypt"
)

type PersonRequestBody struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func PersonCreate(c *gin.Context) {

	body := PersonRequestBody{}

	c.BindJSON(&body)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	person := &models.Person{Name: body.Name, Address: body.Address, Phone: body.Phone, Username: body.Username, Password: string(hashedPassword)}

	result := configs.DB.Create(&person)

	if result.Error != nil {
		fmt.Println("result:", true)
		c.JSON(500, gin.H{"Error": "Failed to insert"})
		return
	}

	c.JSON(200, &person)
}

func PersonGet(c *gin.Context) {
	var persons []models.Person
	configs.DB.Find(&persons)
	c.JSON(200, &persons)

}

func PersonGetById(c *gin.Context) {
	id := c.Param("id")
	var person models.Person
	configs.DB.First(&person, id)
	c.JSON(200, &person)

}

func PersonUpdate(c *gin.Context) {

	id := c.Param("id")
	var person models.Person
	configs.DB.First(&person, id)

	body := PersonRequestBody{}
	c.BindJSON(&body)
	data := &models.Person{Name: body.Name, Address: body.Address, Phone: body.Phone}

	result := configs.DB.Model(&person).Updates(data)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": true, "message": "Failed to update"})
		return
	}

	c.JSON(200, &person)
}

func PersonDelete(c *gin.Context) {
	id := c.Param("id")
	var person models.Person
	configs.DB.Delete(&person, id)
	c.JSON(200, gin.H{"deleted": true})

}
