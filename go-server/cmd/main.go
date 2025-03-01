package main

import (
	"log"
	"os"

	"github.com/Mahaveer86619/Entity-Interviewer.so/internal/config"
	"github.com/Mahaveer86619/Entity-Interviewer.so/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	config.ConnectDatabase()

	// Set up routes
	router.GET("/admin", handlers.AdminHandler)

	// Serve static files
	router.Static("/css", "./views/css")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	router.Run(":" + port)
}
