package services

import (
	"cinema-admin/controllers/payment"
	"cinema-admin/models"

	"github.com/qor/admin"
)

//SetupPaymentManagementResource ...
func SetupPaymentManagementResource(Admin *admin.Admin) {
	menu := Admin.AddMenu(&admin.Menu{Name: "Quản lý Thanh Toán"})

	// Add Resource
	paymentUserResource := Admin.AddResource(&models.Payment{}, &admin.Config{
		Menu: []string{menu.Name},
	})

	// Setup Resource

	payment.SetupResource(paymentUserResource)

}
