package main

import (
	"cinema-admin/app"
	"cinema-admin/db"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	//Connect to database
	db.Init()
	defer db.CloseDB()

	// Init for app to work
	app.Init()
}
