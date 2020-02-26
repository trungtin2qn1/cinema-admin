package app

import (
	"cinema-admin/app/routers"
	"log"
)

// Init ...
func Init() {
	_, mux := routers.SetupAdmin()
	router := routers.SetupRouter(mux)

	err := router.Run(":4000")
	if err != nil {
		log.Fatal(err)
	}
}
