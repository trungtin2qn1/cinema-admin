package services

import (
	"cinema-admin/controllers/consumer"
	"cinema-admin/models"

	"github.com/qor/admin"
)

//SetupConsumerManagementResource ...
func SetupConsumerManagementResource(Admin *admin.Admin) {
	menu := Admin.AddMenu(&admin.Menu{Name: "Quản lý Người Tiêu Dùng"})

	// Add Resource
	consumerUserResource := Admin.AddResource(&models.Consumer{}, &admin.Config{
		Menu: []string{menu.Name},
	})

	// Setup Resource

	consumer.SetupResource(consumerUserResource)

}
