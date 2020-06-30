package main

import (
	"tommychu/workdir/027_api-example-v2/app"
	"tommychu/workdir/027_api-example-v2/config"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "tommychu/workdir/027_api-example-v2/docs"
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

	// documentation
	url := ginSwagger.URL("http://localhost:8081/swagger/doc.json") // The url pointing to API definition
	a.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	a.Run()
}
