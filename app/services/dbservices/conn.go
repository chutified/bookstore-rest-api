package dbservices

import (
	"fmt"

	"github.com/chutified/bookstore-api-example/app/models"
	"github.com/chutified/bookstore-api-example/config"
	"github.com/jinzhu/gorm"
)

// GetDB returns the set up database connection.
func GetDB(cfg *config.Config) (*gorm.DB, error) {
	dbURI := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.User,
		cfg.DB.DBName,
		cfg.DB.Password,
	)
	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		return nil, fmt.Errorf("could not Initialize a db connection: %v", err)
	}

	if !config.DEBUG_MODE {
		db.LogMode(false)
	}
	return dbMigrate(db), nil
}

// dbMigrate is a database migration.
func dbMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&models.Book{})
	return db
}
