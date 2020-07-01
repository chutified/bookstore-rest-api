package dbservices

import (
	"fmt"
	"log"
	"tommychu/workdir/026_api-example-v2/app/models"
	"tommychu/workdir/026_api-example-v2/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DBMigrate is a database migration.
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&models.Book{})
	return db
}

// GetDB returns the set up database connection.
func GetDB(cfg *config.Config) *gorm.DB {
	dbURI := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.User,
		cfg.DB.DBName,
		cfg.DB.Password,
	)
	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		log.Fatal(fmt.Errorf("could not Initialize a db connection: %v", err))
	}
	db.LogMode(false) // disable annoying database logs
	return db
}
