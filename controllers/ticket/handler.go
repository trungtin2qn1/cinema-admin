package ticket

import (
	"cinema-admin/models"
	"cinema-admin/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//BookTicket ...
func BookTicket(c *gin.Context) {
	interfaceUserID, _ := c.Get("user_id")
	userID := fmt.Sprintf("%v", interfaceUserID)

	var consumerID *string

	if interfaceUserID != nil {
		consumerID = new(string)
		*consumerID = userID
	} else {
		consumerID = nil
	}

	ticketReq := models.Ticket{}

	err := c.ShouldBind(&ticketReq)

	if err != nil {
		go utils.LogErrToFile(err.Error())
		c.JSON(http.StatusBadRequest, "Data or data type is invalid")
		return
	}

	ticket, err := models.CreateTicket(ticketReq.NumberSeat,
		ticketReq.MovieSessionInTheaterID, consumerID,
		ticketReq.NumberTheater, models.UnpaidTicket, ticketReq.Value)

	c.JSON(200, ticket)
}
