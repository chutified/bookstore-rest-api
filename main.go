package main

import (
	"tommychu/workdir/027_api-example-v2/app"
	"tommychu/workdir/027_api-example-v2/config"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// @title Bookstore API example with Gin
// @version 1.0
// @description This is a sample of a Gin API framework.

// @contact.name Tommy Chu
// @contact.email tommychu2256@gmail.com

// @schemes http
// @host localhost:8081
// @BasePath /api/v1
func main() {
	cfg := config.GetConfig()

	// set app
	a := app.New()
	a.Initialize(cfg)
	defer a.Close()

	a.Run()
}
