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


// GetBillingById fetches a billing record by its ID
func (pc *PaymentController) GetBillingById(c *gin.Context) {
    id := c.Param("id")

    var billing struct {
        ID        int     `json:"id"`
        BookingID int     `json:"booking_id"`
        Amount    float64 `json:"amount"`
        Status    string  `json:"status"`
    }

    err := pc.DB.QueryRow(`
        SELECT id, booking_id, amount, status FROM billing WHERE id = ?`, id,
    ).Scan(
        &billing.ID, &billing.BookingID, &billing.Amount, &billing.Status,
    )

    if err != nil {
        if err == sql.ErrNoRows {
            c.JSON(http.StatusNotFound, gin.H{"error": "Billing record not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch billing record"})
        }
        return
    }

    c.JSON(http.StatusOK, gin.H{"billing": billing})
}

// PostBilling creates a new billing record
func (pc *PaymentController) PostBilling(c *gin.Context) {
    var billing struct {
        BookingID int     `json:"booking_id" binding:"required"`
        Amount    float64 `json:"amount" binding:"required"`
        Status    string  `json:"status"`
    }

    if err := c.ShouldBindJSON(&billing); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid request data",
            "details": err.Error(),
        })
        return
    }

    // Ensure the booking exists before creating the billing
    var bookingExists bool
    err := pc.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM booking WHERE id = ?)", billing.BookingID).Scan(&bookingExists)
    if err != nil || !bookingExists {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid booking ID",
        })
        return
    }

    query := `
        INSERT INTO billing (booking_id, amount, status)
        VALUES (?, ?, ?)
    `
    _, err = pc.DB.Exec(query, billing.BookingID, billing.Amount, billing.Status)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to create billing record",
            "details": err.Error(),
        })
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "Billing record created successfully",
    })
}
