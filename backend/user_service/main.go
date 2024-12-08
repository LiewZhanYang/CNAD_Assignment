package main

import (
	"CNAD_Assignment/backend/config"
	"CNAD_Assignment/backend/user_service/controllers"
	"CNAD_Assignment/backend/user_service/routes"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load("../.env") // Adjust path if .env file is elsewhere
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Connect to shared database
	db := config.ConnectDatabase("DBCNAD_Assignment") // Using DBCNAD_Assignment
	defer db.Close()

	// Initialize controllers
	userControllers := &controllers.UserControllers{DB: db}

	// Initialize Gin router
	router := gin.Default()

	// Add CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5500", "http://127.0.0.1:5500"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))


	// Setup routes
	routes.SetupUserRoutes(router, userControllers)

	// Start server
	log.Println("User Service is running on http://localhost:8080")
	router.Run(":8080")
}
