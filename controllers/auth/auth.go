package auth

import (
	"cinema-admin/models"
	"cinema-admin/utils"
	"fmt"
	"log"

	"net/http"

	"github.com/qor/qor"

	"github.com/gin-gonic/gin"
	"github.com/qor/admin"
)

// Authentication ...
type Authentication struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

// LoginURL ...
func (Authentication) LoginURL(context *admin.Context) string {
	return "/login"
}

// LogoutURL ...
func (Authentication) LogoutURL(context *admin.Context) string {
	return "/logout"
}

// GetCurrentUser ...
func (Authentication) GetCurrentUser(context *admin.Context) qor.CurrentUser {
	cookie, err := context.Request.Cookie("admin_id")
	if err != nil || cookie == nil || cookie.Value == "" {
		return nil
	}

	log.Println("cookie:", cookie)

	admin, err := models.GetAdminByID(cookie.Value)
	if err != nil {
		return nil
	}
	return admin
}

// GetDefault ...
func GetDefault(c *gin.Context) {
	c.Redirect(http.StatusSeeOther, "/login")
}

// GetLogin ...
func GetLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

// GetLogout ...
func GetLogout(c *gin.Context) {
	DeleteCookie(c)
	c.Redirect(http.StatusSeeOther, "/login")
}

// PostLogin ...
func PostLogin(c *gin.Context) {
	auth := Authentication{}

	err := c.ShouldBind(&auth)
	if err != nil {
		log.Println("err:", err)
	}

	if auth.Email == "" || auth.Password == "" {
		fmt.Println("auth.Email:", auth.Email)
		fmt.Println("auth.Password:", auth.Password)
		c.JSON(http.StatusNotAcceptable, "Data or data type is invalid")
		return
	}

	log.Println("auth.Email:", auth.Email)
	admin, err := models.GetAdminByEmail(auth.Email)
	log.Println("admin:", admin)

	if err != nil {
		fmt.Println(auth.Email)
		c.JSON(http.StatusNotAcceptable, err.Error())
		return
	}
	if check, _ := utils.Compare(admin.Password, auth.Password); check == false {
		c.JSON(http.StatusNotAcceptable, "Email or password is invalid")
		return
	}

	SetCookie(c, admin.ID)

	c.Redirect(http.StatusSeeOther, "/admin")
}
