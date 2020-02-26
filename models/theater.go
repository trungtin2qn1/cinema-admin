package models

import "time"

// Theater ...
type Theater struct {
	ID          string `json:"id,omitempty"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	State       string `json:"state,omitempty"`
	City        string `json:"city,omitempty"`
	District    string `json:"district,omitempty"`
	Ward        string `json:"ward,omitempty"`
	Street      string `json:"street,omitempty"`
}
