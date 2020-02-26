package models

import "time"

//Consumer ...
type Consumer struct {
	ID        string `json:"id,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Email     string `json:"email,omitempty"`
	Name      string `json:"name,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Address   string `json:"address,omitempty"`
}
