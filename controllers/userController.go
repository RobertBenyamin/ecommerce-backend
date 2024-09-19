package controllers

import (
	"ecommerce-backend/config"
	"ecommerce-backend/middlewares"
	"ecommerce-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }

    if user.Name == "" || user.Email == "" || user.Password == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "All fields (name, email, password) are required"})
        return
    }

    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
    user.Password = string(hashedPassword)

    config.DB.Create(&user)

    token, _ := middlewares.GenerateJWT(user.ID)

    c.JSON(http.StatusOK, gin.H{"token": token})
}

func LoginUser(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token, _ := middlewares.GenerateJWT(user.ID)

	c.JSON(http.StatusOK, gin.H{"token": token})
}
