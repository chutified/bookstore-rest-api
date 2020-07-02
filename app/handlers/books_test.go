package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/chutified/bookstore-api-example/app/dbservices"
	"github.com/chutified/bookstore-api-example/app/models"
	"github.com/chutified/bookstore-api-example/config"
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

	// test model
	var (
		sku         = "######"
		title       = "testModel"
		author      = "testModel"
		description = "testModel"
		price       = 0.0
	)
	testModel := models.Book{
		Model: models.Model{
			CreatedAt: time.Unix(0, 0),
			UpdatedAt: time.Unix(0, 0),
		},
		SKU:         &sku,
		Title:       &title,
		Author:      &author,
		Description: description,
		Price:       &price,
	}
	testModelBytes, _ := json.Marshal(&testModel)

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
			id:             "0",
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
			id:             "0",
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

			// insert
			db = db.Create(&testModel)
			db.DB().Exec(fmt.Sprintf("UPDATE books SET id = 0 WHERE id = %v;", testModel.Model.ID))

			router := GetRouter(cfg, db)

			// set
			test.action(db)
			w := httptest.NewRecorder()
			r, _ := http.NewRequest("GET", fmt.Sprintf("/api/v1/books/%v", test.id), nil)
			router.ServeHTTP(w, r)

			// clean up
			db, _ = dbservices.GetDB(cfg)
			defer db.Close()
			db.DB().Exec("DELETE FROM books WHERE id = 0;")

			// check
			if test.expectedOutput != "" {
				var m1 models.Book
				json.Unmarshal([]byte(test.expectedOutput), &m1)

				var m2 models.Book
				json.Unmarshal(w.Body.Bytes(), &m2)

				m2.UpdatedAt = m1.UpdatedAt

				assert.Equal(t, m1, m2)
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
	invalidBody1 := `{"sku":"test",title":"test","author":"test","price":0}`
	invalidBody2 := `{"sku":"test","author":"test","price":0}`
	invalidBody3 := `test`

	// tests table
	tests := []struct {
		name         string
		actionBefore func(*gorm.DB)
		actionAfter  func(*gorm.DB)
		body         string
		expectedCode int
	}{
		{
			name:         "ok",
			actionBefore: func(db *gorm.DB) {},
			actionAfter:  func(db *gorm.DB) {},
			body:         testBody,
			expectedCode: 200,
		},
		{
			name: "invalid model validation",
			actionBefore: func(db *gorm.DB) {
				db.DB().Exec(`INSERT INTO books (sku,title,author,price) VALUES ('test','test','test',0);`)
			},
			actionAfter: func(db *gorm.DB) {
				db.DB().Exec(`DELETE FROM books WHERE sku = 'test';`)
			},
			body:         invalidBody1,
			expectedCode: 400,
		},
		{
			name:         "invalid model missing required",
			actionBefore: func(db *gorm.DB) {},
			actionAfter:  func(db *gorm.DB) {},
			body:         invalidBody2,
			expectedCode: 400,
		},
		{
			name:         "invalid model wrong data type",
			actionBefore: func(db *gorm.DB) {},
			actionAfter:  func(db *gorm.DB) {},
			body:         invalidBody3,
			expectedCode: 400,
		},
		{
			name:         "database down",
			actionBefore: func(db *gorm.DB) { db.Close() },
			actionAfter:  func(db *gorm.DB) {},
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
			test.actionBefore(db)
			router.ServeHTTP(w, r)
			test.actionAfter(db)

			// success
			if w.Code == 200 {
				var b models.Book
				db.Where("sku = 'test'").First(&b)
				assert.Equal(t, true, *b.Title == "test")
			}

			// clean up
			var m models.Book
			json.Unmarshal(w.Body.Bytes(), &m)
			if !(cmp.Equal(m, models.Book{})) {
				db.Unscoped().Delete(&m)
			}

			// check
			assert.Equal(t, test.expectedCode, w.Code)
		})
	}
}

func TestUpdateBook(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	cfg := config.GetConfig()
	cfg.Log.Output = nilWriter{}

	// test model
	var (
		sku    = "test"
		title  = "test"
		author = "test"
		price  = 0.0
	)
	testModel := models.Book{
		SKU:    &sku,
		Title:  &title,
		Author: &author,
		Price:  &price,
	}
	var (
		sku2    = "test2"
		title2  = "test2"
		author2 = "test2"
		price2  = 0.0
	)
	testModel2 := models.Book{
		SKU:    &sku2,
		Title:  &title2,
		Author: &author2,
		Price:  &price2,
	}

	// updates
	validModel := `{"sku":"test3"}`
	invalidModel := `{"sku":"test"}`

	// tests table
	tests := []struct {
		name         string
		id           string
		body         string
		action       func(*gorm.DB)
		expectedCode int
	}{
		{
			name:         "ok",
			action:       func(db *gorm.DB) {},
			id:           "1",
			body:         validModel,
			expectedCode: 200,
		},
		{
			name:         "invalid id",
			action:       func(db *gorm.DB) {},
			id:           "test",
			body:         validModel,
			expectedCode: 400,
		},
		{
			name:         "invalid model duplicate",
			action:       func(db *gorm.DB) {},
			id:           "1",
			body:         invalidModel,
			expectedCode: 400,
		},
		{
			name:         "invalid model type",
			action:       func(db *gorm.DB) {},
			id:           "1",
			body:         "test",
			expectedCode: 400,
		},
		{
			name:         "database down",
			action:       func(db *gorm.DB) { db.Close() },
			id:           "1",
			body:         validModel,
			expectedCode: 500,
		},
	}

	// run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db, _ := dbservices.GetDB(cfg)
			defer db.Close()

			// insert
			db = db.Create(&testModel)
			db.DB().Exec(fmt.Sprintf("UPDATE books SET id = 0 WHERE id = %v;", testModel.Model.ID))
			db = db.Create(&testModel2)
			db.DB().Exec(fmt.Sprintf("UPDATE books SET id = 1 WHERE id = %v;", testModel2.Model.ID))

			router := GetRouter(cfg, db)

			// set
			test.action(db)
			w := httptest.NewRecorder()
			r, _ := http.NewRequest("PUT", fmt.Sprintf("/api/v1/books/%s", test.id), strings.NewReader(test.body))
			router.ServeHTTP(w, r)

			// success
			if test.expectedCode == 200 {
				var b models.Book
				db.Where("id = 1").First(&b)
				assert.Equal(t, true, *b.SKU == "test3")
			}

			// clean up
			db, _ = dbservices.GetDB(cfg)
			defer db.Close()
			db.DB().Exec("DELETE FROM books WHERE id = 0;")
			db.DB().Exec("DELETE FROM books WHERE id = 1;")

			// check
			assert.Equal(t, test.expectedCode, w.Code)
		})
	}
}

func TestRemoveBook(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	cfg := config.GetConfig()
	cfg.Log.Output = nilWriter{}

	// test model
	var (
		sku         = "test"
		title       = "test"
		author      = "test"
		description = "test"
		price       = 0.0
	)
	testModel := models.Book{
		Model: models.Model{
			CreatedAt: time.Unix(0, 0),
			UpdatedAt: time.Unix(0, 0),
		},
		SKU:         &sku,
		Title:       &title,
		Author:      &author,
		Description: description,
		Price:       &price,
	}

	// tests table
	tests := []struct {
		name         string
		id           string
		action       func(*gorm.DB)
		expectedCode int
	}{
		{
			name:         "ok",
			id:           "0",
			action:       func(db *gorm.DB) {},
			expectedCode: 200,
		},
		{
			name:         "invalid id non-parseable",
			id:           "non",
			action:       func(db *gorm.DB) {},
			expectedCode: 400,
		},
		{
			name:         "database down",
			id:           "0",
			action:       func(db *gorm.DB) { db.Close() },
			expectedCode: 500,
		},
	}

	// run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db, _ := dbservices.GetDB(cfg)
			defer db.Close()

			// insert
			db = db.Create(&testModel)
			db.DB().Exec(fmt.Sprintf("UPDATE books SET id = 0 WHERE id = %v;", testModel.Model.ID))

			router := GetRouter(cfg, db)

			// set
			test.action(db)
			w := httptest.NewRecorder()
			r, _ := http.NewRequest("DELETE", fmt.Sprintf("/api/v1/books/%v", test.id), nil)
			router.ServeHTTP(w, r)

			// success
			if test.expectedCode == 200 {
				var b models.Book
				db.Unscoped().Where("id = 0").First(&b)
				assert.Equal(t, true, b.DeletedAt != nil)
			}

			// clean up
			db, _ = dbservices.GetDB(cfg)
			defer db.Close()
			db.DB().Exec("DELETE FROM books WHERE id = 0;")

			// check
			assert.Equal(t, test.expectedCode, w.Code)
		})
	}
}

func TestRecoverBook(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	cfg := config.GetConfig()
	cfg.Log.Output = nilWriter{}

	// test model
	var (
		sku         = "test"
		title       = "test"
		author      = "test"
		description = "test"
		price       = 0.0
	)
	testModel := models.Book{
		SKU:         &sku,
		Title:       &title,
		Author:      &author,
		Description: description,
		Price:       &price,
	}

	// tests table
	tests := []struct {
		name         string
		id           string
		action       func(*gorm.DB)
		expectedCode int
	}{
		{
			name:         "ok",
			id:           "0",
			action:       func(db *gorm.DB) {},
			expectedCode: 200,
		},
		{
			name:         "invalid id non-parseable",
			id:           "non",
			action:       func(db *gorm.DB) {},
			expectedCode: 400,
		},
		{
			name:         "invalid id non-existing",
			id:           "-1",
			action:       func(db *gorm.DB) {},
			expectedCode: 400,
		},
		{
			name:         "database down",
			id:           "0",
			action:       func(db *gorm.DB) { db.Close() },
			expectedCode: 500,
		},
	}

	// run tests
	for _, test := range tests {
		// create, delete, recover, check, cleanup
		t.Run(test.name, func(t *testing.T) {
			db, _ := dbservices.GetDB(cfg)
			defer db.Close()

			// insert
			db = db.Create(&testModel)
			db.DB().Exec(fmt.Sprintf("UPDATE books SET id = 0 WHERE id = %v;", testModel.Model.ID))
			db.Delete(models.Book{}, "id = 0")

			router := GetRouter(cfg, db)

			// set
			test.action(db)
			w := httptest.NewRecorder()
			r, _ := http.NewRequest("POST", fmt.Sprintf("/api/v1/books/%v/recover", test.id), nil)
			router.ServeHTTP(w, r)

			// success
			if test.expectedCode == 200 {
				var b models.Book
				db.Where("id = 0").First(&b)
				assert.Equal(t, true, b.DeletedAt == nil)
			}

			// clean up
			db, _ = dbservices.GetDB(cfg)
			defer db.Close()
			db.DB().Exec("DELETE FROM books WHERE id = 0;")

			// check
			assert.Equal(t, test.expectedCode, w.Code)
		})
	}
}
