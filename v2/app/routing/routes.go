package routing

import (
	"tommychu/workdir/027_api-example-v2/app/handlers"
	"tommychu/workdir/027_api-example-v2/app/middlewares"
	"tommychu/workdir/027_api-example-v2/config"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GetRouter returns set up gin router.
func GetRouter(cfg *config.Config, db *gorm.DB) *gin.Engine {
	r := gin.New()

	// crash free
	r.Use(gin.Recovery())
	// log
	r.Use(func() gin.HandlerFunc {
		return gin.LoggerWithConfig(gin.LoggerConfig{
			Output: cfg.Log.Output,
		})
	}())

	// routing
	v2 := r.Group("/api/v2/")
	v2.Use(middlewares.DBConn(db))
	{
		v2.GET("/books", handlers.GetAllBooks)
		v2.GET("/books/:id", handlers.GetBook)
		v2.POST("/books", handlers.NewBook)
		v2.PUT("/books/:id", handlers.UpdateBook)
		v2.DELETE("/books/:id", handlers.RemoveBook)
	}

	return r
}
