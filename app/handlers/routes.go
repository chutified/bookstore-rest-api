package handlers

import (
	"tommychu/workdir/026_api-example-v2/app/middlewares"
	"tommychu/workdir/026_api-example-v2/config"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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

	return r
}
