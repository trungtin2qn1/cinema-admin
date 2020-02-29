package theater

import (
	"cinema-admin/models"
	"cinema-admin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetTheaterInfoByID ...
func GetTheaterInfoByID(c *gin.Context) {
	theaterID := c.Param("theater_id")

	theater, err := models.GetTheaterByID(theaterID)

	if err != nil {
		go utils.LogErrToFile(err.Error())
		c.JSON(http.StatusNotAcceptable, "Can't find theater")
		return
	}

	display := Display{Theater: theater}

	c.JSON(200, display)
}
