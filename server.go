package main

import (
	"tommychu/workdir/026_api-example/app"
	"tommychu/workdir/026_api-example/config"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	cfg := config.GetConfig()

	app := app.NewApp()
	defer app.DB.Close()

	app.Initialize(cfg)
	app.Run(":8080")
}
