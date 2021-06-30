package dbservices

import (
	"fmt"

	"github.com/chutommy/bookstore-api/app/models"
	"github.com/chutommy/bookstore-api/config"
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

	if !cfg.DebugMode {
		db.LogMode(false)
	}
	return dbMigrate(db), nil
}

// dbMigrate is a database migration.
func dbMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&models.Book{})
	return db
}
