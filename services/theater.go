package services

import (
	"cinema-admin/controllers/theater"
	"cinema-admin/models"

	"github.com/qor/admin"
)

//SetupTheaterManagementResource ...
func SetupTheaterManagementResource(Admin *admin.Admin) {
	menu := Admin.AddMenu(&admin.Menu{Name: "Quản lý Rạp Chiếu Phim Lớn"})

	// Add Resource
	theaterUserResource := Admin.AddResource(&models.Theater{}, &admin.Config{
		Menu: []string{menu.Name},
	})

	// Setup Resource

	theater.SetupResource(theaterUserResource)

}
