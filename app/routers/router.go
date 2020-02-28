package routers

import (
	"cinema-admin/admin/bindatafs"
	controllerAuth "cinema-admin/controllers/auth"
	controllerConsumer "cinema-admin/controllers/consumer"
	controllerMovie "cinema-admin/controllers/movie"
	controllerMovieSession "cinema-admin/controllers/moviesession"
	controllerPayment "cinema-admin/controllers/payment"
	controllerTheater "cinema-admin/controllers/theater"
	controllerTicket "cinema-admin/controllers/ticket"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// SetupRouter ...
func SetupRouter(mux *http.ServeMux) *gin.Engine {
	router := gin.Default()

	// Set login view
	router = setViewLogin(router)

	router.GET("/", controllerAuth.GetDefault)
	router.GET("/login", controllerAuth.GetLogin)
	router.GET("/logout", controllerAuth.GetLogout)
	router.POST("/login", controllerAuth.PostLogin)

	adm := router.Group("/admin")
	adm.Use(controllerAuth.CheckCookie)
	adm.Any("/*resource", gin.WrapH(mux))

	api := router.Group("/apis")
	{
		api.POST("/login", controllerConsumer.Login)

		auth := api.Group("/auth")
		{
			auth.Use(controllerAuth.VerifyJWTToken)
			auth.GET("/theater/:theater_id", controllerTheater.GetTheaterInfoByID)
			auth.GET("/movie/:movie_id", controllerMovie.GetMovieInfoByID)
			auth.GET("/theater/:theater_id/movie-sessions",
				controllerMovieSession.GetMovieSessionsByTheaterID)
			auth.POST("/ticket", controllerTicket.BookTicket)
		}

		public := api.Group("/public")
		{
			public.Use(controllerAuth.CheckAPIKey)
			public.GET("/theater/:theater_id", controllerTheater.GetTheaterInfoByID)
			public.GET("/movie/:theater_id", controllerMovie.GetMovieInfoByID)
			public.GET("/theater/:theater_id/movie-sessions",
				controllerMovieSession.GetMovieSessionsByTheaterID)
			public.POST("/ticket", controllerTicket.BookTicket)
			public.POST("/check-out", controllerPayment.CheckOut)
		}
	}

	return router
}

// setViewLogin ...
func setViewLogin(r *gin.Engine) *gin.Engine {
	lfs := bindatafs.AssetFS.NameSpace("login")
	lfs.RegisterPath("./template/")
	logintpl, err := lfs.Asset("login.html")
	if err != nil {
		logrus.WithError(err).Fatal("Unable to find HTML template for login page in admin")
	}
	r.SetHTMLTemplate(template.Must(template.New("login.html").Parse(string(logintpl))))
	return r
}
