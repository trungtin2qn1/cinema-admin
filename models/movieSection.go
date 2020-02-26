package models

import "time"

//MovieSessionInTheater ...
type MovieSessionInTheater struct {
	ID        string `json:"id,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	StartTime *time.Time `json:"start_time,omitempty"`
	EndTime   *time.Time `json:"end_time,omitempty"`
	MovieID   *string    `json:"movie_id,omitempty"`
	TheaterID *string    `json:"theater_id,omitempty"`
}
