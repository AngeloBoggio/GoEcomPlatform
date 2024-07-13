package handlers

import (
	"ecommerce-backend/config"
	"ecommerce-backend/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Signup (c *gin.Context){
	var user models.User
	
	// Bind Input for our User with gin
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error": "Invalid user data"})
		return
	}

	// Now check if user already exists 
	var existingUser models.User
	if err := config.DB.Where("username = ?", user.Username).First(&existingUser).Error; err == nil{
		c.JSON(http.StatusConflict, gin.H{"error": "Username already taken"})
		return
	}else if !errors.Is(err, gorm.ErrRecordNotFound){
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user.Password = string(hashedPassword)
	// Create user
	if err := config.DB.Create(&user).Error; err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return 
	}
	
	// Successful message
	c.JSON(http.StatusCreated,user)

}