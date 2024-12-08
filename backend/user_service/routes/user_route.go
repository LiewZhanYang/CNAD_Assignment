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
		userRoutes.POST("/register", userControllers.RegisterUser)
		userRoutes.POST("/login", userControllers.LoginUser)
		userRoutes.GET("/:id", userControllers.GetUserDetails)
		userRoutes.PUT("/:id", userControllers.UpdateUserProfile)
	}
}
