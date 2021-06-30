package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chutommy/bookstore-api/app/dbservices"
	"github.com/chutommy/bookstore-api/config"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/assert.v1"
)

func TestGetRouter(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	cfg, err := config.GetConfig()
	if err != nil {
		t.Fatalf("could not get config: %v", err)
	}
	cfg.Log.Output = nilWriter{}
	db, _ := dbservices.GetDB(cfg)
	db.LogMode(false)
	router := GetRouter(cfg, db)

	// set
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, r)

	// check
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
