package app

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"tommychu/workdir/026_api-example-v2/app/handlers"
	"tommychu/workdir/026_api-example-v2/app/services/dbservices"
	"tommychu/workdir/026_api-example-v2/config"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// App is an application struct.
type App struct {
	Log    io.Writer
	Srv    *http.Server
	Router *gin.Engine
	DB     *gorm.DB
}

// New returns a new App.
func New() *App {
	return &App{}
}

// Initialize applies the config of the App.
func (a *App) Initialize(cfg *config.Config) {

	// log
	a.Log = cfg.Log.Output

	// database
	db := dbservices.GetDB(cfg)
	a.DB = dbservices.DBMigrate(db)

	// router
	a.Router = handlers.GetRouter(cfg, a.DB)

	// server
	a.Srv = &http.Server{
		Addr:              cfg.Srv.Addr,
		Handler:           a.Router, // router
		ReadTimeout:       cfg.Srv.ReadTimeout,
		ReadHeaderTimeout: cfg.Srv.ReadHeaderTimeout,
		WriteTimeout:      cfg.Srv.WriteTimeout,
		IdleTimeout:       cfg.Srv.IdleTimeout,
		MaxHeaderBytes:    cfg.Srv.MaxHeaderBytes,
	}
}

// Close takes care of the whole application closure.
func (a *App) Close() []error {
	return []error{
		a.DB.Close(),
	}
}

// Run starts the application server.
func (a *App) Run() {
	fmt.Fprintf(a.Log, "Listening and serving HTTP on %s\n", a.Srv.Addr)
	log.Fatal(a.Srv.ListenAndServe())
}
