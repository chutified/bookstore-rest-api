package handlers

import (
	"fmt"

	"github.com/chutified/bookstore-api/app/middlewares"
	"github.com/chutified/bookstore-api/config"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// GetRouter returns the set up gin router.
func GetRouter(cfg *config.Config, db *gorm.DB) *gin.Engine {
	r := gin.New()

	// crash free middleware
	r.Use(gin.Recovery())

	//logging middleware
	r.Use(func() gin.HandlerFunc {
		return gin.LoggerWithConfig(gin.LoggerConfig{
			Output: cfg.Log.Output,
		})
	}())

	// routing
	v1 := r.Group("/api/v1/")
	v1.Use(middlewares.DBConn(db))
	{
		v1.GET("/books", GetAllBooks)
		v1.GET("/books/:id", GetBook)
		v1.POST("/books", NewBook)
		v1.PUT("/books/:id", UpdateBook)
		v1.DELETE("/books/:id", RemoveBook)
		v1.POST("/books/:id/recover", RecoverBook)
	}

	// ping
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// documentation
	url := ginSwagger.URL(fmt.Sprintf("http://localhost%s/swagger/doc.json", cfg.Srv.Addr))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}
