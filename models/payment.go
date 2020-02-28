package models

import (
	"cinema-admin/db"
	"time"
)

//Payment ...
type Payment struct {
	ID               string     `json:"id,omitempty"`
	CreatedAt        time.Time  `json:"created_at,omitempty"`
	UpdatedAt        time.Time  `json:"-"`
	DeletedAt        *time.Time `json:"-"`
	Amount           int        `json:"amount,omitempty"`
	TicketID         *string    `json:"ticket_id,omitempty"`
	PaymentPartnerID *string    `json:"payment_partner_id,omitempty"`
}

//CreatePayment ...
func CreatePayment(amount int, ticketID, paymentPartnerID *string) (*Payment, error) {
	dbConn := db.GetDB()
	payment := &Payment{
		Amount:           amount,
		TicketID:         ticketID,
		PaymentPartnerID: paymentPartnerID,
	}
	dbConn = dbConn.Create(payment)
	return payment, dbConn.Error
}
