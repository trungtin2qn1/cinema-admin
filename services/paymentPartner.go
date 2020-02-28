package services

import (
	"cinema-admin/controllers/paymentpartner"
	"cinema-admin/models"

	"github.com/qor/admin"
)

//SetupPaymentPartnerManagementResource ...
func SetupPaymentPartnerManagementResource(Admin *admin.Admin) {
	menu := Admin.AddMenu(&admin.Menu{Name: "Quản lý Đối Tác Thanh Toán"})

	// Add Resource
	paymentUserResource := Admin.AddResource(&models.PaymentPartner{}, &admin.Config{
		Menu: []string{menu.Name},
	})

	// Setup Resource

	paymentpartner.SetupResource(paymentUserResource)

}
