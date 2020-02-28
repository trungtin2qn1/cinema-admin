package payment

import (
	"cinema-admin/models"
	"fmt"
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

	if paymentReq.TicketID == nil {
		if err != nil {
			c.JSON(http.StatusBadRequest, "Data or data type is invalid")
			return
		}
	}

	interfaceServiceKeyID, _ := c.Get("service-key-id")
	serviceKeyID := fmt.Sprintf("%v", interfaceServiceKeyID)

	paymentPartner, err := models.GetPaymentPartnerByAPIKeyID(serviceKeyID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Server is busy")
		return
	}

	ticket, err := models.GetTicketByID(*paymentReq.TicketID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Server is busy")
		return
	}

	payment, err := models.CreatePayment(ticket.Value,
		&ticket.ID,
		&paymentPartner.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Server is busy")
		return
	}

	newTicket := ticket
	newTicket.Type = models.PaidTicket
	err = ticket.UpdateWithNewTicketValue(newTicket)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Server is busy")
		return
	}

	c.JSON(200, payment)
}
