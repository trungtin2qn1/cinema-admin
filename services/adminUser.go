package services

import (
	"cinema-admin/controllers/adminuser"
	"cinema-admin/models"

	"github.com/qor/admin"
)

// SetupAdminUserManagementResource ...
func SetupAdminUserManagementResource(Admin *admin.Admin) {
	menu := Admin.AddMenu(&admin.Menu{Name: "Quản lý Admin"})

	// Add Resource
	adminUserResource := Admin.AddResource(&models.AdminUser{}, &admin.Config{
		Menu: []string{menu.Name},
	})

	// Setup Resource

	adminuser.SetupResource(adminUserResource)
}
