package services

import (
	"cinema-admin/controllers/movie"
	"cinema-admin/models"

	"github.com/qor/admin"
)

//SetupMovieManagementResource ...
func SetupMovieManagementResource(Admin *admin.Admin) {
	menu := Admin.AddMenu(&admin.Menu{Name: "Quản lý phim"})

	// Add Resource
	movieResource := Admin.AddResource(&models.Movie{}, &admin.Config{
		Menu: []string{menu.Name},
	})

	// Setup Resource

	movie.SetupResource(movieResource)
}
