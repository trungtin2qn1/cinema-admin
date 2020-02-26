package models

import "time"

//Payment ...
type Payment struct {
	ID        string `json:"id,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Amount    int     `json:"amount,omitempty"`
	TicketID  *string `json:"ticket_id,omitempty"`
}
