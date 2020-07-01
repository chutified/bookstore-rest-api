package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
	"tommychu/workdir/026_api-example-v2/app/models"
	"tommychu/workdir/026_api-example-v2/app/services/dbservices"
	"tommychu/workdir/026_api-example-v2/config"

	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/go-playground/assert.v1"
)

type nilWriter struct{}

func (w nilWriter) Write([]byte) (int, error) {
	return 0, nil
}

func TestGetAllBooks(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	cfg := config.GetConfig()
	cfg.Log.Output = nilWriter{}

	// tests table
	tests := []struct {
		name         string
		action       func(*gorm.DB)
		expectedCode int
	}{
		{
			name:         "ok",
			action:       func(db *gorm.DB) {},
			expectedCode: 200,
		},
		{
			name:         "database down",
			action:       func(db *gorm.DB) { db.Close() },
			expectedCode: 500,
		},
	}

	// run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db, _ := dbservices.GetDB(cfg)
			defer db.Close()
			router := GetRouter(cfg, db)

			// set
			test.action(db)
			w := httptest.NewRecorder()
			r, _ := http.NewRequest("GET", "/api/v1/books", nil)
			router.ServeHTTP(w, r)

			// check
			assert.Equal(t, test.expectedCode, w.Code)
		})
	}
}

func TestGetBook(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	cfg := config.GetConfig()
	cfg.Log.Output = nilWriter{}

	// create a test model
	var (
		sku         = "######"
		title       = "testModel"
		author      = "testModel"
		description = "testModel"
		price       = 0.0
	)
	testModel := models.Book{
		Model: models.Model{
			ID:        0,
			CreatedAt: time.Unix(0, 0),
			UpdatedAt: time.Unix(0, 0),
		},
		SKU:         &sku,
		Title:       &title,
		Author:      &author,
		Description: description,
		Price:       &price,
	}

	// insert
	db, _ := dbservices.GetDB(cfg)
	if db.NewRecord(testModel) {
		db = db.Create(&testModel)
	}
	defer db.Unscoped().Delete(&testModel)
	testModelBytes, _ := json.Marshal(&testModel)
	db.Close()

	// tests table
	tests := []struct {
		name           string
		id             string
		action         func(*gorm.DB)
		expectedCode   int
		expectedOutput string
	}{
		{
			name:           "ok",
			id:             fmt.Sprintf("%v", testModel.ID),
			action:         func(db *gorm.DB) {},
			expectedCode:   200,
			expectedOutput: string(testModelBytes),
		},
		{
			name:           "invalid id 1",
			id:             "none",
			action:         func(db *gorm.DB) {},
			expectedCode:   400,
			expectedOutput: "",
		},
		{
			name:           "invalid id 2",
			id:             "-1",
			action:         func(db *gorm.DB) {},
			expectedCode:   400,
			expectedOutput: "",
		},
		{
			name:           "database down",
			id:             fmt.Sprintf("%v", testModel.ID),
			action:         func(db *gorm.DB) { db.Close() },
			expectedCode:   500,
			expectedOutput: "",
		},
	}

	// run tests
	for _, test := range tests {
		t.Run(test.name, func(*testing.T) {
			db, _ := dbservices.GetDB(cfg)
			defer db.Close()
			router := GetRouter(cfg, db)

			// set
			test.action(db)
			w := httptest.NewRecorder()
			r, _ := http.NewRequest("GET", fmt.Sprintf("/api/v1/books/%v", test.id), nil)
			router.ServeHTTP(w, r)

			// check
			if test.expectedOutput != "" {
				assert.Equal(t, test.expectedOutput, w.Body.String())
			}
			assert.Equal(t, test.expectedCode, w.Code)
		})
	}
}

func TestNewBook(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	cfg := config.GetConfig()
	cfg.Log.Output = nilWriter{}

	testBody := `{"sku":"test","title":"test","author":"test","price":0}`
	invalidBody1 := `{"sku":"######","title":"test","author":"test","price":0}`
	invalidBody2 := `{"sku":"test","author":"test","price":0}`
	invalidBody3 := `test`

	// tests table
	tests := []struct {
		name         string
		action       func(*gorm.DB)
		body         string
		expectedCode int
	}{
		{
			name:         "ok",
			action:       func(db *gorm.DB) {},
			body:         testBody,
			expectedCode: 200,
		},
		{
			name:         "invalid model validation",
			action:       func(db *gorm.DB) {},
			body:         invalidBody1,
			expectedCode: 400,
		},
		{
			name:         "invalid model missing required",
			action:       func(db *gorm.DB) {},
			body:         invalidBody2,
			expectedCode: 400,
		},
		{
			name:         "invalid model wrong data type",
			action:       func(db *gorm.DB) {},
			body:         invalidBody3,
			expectedCode: 400,
		},
		{
			name:         "database down",
			action:       func(db *gorm.DB) { db.Close() },
			body:         testBody,
			expectedCode: 500,
		},
	}

	// run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db, _ := dbservices.GetDB(cfg)
			defer db.Close()
			router := GetRouter(cfg, db)

			// set
			w := httptest.NewRecorder()
			r, _ := http.NewRequest("POST", "/api/v1/books", strings.NewReader(test.body))
			test.action(db)
			router.ServeHTTP(w, r)

			// clean up
			var m models.Book
			err := json.Unmarshal(w.Body.Bytes(), &m)
			if err == nil {
				if !(cmp.Equal(m, models.Book{})) {
					db.Unscoped().Delete(&m)
				}
			}

			// check
			assert.Equal(t, test.expectedCode, w.Code)
		})
	}
}

func TestUpdateBook(t *testing.T) {
}

func TestRemoveBook(t *testing.T) {
}

func TestRecoverBook(t *testing.T) {
}
