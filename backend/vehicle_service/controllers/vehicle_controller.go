package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VehicleControllers struct {
	DB *sql.DB
}

// GetAllVehicles fetches all vehicles
func (vc *VehicleControllers) GetAllVehicles(c *gin.Context) {
	rows, err := vc.DB.Query("SELECT id, brand, model, vehicleType, area, personCapacity, price, image_url FROM vehicle")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch vehicles"})
		return
	}
	defer rows.Close()

	var vehicles []map[string]interface{}
	for rows.Next() {
		var vehicle struct {
			ID             int     `json:"id"`
			Brand          string  `json:"brand"`
			Model          string  `json:"model"`
			VehicleType    string  `json:"vehicleType"`
			Area           string  `json:"area"`
			PersonCapacity int     `json:"personCapacity"`
			Price          float64 `json:"price"`
			ImageURL       string  `json:"image_url"`
		}

		if err := rows.Scan(&vehicle.ID, &vehicle.Brand, &vehicle.Model, &vehicle.VehicleType, &vehicle.Area,
			&vehicle.PersonCapacity, &vehicle.Price, &vehicle.ImageURL); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse vehicle data"})
			return
		}

		vehicles = append(vehicles, map[string]interface{}{
			"id":             vehicle.ID,
			"brand":          vehicle.Brand,
			"model":          vehicle.Model,
			"vehicleType":    vehicle.VehicleType,
			"area":           vehicle.Area,
			"personCapacity": vehicle.PersonCapacity,
			"price":          vehicle.Price,
			"image_url":      vehicle.ImageURL,
		})
	}

	c.JSON(http.StatusOK, gin.H{"vehicles": vehicles})
}

// GetVehicleById fetches a specific vehicle by ID
func (vc *VehicleControllers) GetVehicleById(c *gin.Context) {
	vehicleID := c.Param("id")
	var vehicle struct {
		ID             int     `json:"id"`
		Brand          string  `json:"brand"`
		Model          string  `json:"model"`
		VehicleType    string  `json:"vehicleType"`
		Area           string  `json:"area"`
		PersonCapacity int     `json:"personCapacity"`
		Price          float64 `json:"price"`
		ImageURL       string  `json:"image_url"`
	}

	query := "SELECT id, brand, model, vehicleType, area, personCapacity, price, image_url FROM vehicle WHERE id = ?"
	err := vc.DB.QueryRow(query, vehicleID).Scan(&vehicle.ID, &vehicle.Brand, &vehicle.Model, &vehicle.VehicleType,
		&vehicle.Area, &vehicle.PersonCapacity, &vehicle.Price, &vehicle.ImageURL)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vehicle not found"})
		return
	}

	c.JSON(http.StatusOK, vehicle)
}
