package consumer

import (
	"cinema-admin/models"
	"cinema-admin/utils"
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
	consumer := models.Consumer{}
	_, err := strconv.ParseInt(keyword, 10, 64)
	if err == nil {
		go utils.LogErrToFile(err.Error())
		db := models.SearchHandlerInteger(context.GetDB(), &consumer, keyword, "id")
		context.SetDB(db)
		return db
	}
	db := models.SearchHandlerString(context.GetDB(), &consumer, keyword, "name")
	context.SetDB(db)
	return db
}
