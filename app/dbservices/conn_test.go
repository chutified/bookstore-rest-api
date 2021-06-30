package dbservices

import (
	"testing"

	"github.com/chutommy/bookstore-api/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/go-playground/assert.v1"
)

func TestGetDB(t *testing.T) {

	// tests table
	tests := []struct {
		name         string
		actionBefore func(*config.Config)
		actionAfter  func(*gorm.DB)
		noErr1       bool
		noErr2       bool
	}{
		{
			name:         "ok",
			actionBefore: func(cfg *config.Config) {},
			actionAfter:  func(db *gorm.DB) {},
			noErr1:       true,
			noErr2:       true,
		},
		{
			name:         "database down",
			actionBefore: func(cfg *config.Config) {},
			actionAfter:  func(db *gorm.DB) { db.Close() },
			noErr1:       true,
			noErr2:       false,
		},
		{
			name:         "invalid config",
			actionBefore: func(cfg *config.Config) { cfg.DB.DBName = "###_invalid_database_###" },
			actionAfter:  func(db *gorm.DB) {},
			noErr1:       false,
			noErr2:       false,
		},
	}

	// run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cfg, err := config.GetConfig()
			if err != nil {
				t.Fatalf("could not get config: %v", err)
			}

			// get
			test.actionBefore(cfg)
			db, err1 := GetDB(cfg)
			test.actionAfter(db)

			// check
			assert.Equal(t, test.noErr1, (err1 == nil))
			if err1 == nil {
				defer db.Close()
				err2 := db.DB().Ping()
				assert.Equal(t, test.noErr2, (err2 == nil))
			}
		})
	}
}
