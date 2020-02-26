package models

import "time"

//Product ...
type Product struct {
	//ID        uint64 `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	ID        string `json:"id,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	Name  string `json:"name,omitempty"`
	Price int    `json:"price,omitempty"`
}
