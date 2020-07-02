package main

import (
	"log"
	"tommychu/workdir/026_api-example-v2/app"
	"tommychu/workdir/026_api-example-v2/config"

	"github.com/gin-gonic/gin"
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
	if !config.DEBUG_MODE {
		gin.SetMode(gin.ReleaseMode)
	}

	cfg := config.GetConfig()

	// set app
	a := app.New()
	err := a.Initialize(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer log.Fatal(a.Close())

	log.Panic(a.Run())
}
