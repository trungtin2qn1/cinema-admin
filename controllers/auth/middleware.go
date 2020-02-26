package controller

import (
	"cinema-admin/models"
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
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	_, err = models.GetAdminByID(cookie.Value)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	c.Next()
}

// DeleteCookie ...
func DeleteCookie(c *gin.Context) {
	cookie, err := c.Request.Cookie("admin_id")
	if err != nil || cookie == nil {
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
			"message": "Token can not be null",
		})
		return
	}
	rawToken := string(token[len("Tin "):])
	userID, _, err := jwt.VerificationToken(rawToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Token is in valid",
		})
		return
	}
	fmt.Println(userID)
	c.Set("user_id", userID)
	c.Next()
}

//CheckAPIKey ...
func CheckAPIKey(c *gin.Context) {
	c.Next()
}
