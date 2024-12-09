package controllers

import (
	"bytes"
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"net/smtp"
	"time"

	"github.com/gin-gonic/gin"
)

type InvoiceController struct {
	DB *sql.DB
}

// Invoice struct
type Invoice struct {
	ID          int       `json:"id"`
	BookingID   int       `json:"booking_id" binding:"required"`
	UserID      int       `json:"user_id" binding:"required"`
	Amount      float64   `json:"amount" binding:"required"`
	InvoiceDate time.Time `json:"invoice_date"`
	Status      string    `json:"status" binding:"required"`
	UserEmail   string    `json:"user_email"`
	UserName    string    `json:"user_name"`
}

// GenerateAndSendInvoice generates an invoice and sends it to the user's email
func (ic *InvoiceController) GenerateAndSendInvoice(c *gin.Context) {
	var invoice Invoice
	if err := c.ShouldBindJSON(&invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "details": err.Error()})
		return
	}

	// Verify user_id exists in the user table
	var userExists bool
	err := ic.DB.QueryRow(`SELECT EXISTS(SELECT 1 FROM user WHERE id = ?)`, invoice.UserID).Scan(&userExists)
	if err != nil || !userExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or non-existent user_id"})
		return
	}

	// Fetch user email and name from the database
	var userEmail, userName string
	err = ic.DB.QueryRow(`SELECT email, name FROM user WHERE id = ?`, invoice.UserID).Scan(&userEmail, &userName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user details", "details": err.Error()})
		return
	}
	invoice.UserEmail = userEmail
	invoice.UserName = userName

	// Set the current time as the invoice date
	invoice.InvoiceDate = time.Now()

	// Insert the invoice into the database
	query := `
		INSERT INTO invoice (booking_id, user_id, amount, invoice_date, status)
		VALUES (?, ?, ?, ?, ?)
	`
	result, err := ic.DB.Exec(query, invoice.BookingID, invoice.UserID, invoice.Amount, invoice.InvoiceDate, invoice.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate invoice", "details": err.Error()})
		return
	}

	// Get the last inserted invoice ID
	invoiceID, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve invoice ID", "details": err.Error()})
		return
	}

	invoice.ID = int(invoiceID)

	// Send the invoice via email
	if err := sendInvoiceEmail(invoice); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send invoice email", "details": err.Error()})
		return
	}

	// Respond with the invoice details
	c.JSON(http.StatusCreated, gin.H{"message": "Invoice generated and sent successfully", "invoice": invoice})
}

func sendInvoiceEmail(invoice Invoice) error {
	// Replace with your SMTP configuration
	smtpHost := "smtp.example.com"
	smtpPort := "587"
	senderEmail := "your-email@example.com"
	senderPassword := "your-email-password"

	// Create email body using an HTML template
	emailTemplate := `
		<html>
		<body>
			<h1>Invoice</h1>
			<p>Dear {{.UserName}},</p>
			<p>Thank you for your payment. Here are the details of your invoice:</p>
			<table>
				<tr><th>Invoice ID</th><td>{{.ID}}</td></tr>
				<tr><th>Booking ID</th><td>{{.BookingID}}</td></tr>
				<tr><th>Amount</th><td>${{.Amount}}</td></tr>
				<tr><th>Date</th><td>{{.InvoiceDate}}</td></tr>
				<tr><th>Status</th><td>{{.Status}}</td></tr>
			</table>
			<p>If you have any questions, feel free to contact us.</p>
			<p>Best regards,</p>
			<p>MyCarRental Team</p>
		</body>
		</html>
	`

	tmpl, err := template.New("email").Parse(emailTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse email template: %v", err)
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, invoice); err != nil {
		return fmt.Errorf("failed to execute email template: %v", err)
	}

	// Setup email message
	message := fmt.Sprintf("Subject: Your Invoice\nContent-Type: text/html\n\n%s", body.String())

	// Send email
	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpHost)
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, senderEmail, []string{invoice.UserEmail}, []byte(message))
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}
