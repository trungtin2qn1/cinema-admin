package models

import "time"

//Movie ...
type Movie struct {
	ID          string `json:"id,omitempty"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
	Name        string     `json:"name,omitempty"`
	Description string     `json:"description,omitempty"`
	Image       string     `json:"image,omitempty"`
	Trailer     string     `json:"trailer,omitempty"`
	StartedAt   *time.Time `json:"started_at,omitempty"`
	Duration    int        `json:"duration,omitempty"`
	Rating      float32    `json:"rating,omitempty"`
	Type        int        `json:"type,omitempty"`
}
