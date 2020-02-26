package adminuser

import (
	"cinema-admin/models"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/qor/admin"
	"github.com/qor/qor"
)

func setupSearchHandler(resource *admin.Resource) {
	resource.SearchHandler = searchHandler
}

func searchHandler(keyword string, context *qor.Context) *gorm.DB {
	if keyword == "" {
		return context.GetDB()
	}
	adminUser := models.AdminUser{}
	_, err := strconv.ParseInt(keyword, 10, 64)
	if err == nil {
		db := models.SearchHandlerInteger(context.GetDB(), &adminUser, keyword, "id")
		context.SetDB(db)
		return db
	}
	db := models.SearchHandlerString(context.GetDB(), &adminUser, keyword, "username",
		"first_name", "last_name", "email")
	context.SetDB(db)
	return db
}
