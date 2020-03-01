package main

import (
	"cinema-admin/db"
	"cinema-admin/migration/models"
	"cinema-admin/utils"
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/gormigrate.v1"
)

func main() {
	//Connect to database
	db.Init(nil)
	defer db.CloseDB()

	dbConn := db.GetDB()

	m := gormigrate.New(dbConn, gormigrate.DefaultOptions, []*gormigrate.Migration{
		models.InitAdmin,
	})

	if err := m.Migrate(); err != nil {
		go utils.LogErrToFile(err.Error())
	}
	log.Printf("Migration did run successfully")

}
