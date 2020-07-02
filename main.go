package main

import (
	"log"

	"github.com/chutified/bookstore-api-example/app"
	"github.com/chutified/bookstore-api-example/config"
	_ "github.com/chutified/bookstore-api-example/docs"
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
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	if !cfg.DebugMode {
		gin.SetMode(gin.ReleaseMode)
	}

	// set app
	a := app.New()
	err = a.Initialize(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		log.Fatal(a.Close())
	}()

	log.Panic(a.Run())
}
