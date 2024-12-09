package routes

import (
	"CNAD_Assignment/backend/payment_service/controllers"

	"github.com/gin-gonic/gin"
)

func SetupInvoiceRoutes(router *gin.Engine, ic *controllers.InvoiceController) {
	invoiceGroup := router.Group("/invoices")
	{
		// Route to generate and send invoice
		invoiceGroup.POST("/generate", ic.GenerateAndSendInvoice)
	}
}
