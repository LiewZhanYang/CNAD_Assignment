package main

import (
	"CNAD_Assignment/backend/config"
	"CNAD_Assignment/backend/payment_service/controllers"
	"CNAD_Assignment/backend/payment_service/routes"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Connect to the database
	db := config.ConnectDatabase("DBCNAD_Assignment")
	defer db.Close()

	// Initialize controller
	paymentController := &controllers.PaymentController{DB: db}

	// Initialize Gin router
	router := gin.Default()

	// Add CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5500", "http://127.0.0.1:5500"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Setup routes
	routes.SetupPaymentRoutes(router, paymentController)

	// Start server
	log.Println("Payment Service is running on http://localhost:8082")
	router.Run(":8082")
}
