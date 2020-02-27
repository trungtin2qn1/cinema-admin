package payment

import (
	"cinema-admin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//CheckOut ...
func CheckOut(c *gin.Context) {
	paymentReq := models.Payment{}

	err := c.ShouldBind(&paymentReq)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Data or data type is invalid")
		return
	}

	payment, err := models.CreatePayment(paymentReq.Amount, paymentReq.TicketID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Server is busy")
		return
	}

	c.JSON(200, payment)
}
