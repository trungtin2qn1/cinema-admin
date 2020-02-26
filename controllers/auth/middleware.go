package controller

import (
	"cinema-admin/models"
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

//CheckAuthenticationToken ...
func CheckAuthenticationToken(c *gin.Context) {
	c.Next()
}

//CheckAPIKey ...
func CheckAPIKey(c *gin.Context) {
	c.Next()
}
