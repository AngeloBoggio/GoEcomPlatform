package main

import (
	"ecommerce-backend/config"
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
	
	router.GET("/ping", func(c *gin.Context) {
		fmt.Println("Pong")
	})

	 // Run the Gin server
	 router.Run(":8080")
}