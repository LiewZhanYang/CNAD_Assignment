package routes

import (
	"CNAD_Assignment/backend/vehicle_service/controllers"
	"github.com/gin-gonic/gin"
)

// SetupVehicleRoutes sets up routes for vehicle-related APIs
func SetupVehicleRoutes(router *gin.Engine, vehicleControllers *controllers.VehicleControllers) {
	// Group vehicle-related routes under /vehicles
	vehicleRoutes := router.Group("/vehicles")
	{
		vehicleRoutes.GET("/", vehicleControllers.GetAllVehicles)
		vehicleRoutes.GET("/:id", vehicleControllers.GetVehicleById)
	}
}
