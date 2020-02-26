package models

import "time"

//Ticket ...
type Ticket struct {
	ID                      string `json:"id,omitempty"`
	CreatedAt               time.Time
	UpdatedAt               time.Time
	DeletedAt               *time.Time
	Value                   int     `json:"value,omitempty"`
	NumberSeat              string  `json:"number_seat,omitempty"`
	NumberTheater           string  `json:"number_theater,omitempty"`
	MovieSessionInTheaterID *string `json:"movie_session_in_theater_id,omitempty"`
	CustomerID              *string `json:"customer_id,omitempty"`
	Type                    int     `json:"type,omitempty"`
}
