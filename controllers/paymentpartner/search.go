package paymentpartner

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
	paymentPartner := models.PaymentPartner{}
	_, err := strconv.ParseInt(keyword, 10, 64)
	if err == nil {
		go utils.LogErrToFile(err.Error())
		db := models.SearchHandlerInteger(context.GetDB(), &paymentPartner, keyword, "id")
		context.SetDB(db)
		return db
	}
	db := models.SearchHandlerString(context.GetDB(), &paymentPartner, keyword, "name")
	context.SetDB(db)
	return db
}
