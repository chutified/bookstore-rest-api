package middlewares

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func TestDBConn(t *testing.T) {

	// set
	db := &gorm.DB{Value: "test"}
	c := &gin.Context{}
	DBConn(db)(c)

	// check
	if got := c.Value("db").(*gorm.DB).Value.(string); got != "test" {
		t.Errorf("expected: %s, got: %s", "test", got)
	}
}
