package models

import (
	"cinema-admin/db"
	"errors"
	"time"
)

// Theater ...
type Theater struct {
	ID          string     `json:"id,omitempty"`
	CreatedAt   time.Time  `json:"created_at,omitempty"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `json:"-"`
	Name        string     `json:"name,omitempty"`
	Description string     `json:"description,omitempty"`
	State       string     `json:"state,omitempty"`
	City        string     `json:"city,omitempty"`
	District    string     `json:"district,omitempty"`
	Ward        string     `json:"ward,omitempty"`
	Street      string     `json:"street,omitempty"`
}

// GetTheaterByID ...
func GetTheaterByID(id string) (Theater, error) {
	dbConn := db.GetDB()

	theater := Theater{}
	res := dbConn.Where("id = ?", id).Find(&theater)
	if res.Error != nil {
		return theater, errors.New("Data or data type is invalid")
	}
	return theater, nil
}
