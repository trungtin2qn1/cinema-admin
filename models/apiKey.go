package models

import (
	"cinema-admin/db"
	"errors"
	"time"
)

//APIKey ...
type APIKey struct {
	ID        string     `json:"id,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
	Value     string     `json:"value,omitempty"`
	Type      string     `json:"type,omitempty"`
}

// GetAPIKeyByKey ...
func GetAPIKeyByKey(key string) (APIKey, error) {
	dbConn := db.GetDB()

	apiKey := APIKey{}
	res := dbConn.Where("value = ?", key).Find(&apiKey)
	if res.Error != nil {
		return apiKey, errors.New("Data or data type is invalid")
	}

	return apiKey, nil
}
