package routes

import (
	"CNAD_Assignment/backend/user_service/controllers"

	"github.com/gin-gonic/gin"
)

// SetupUserRoutes sets up routes for user-related APIs
func SetupUserRoutes(router *gin.Engine, userControllers *controllers.UserControllers) {
	// Group user-related routes under /users
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/signup", userControllers.RegisterUser)         // For user signup
		userRoutes.POST("/signin", userControllers.LoginUser)            // For user login
		userRoutes.GET("/", userControllers.GetAllUsers)                 // To get all users
		userRoutes.GET("/:id", userControllers.GetUserDetails)           // To get user by ID
		userRoutes.PUT("/profile/:id", userControllers.UpdateUserProfile) // To update user profile
	}
}
