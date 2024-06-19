package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mhusendev/go-mysql-api/configs"
	"github.com/mhusendev/go-mysql-api/models"
	"github.com/mhusendev/go-mysql-api/utils"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	fmt.Println("result:", input)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.Person
	if err := configs.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate refresh token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":        token,
		"refreshToken": refreshToken,
	})
}

func RefreshToken(c *gin.Context) {
	var request struct {
		RefreshToken string `json:"refreshToken"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	claims, err := utils.ValidateToken(request.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	newToken, err := utils.GenerateJWT(claims.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate new Access token"})
		return
	}
	newRefToken, err := utils.GenerateRefreshToken(claims.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to generate new Refresh token"})
	}

	c.JSON(http.StatusOK, gin.H{"accessToken": newToken, "refreshToken": newRefToken})
}
