package models

import (
	"cinema-admin/db"
	"cinema-admin/utils"
	"errors"
	"time"
)

//PaymentPartner ...
type PaymentPartner struct {
	ID        string     `json:"id,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
	Name      string     `json:"name,omitempty"`
	Type      int        `json:"type,omitempty"`
	APIKeyID  *string    `json:"api_key_id,omitempty"`
}

// GetPaymentPartnerByID ...
func GetPaymentPartnerByID(id string) (PaymentPartner, error) {
	dbConn := db.GetDB()

	paymentPartner := PaymentPartner{}

	res := dbConn.Where("id = ?", id).Find(&paymentPartner)
	if res.Error != nil {
		go utils.LogErrToFile(res.Error.Error())
		return paymentPartner, errors.New("Data or data type is invalid")
	}
	return paymentPartner, nil
}

// GetPaymentPartnerByAPIKeyID ...
func GetPaymentPartnerByAPIKeyID(apiKeyID string) (PaymentPartner, error) {
	dbConn := db.GetDB()

	paymentPartner := PaymentPartner{}

	res := dbConn.Where("api_key_id = ?", apiKeyID).Find(&paymentPartner)
	if res.Error != nil {
		go utils.LogErrToFile(res.Error.Error())
		return paymentPartner, errors.New("Data or data type is invalid")
	}
	return paymentPartner, nil
}
