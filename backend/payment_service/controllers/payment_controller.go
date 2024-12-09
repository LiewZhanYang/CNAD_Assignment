package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PaymentController struct {
	DB *sql.DB
}

type Booking struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Address          string `json:"address"`
	PickUpLocation   string `json:"pickUpLocation"`
	PickUpDate       string `json:"pickUpDate"`
	PickUpTime       string `json:"pickUpTime"`
	DropOffLocation  string `json:"dropOffLocation"`
	DropOffDate      string `json:"dropOffDate"`
	DropOffTime      string `json:"dropOffTime"`
	CreditCardNumber string `json:"creditCardNumber"`
	VehicleID        int    `json:"vehicle_id"`
}

// GetBookingById fetches a booking by its ID
func (pc *PaymentController) GetBookingById(c *gin.Context) {
	id := c.Param("id")

	var booking Booking
	err := pc.DB.QueryRow(`
		SELECT id, user_id, address, pickUpLocation, pickUpDate, pickUpTime, dropOffLocation, dropOffDate, dropOffTime, creditCardNumber, vehicle_id 
		FROM booking WHERE id = ?`, id).
		Scan(
			&booking.ID, &booking.UserID, &booking.Address, &booking.PickUpLocation, &booking.PickUpDate,
			&booking.PickUpTime, &booking.DropOffLocation, &booking.DropOffDate, &booking.DropOffTime,
			&booking.CreditCardNumber, &booking.VehicleID,
		)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch booking"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"booking": booking})
}

// PostBooking creates a new booking
func (pc *PaymentController) PostBooking(c *gin.Context) {
	var booking Booking

	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request data",
			"details": err.Error(),
		})
		return
	}

	query := `
		INSERT INTO booking (
			user_id, address, pickUpLocation, pickUpDate, pickUpTime, 
			dropOffLocation, dropOffDate, dropOffTime, creditCardNumber, vehicle_id
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := pc.DB.Exec(
		query,
		booking.UserID,
		booking.Address,
		booking.PickUpLocation,
		booking.PickUpDate,
		booking.PickUpTime,
		booking.DropOffLocation,
		booking.DropOffDate,
		booking.DropOffTime,
		booking.CreditCardNumber,
		booking.VehicleID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create booking",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Booking created successfully"})
}

// PutBooking updates an existing booking
func (pc *PaymentController) PutBooking(c *gin.Context) {
	id := c.Param("id")
	var booking Booking

	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	_, err := pc.DB.Exec(`
		UPDATE booking 
		SET address=?, pickUpLocation=?, pickUpDate=?, pickUpTime=?, dropOffLocation=?, dropOffDate=?, dropOffTime=?, creditCardNumber=?, vehicle_id=? 
		WHERE id=?`,
		booking.Address, booking.PickUpLocation, booking.PickUpDate, booking.PickUpTime,
		booking.DropOffLocation, booking.DropOffDate, booking.DropOffTime,
		booking.CreditCardNumber, booking.VehicleID, id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update booking"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Booking updated successfully"})
}

// CancelBooking deletes a booking
func (pc *PaymentController) CancelBooking(c *gin.Context) {
	id := c.Param("id")

	_, err := pc.DB.Exec("DELETE FROM booking WHERE id=?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel booking"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Booking canceled successfully"})
}

// GetBookingByUserId fetches all bookings for a specific user
func (pc *PaymentController) GetBookingByUserId(c *gin.Context) {
	userID := c.Param("userid")

	rows, err := pc.DB.Query(`
		SELECT id, user_id, address, pickUpLocation, pickUpDate, pickUpTime, dropOffLocation, dropOffDate, dropOffTime, creditCardNumber, vehicle_id 
		FROM booking WHERE user_id = ?`, userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bookings"})
		return
	}
	defer rows.Close()

	bookings := []Booking{}
	for rows.Next() {
		var booking Booking
		if err := rows.Scan(
			&booking.ID, &booking.UserID, &booking.Address, &booking.PickUpLocation, &booking.PickUpDate,
			&booking.PickUpTime, &booking.DropOffLocation, &booking.DropOffDate, &booking.DropOffTime,
			&booking.CreditCardNumber, &booking.VehicleID,
		); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning booking data"})
			return
		}
		bookings = append(bookings, booking)
	}

	c.JSON(http.StatusOK, gin.H{"bookings": bookings})
}
