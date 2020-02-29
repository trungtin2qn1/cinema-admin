package consumer

import (
	"cinema-admin/models"
	"cinema-admin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login router handler for frontend
func Login(c *gin.Context) {
	authReq := models.AuthReq{}
	err := c.ShouldBind(&authReq)
	if err != nil {
		go utils.LogErrToFile(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data or data type is invalid",
		})
		return
	}

	consumer, err := models.Login(authReq.Email, authReq.Password)

	if err != nil {
		go utils.LogErrToFile(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server is busy",
		})
		return
	}

	c.JSON(200, consumer)
}
