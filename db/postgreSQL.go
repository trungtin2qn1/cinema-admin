package db

import (
	"cinema-admin/utils"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

//Init ...
func Init() {
	dbUser := utils.GetEnv("POSTGREST_DB_USER", "user")
	dbPassword := utils.GetEnv("POSTGREST_DB_PASSWORD", "123456")
	dbName := utils.GetEnv("POSTGREST_DB_NAME", "cinema-admin")
	dbHost := utils.GetEnv("POSTGREST_DB_HOST", "localhost")
	dbPort := utils.GetEnv("POSTGREST_DB_PORT", "5432")

	port, e := utils.ConvertStringToInt(dbPort)
	if e != nil {
		log.Fatal(e)
	}
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, port, dbUser, dbPassword, dbName)

	var err error
	fmt.Println(dbinfo)
	db, err = ConnectDB(dbinfo)
	if err != nil {
		fmt.Println(1)
		log.Fatal(err)
	}
	fmt.Println("Success")
}

//ConnectDB ...
func ConnectDB(dbinfo string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	db = db.Debug()
	return db, nil
}

//GetDB ...
func GetDB() *gorm.DB {
	return db
}

//CloseDB before close application
func CloseDB() {
	db.Close()
	db = nil
}
