package models

import (
	"cinema-admin/db"
	"time"
)

//Payment ...
type Payment struct {
	ID        string     `json:"id,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
	Amount    int        `json:"amount,omitempty"`
	TicketID  *string    `json:"ticket_id,omitempty"`
}

//CreatePayment ...
func CreatePayment(amount int, ticketID *string) (*Payment, error) {
	dbConn := db.GetDB()
	payment := &Payment{
		Amount:   amount,
		TicketID: ticketID,
	}
	dbConn = dbConn.Create(payment)
	return payment, dbConn.Error
}
