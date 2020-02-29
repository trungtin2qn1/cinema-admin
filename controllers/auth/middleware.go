package auth

import (
	"cinema-admin/models"
	"cinema-admin/utils"
	"cinema-admin/utils/jwt"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetCookie ...
func SetCookie(c *gin.Context, adminID string) {
	cookie := http.Cookie{
		Name:   "admin_id",
		Value:  adminID,
		MaxAge: 24 * 60 * 60,
	}
	http.SetCookie(c.Writer, &cookie)
}

// CheckCookie ...
func CheckCookie(c *gin.Context) {
	cookie, err := c.Request.Cookie("admin_id")
	if err != nil || cookie == nil || cookie.Value == "" {
		go utils.LogErrToFile(err.Error())
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	_, err = models.GetAdminByID(cookie.Value)
	if err != nil {
		go utils.LogErrToFile(err.Error())
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	c.Next()
}

// DeleteCookie ...
func DeleteCookie(c *gin.Context) {
	cookie, err := c.Request.Cookie("admin_id")
	if err != nil || cookie == nil {
		go utils.LogErrToFile(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	cookie.MaxAge = -1
	http.SetCookie(c.Writer, cookie)
}

//VerifyJWTToken ...
func VerifyJWTToken(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	rawToken := string(token[len("Tin "):])
	userID, _, err := jwt.VerificationToken(rawToken)
	if err != nil {
		go utils.LogErrToFile(err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	fmt.Println(userID)
	c.Set("user_id", userID)
	c.Next()
}

//CheckAPIKey ...
func CheckAPIKey(c *gin.Context) {
	key := c.Request.Header.Get("api-key")

	if key == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	serviceKey, err := models.GetAPIKeyByKey(key)
	if err != nil {
		go utils.LogErrToFile(err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	c.Set("service-key-id", serviceKey.ID)
	c.Next()
}
