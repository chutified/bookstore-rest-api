package main

import (
	"tommychu/workdir/027_api-example-v2/app"
	"tommychu/workdir/027_api-example-v2/config"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// @Title Bookstore API example with Gin
// @Version 1.0
// @Description This is a sample of a Gin API framework.

// @Contact.name Tommy Chu
// @Contact.email tommychu2256@gmail.com

// @Schemes http
// @Host localhost:8081
// @BasePath /api/v1
func main() {
	cfg := config.GetConfig()

	// set app
	a := app.New()
	a.Initialize(cfg)
	defer a.Close()

	a.Run()
}
