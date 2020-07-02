package main

import (
	"fmt"
	"log"

	"github.com/chutified/bookstore-api/app"
	"github.com/chutified/bookstore-api/config"
	_ "github.com/chutified/bookstore-api/docs"
	"github.com/gin-gonic/gin"
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

	fmt.Fprintf(a.Log, "Listening and serving HTTP on %s\n", a.Srv.Addr)
	fmt.Fprintf(a.Log, "API Documentation: http://localhost%s/swagger/index.html\n", a.Srv.Addr)
	log.Panic(a.Run())
}
