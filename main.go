package main

import (
	"ecommerce-backend/config"
	"ecommerce-backend/handlers"
	"ecommerce-backend/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
    godotenv.Load()
}

func main() {
	
	// Set Gin to release mode
	gin.SetMode(gin.ReleaseMode)

	// initialize Gin router
	router := gin.Default()

	// Set trusted procies
	router.SetTrustedProxies([]string{"127.0.0.1"})

	//Connect to the database
	config.ConnectDatabase()
	
	// Auto-migrate models
	config.DB.AutoMigrate(
		&models.User{},
        &models.ShoppingCart{},
        &models.CartItem{},
        &models.PaymentMethod{},
        &models.Product{},
        &models.Wishlist{},
        &models.WishlistItem{},
	)

	router.GET("/ping", func(c *gin.Context) {
		fmt.Println("Pong")
	})

	router.POST("/signup", handlers.Signup)
	router.POST("/login", handlers.Login)

	
	 // Run the Gin server
	 router.Run(":8080")
}