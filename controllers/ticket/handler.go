package ticket

import (
	"cinema-admin/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//BookTicketAuth ...
func BookTicketAuth(c *gin.Context) {

	interfaceUserID, _ := c.Get("user_id")
	consumerID := fmt.Sprintf("%v", interfaceUserID)

	ticketReq := models.Ticket{}

	err := c.ShouldBind(&ticketReq)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Data or data type is invalid")
		return
	}

	ticket, err := models.CreateTicket(ticketReq.NumberSeat,
		ticketReq.MovieSessionInTheaterID, &consumerID,
		ticketReq.NumberTheater, models.UnpaidTicket, ticketReq.Value)

	c.JSON(200, ticket)
}

//BookTicketService ...
func BookTicketService(c *gin.Context) {
	ticketReq := models.Ticket{}

	err := c.ShouldBind(&ticketReq)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Data or data type is invalid")
		return
	}

	ticket, err := models.CreateTicket(ticketReq.NumberSeat,
		ticketReq.MovieSessionInTheaterID, ticketReq.CustomerID,
		ticketReq.NumberTheater, models.PaidTicket, ticketReq.Value)

	c.JSON(200, ticket)
}
