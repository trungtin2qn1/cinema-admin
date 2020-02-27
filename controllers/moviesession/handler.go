package moviesession

import (
	"cinema-admin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetMovieSessionsByTheaterID ...
func GetMovieSessionsByTheaterID(c *gin.Context) {
	theaterID := c.Param("theater_id")

	date := c.Query("date")

	movies, err := models.GetMoviesByTheaterID(theaterID, date)

	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Can't find movies by theater id",
		})
		return
	}

	displays := []Display{}
	for _, movie := range movies {
		display := Display{
			Movie: movie,
		}

		movieSessions, err := models.GetMovieSessionsByTheaterIDAndMovieID(theaterID, movie.ID, date)

		if err != nil {
			continue
		}

		sessions := []Session{}

		for _, movieSession := range movieSessions {
			session := Session{
				TimeStart: movieSession.StartTime,
				TimeEnd:   movieSession.EndTime,
			}
			sessions = append(sessions, session)
		}

		display.Sessions = sessions

		displays = append(displays, display)
	}

	c.JSON(200, displays)
}
