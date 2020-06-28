package main

import (
	"tommychu/workdir/026_api-example/app"
	"tommychu/workdir/026_api-example/config"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// @Title Library Example API
// @Version 1.0
// @Description This is a sample of a REST API

// @Contact.name Tommy Chu
// @Contact.email tommychu2256@gmail.com

// @Host localhost:8081
// @BasePath /
// @Schemes http

func main() {
	cfg := config.GetConfig()

	app := app.NewApp()
	defer app.DB.Close()

	app.Initialize(cfg)
	app.Run(":8081")
}
