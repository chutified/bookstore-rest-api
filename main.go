package main

import (
	"tommychu/workdir/027_api-example-v2/app"
	"tommychu/workdir/027_api-example-v2/config"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	cfg := config.GetConfig()
	a := app.New()
	a.Initialize(cfg)
	defer a.Close()
	a.Run()
}
