package services

import (
	"cinema-admin/controllers/ticket"
	"cinema-admin/models"

	"github.com/qor/admin"
)

//SetupTicketManagementResource ...
func SetupTicketManagementResource(Admin *admin.Admin) {
	menu := Admin.AddMenu(&admin.Menu{Name: "Quản lý Vé"})

	// Add Resource
	ticketUserResource := Admin.AddResource(&models.Ticket{}, &admin.Config{
		Menu: []string{menu.Name},
	})

	// Setup Resource

	ticket.SetupResource(ticketUserResource)

}
