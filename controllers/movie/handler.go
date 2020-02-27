package movie

import (
	"cinema-admin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetMovieInfoByID ...
func GetMovieInfoByID(c *gin.Context) {
	movieID := c.Param("movie_id")

	movie, err := models.GetMovieByID(movieID)

	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Can't find theater")
		return
	}

	display := Display{Movie: movie}

	c.JSON(200, display)
}
