package moviesession

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
	movieSessionInTheater := models.MovieSessionInTheater{}
	_, err := strconv.ParseInt(keyword, 10, 64)
	if err == nil {
		db := models.SearchHandlerInteger(context.GetDB(), &movieSessionInTheater, keyword, "id")
		context.SetDB(db)
		return db
	}
	db := models.SearchHandlerString(context.GetDB(), &movieSessionInTheater, keyword, "name")
	context.SetDB(db)
	return db
}
