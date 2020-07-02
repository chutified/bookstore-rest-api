package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chutified/bookstore-api-example/app/dbservices"
	"github.com/chutified/bookstore-api-example/config"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/assert.v1"
)

func TestGetRouter(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	cfg := config.GetConfig()
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
