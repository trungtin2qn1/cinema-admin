package services

import (
	"cinema-admin/controllers/moviesession"
	"cinema-admin/models"

	"github.com/qor/admin"
)

//SetupMovieSessionInTheaterManagementResource ...
func SetupMovieSessionInTheaterManagementResource(Admin *admin.Admin) {
	menu := Admin.AddMenu(&admin.Menu{Name: "Quản lý các suất chiếu phim"})

	// Add Resource
	movieSessionResource := Admin.AddResource(&models.MovieSessionInTheater{}, &admin.Config{
		Menu: []string{menu.Name},
	})

	// Setup Resource

	moviesession.SetupResource(movieSessionResource)
}
