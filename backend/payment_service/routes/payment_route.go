package routes

import (
	"CNAD_Assignment/backend/payment_service/controllers"

	"github.com/gin-gonic/gin"
)

func SetupPaymentRoutes(router *gin.Engine, pc *controllers.PaymentController) {
	payment := router.Group("/payments")
	{
		payment.GET("/bookings/:id", pc.GetBookingById)       // Get booking by ID
		payment.GET("/user/:userid", pc.GetBookingByUserId)   // Get bookings by user ID
		payment.POST("/bookings", pc.PostBooking)            // Create a booking
		payment.PUT("/bookings/:id", pc.PutBooking)          // Update a booking
		payment.DELETE("/bookings/:id", pc.CancelBooking)    // Cancel a booking
	}
}
