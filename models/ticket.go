package models

import (
	"cinema-admin/db"
	"time"
)

//Ticket ...
type Ticket struct {
	ID                      string     `json:"id,omitempty"`
	CreatedAt               time.Time  `json:"created_at,omitempty"`
	UpdatedAt               time.Time  `json:"-"`
	DeletedAt               *time.Time `json:"-"`
	Value                   int        `json:"value,omitempty"`
	NumberSeat              string     `json:"number_seat,omitempty"`
	NumberTheater           int        `json:"number_theater,omitempty"`
	MovieSessionInTheaterID *string    `json:"movie_session_in_theater_id,omitempty"`
	CustomerID              *string    `json:"customer_id,omitempty"`
	Type                    int        `json:"type,omitempty"`
}

const (
	PaidTicket   = 1 //Tickets have been paid
	UnpaidTicket = 2 //Tickets have not been paid
)

//CreateTicket ...
func CreateTicket(numberSeat string,
	movieSessionInTheaterID, customerID *string,
	numberTheater, t, value int) (*Ticket, error) {
	dbConn := db.GetDB()
	ticket := &Ticket{
		Value:                   value,
		Type:                    t,
		NumberTheater:           numberTheater,
		NumberSeat:              numberSeat,
		MovieSessionInTheaterID: movieSessionInTheaterID,
		CustomerID:              customerID,
	}
	dbConn = dbConn.Create(ticket)
	return ticket, dbConn.Error
}
