package main

import (
	"fmt"
	"log"
	"tommychu/workdir/026_api-example-v2/app"
	"tommychu/workdir/026_api-example-v2/config"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
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
	err := a.Initialize(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		log.Fatal(a.Close())
	}()

	// documentation
	url := ginSwagger.URL(fmt.Sprintf("http://localhost%s/swagger/doc.json", cfg.Srv.Addr)) // The url pointing to API definition
	a.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	log.Panic(a.Run())
}
