package routers

import (
	"cinema-admin/admin/bindatafs"
	"cinema-admin/controllers/auth"
	"cinema-admin/db"
	"cinema-admin/models"
	"cinema-admin/services"
	"net/http"

	"github.com/qor/admin"
	"github.com/qor/roles"
)

// SetupAdmin ...
func SetupAdmin() (*admin.Admin, *http.ServeMux) {
	dbConn := db.GetDB()
	Admin := admin.New(&admin.AdminConfig{
		SiteName: "Cinema Admin",
		DB:       dbConn,
		Auth:     &auth.Authentication{},
	})

	setupRoles()

	// Config router
	mux := http.NewServeMux()

	services.SetupAdminUserManagementResource(Admin)
	services.SetupTheaterManagementResource(Admin)
	services.SetupMovieManagementResource(Admin)
	services.SetupMovieSessionInTheaterManagementResource(Admin)
	services.SetupTicketManagementResource(Admin)
	services.SetupConsumerManagementResource(Admin)
	services.SetupPaymentManagementResource(Admin)

	// Setup theme
	assetFS := bindatafs.AssetFS
	Admin.SetAssetFS(assetFS.NameSpace("admin"))

	// Mount admin interface to mux
	Admin.MountTo("/admin", mux)
	return Admin, mux
}

// setupRoles ...
func setupRoles() {
	roles.Register(models.Role.ROOT, func(req *http.Request, currentUser interface{}) bool {
		if admin, ok := currentUser.(models.AdminUser); ok {
			if admin.Role == models.Role.ROOT {
				return true
			}
		}
		return false
	})
	roles.Register(models.Role.COMMUNITY, func(req *http.Request, currentUser interface{}) bool {
		if admin, ok := currentUser.(models.AdminUser); ok {
			if admin.Role == models.Role.COMMUNITY {
				return true
			}
		}
		return false
	})
	roles.Register(models.Role.BUSINESSMANAGER, func(req *http.Request, currentUser interface{}) bool {
		if admin, ok := currentUser.(models.AdminUser); ok {
			if admin.Role == models.Role.BUSINESSMANAGER {
				return true
			}
		}
		return false
	})
}
