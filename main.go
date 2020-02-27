package main

import (
	"cinema-admin/app"
	"cinema-admin/db"
	"cinema-admin/utils/jwt"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	// Connect to database
	db.Init()
	defer db.CloseDB()

	// Load keys for jwt
	jwt.LoadRSAKeys()

	// Init for app to work
	app.Init()
}
