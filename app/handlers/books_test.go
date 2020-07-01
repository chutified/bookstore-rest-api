package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"tommychu/workdir/026_api-example-v2/app/models"
	"tommychu/workdir/026_api-example-v2/app/services/dbservices"
	"tommychu/workdir/026_api-example-v2/config"

	"gopkg.in/go-playground/assert.v1"
)

func TestGetAllBooks(t *testing.T) {
	cfg := config.GetConfig()
	db := dbservices.GetDB(cfg)
	defer db.Close()
	router := GetRouter(cfg, db)

	// 200
	w1 := httptest.NewRecorder()
	r1, _ := http.NewRequest("GET", "/api/v1/books", nil)
	router.ServeHTTP(w1, r1)

	// 500
	w2 := httptest.NewRecorder()
	r2, _ := http.NewRequest("GET", "/api/v1/books", nil)
	db.Close()
	router.ServeHTTP(w2, r2)

	// compare
	assert.Equal(t, 200, w1.Code)
	assert.Equal(t, 500, w2.Code)
}

func TestGetBook(t *testing.T) {
	cfg := config.GetConfig()
	db := dbservices.GetDB(cfg)
	defer db.Close()
	router := GetRouter(cfg, db)

	var exist models.Book
	db.First(&exist)

	// 200
	w1 := httptest.NewRecorder()
	r1, _ := http.NewRequest("GET", fmt.Sprintf("/api/v1/books/%v", exist.ID), nil)
	router.ServeHTTP(w1, r1)

	// 400
	w2 := httptest.NewRecorder()
	r2, _ := http.NewRequest("GET", fmt.Sprintf("/api/v1/books/none"), nil)
	router.ServeHTTP(w2, r2)

	// 400
	w3 := httptest.NewRecorder()
	r3, _ := http.NewRequest("GET", fmt.Sprintf("/api/v1/books/-1"), nil)
	router.ServeHTTP(w3, r3)

	// 500
	w4 := httptest.NewRecorder()
	r4, _ := http.NewRequest("GET", fmt.Sprintf("/api/v1/books/%v", exist.ID), nil)
	db.Close()
	router.ServeHTTP(w4, r4)

	// compare
	assert.Equal(t, 200, w1.Code)
	assert.Equal(t, 400, w2.Code)
	assert.Equal(t, 400, w3.Code)
	assert.Equal(t, 500, w4.Code)
}

func TestNewBook(t *testing.T) {
}

func TestUpdateBook(t *testing.T) {
}

func TestRemoveBook(t *testing.T) {
}

func TestRecoverBook(t *testing.T) {
}
