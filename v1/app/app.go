package app

import (
	"fmt"
	"log"
	"net/http"
	"tommychu/workdir/026_api-example/app/services/dbservice"
	"tommychu/workdir/026_api-example/config"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// App contains everything needed
// to run the API properly
type App struct {
	Port   int
	Logger *log.Logger
	Router *mux.Router
	DB     *gorm.DB
}

// NewApp returns new App
func NewApp() *App {
	return &App{}
}

// Initialize  applies the configuration and runs the API
func (a *App) Initialize(config *config.Config) {

	// set logging config
	a.Logger = config.Log.Logger

	// set PORT
	a.Port = config.Port

	// set database config
	dbURI := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		config.DB.Host,
		config.DB.Port,
		config.DB.User,
		config.DB.DBname,
		config.DB.Password,
	)

	// db connection
	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		a.Logger.Fatal(err)
	}
	a.Logger.Println("Succesfully connected to DB.")
	a.Logger.Println("--------------------")
	a.DB = dbservice.DBMigrate(db)

	// router
	a.Router = mux.NewRouter()
	a.setRouter()
}

// Run starts the server
func (a *App) Run(port string) {
	err := http.ListenAndServe(port, a.Router)
	if err != nil {
		a.Logger.Fatal(err)
	}
}
