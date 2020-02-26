package models

import "github.com/jinzhu/gorm"

// SearchHandlerInteger ...
func SearchHandlerInteger(db *gorm.DB, table interface{}, keyword string, fields ...string) *gorm.DB {
	db = db.Where(fields[0] + " = " + keyword)
	for i := 1; i < len(fields); i++ {
		db = db.Or(fields[i] + " = " + keyword)
	}
	return db
}

// SearchHandlerString ...
func SearchHandlerString(db *gorm.DB, table interface{}, keyword string, fields ...string) *gorm.DB {
	db = db.Where("upper(" + fields[0] + ")" + " like upper('%" + keyword + "%')")
	for i := 1; i < len(fields); i++ {
		db = db.Or("upper(" + fields[i] + ")" + " like upper('%" + keyword + "%')")
	}
	return db
}
