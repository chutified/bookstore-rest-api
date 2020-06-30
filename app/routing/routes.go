package routing

import (
	"tommychu/workdir/027_api-example-v2/app/handlers"
	"tommychu/workdir/027_api-example-v2/app/middlewares"
	"tommychu/workdir/027_api-example-v2/config"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GetRouter returns the set up gin router.
func GetRouter(cfg *config.Config, db *gorm.DB) *gin.Engine {
	r := gin.New()

	// crash free middleware
	r.Use(gin.Recovery())

	// logging middleware
	r.Use(func() gin.HandlerFunc {
		return gin.LoggerWithConfig(gin.LoggerConfig{
			Output: cfg.Log.Output,
		})
	}())

	// routing
	v1 := r.Group("/api/v1/")
	v1.Use(middlewares.DBConn(db))
	{
		v1.GET("/books", handlers.GetAllBooks)
		v1.GET("/books/:id", handlers.GetBook)
		v1.POST("/books", handlers.NewBook)
		v1.PUT("/books/:id", handlers.UpdateBook)
		v1.DELETE("/books/:id", handlers.RemoveBook)
		v1.POST("/books/:id/recover", handlers.RecoverBook)
	}

	return r
}
