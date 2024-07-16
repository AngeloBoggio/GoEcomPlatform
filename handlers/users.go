package handlers

import (
	"ecommerce-backend/config"
	"ecommerce-backend/middleware"
	"ecommerce-backend/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(c *gin.Context){
	var user models.User
	
	// Bind input for userLogin attempt
	if err := c.ShouldBindBodyWithJSON(&user); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error": "Invalid user data"})
	}

	// now check if there even is a user
	var existingUser models.User
	if err := config.DB.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User Does not exist"})
		return	
	}

	// If not user exists and you need to unhash password and check authentication 
	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Generate token 
	token, err := middleware.GenerateToken(existingUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return 
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user": existingUser,
	})
}


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