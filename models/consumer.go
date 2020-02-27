package models

import (
	"cinema-admin/db"
	"errors"
	"time"
)

//Consumer ...
type Consumer struct {
	ID        string     `json:"id,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
	Email     string     `json:"email,omitempty"`
	Name      string     `json:"name,omitempty"`
	Phone     string     `json:"phone,omitempty"`
	Address   string     `json:"address,omitempty"`
	Password  string     `json:"-" form:"-"`
	Token     string     `json:"token,omitempty" gorm:"-"`
}

// GetConsumerByEmail ...
func GetConsumerByEmail(email string) (Consumer, error) {
	dbConn := db.GetDB()

	consumer := Consumer{}
	res := dbConn.Where("email = ?", email).First(&consumer)
	if res.Error != nil {
		return consumer, errors.New("Data or data type is invalid")
	}

	return consumer, nil
}

// GetConsumerByID ...
func GetConsumerByID(id string) (Consumer, error) {
	dbConn := db.GetDB()

	consumer := Consumer{}
	res := dbConn.Where("id = ?", id).Find(&consumer)
	if res.Error != nil {
		return consumer, errors.New("Data or data type is invalid")
	}
	return consumer, nil
}
