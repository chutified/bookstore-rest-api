package dbservice

import (
	"tommychu/workdir/026_api-example/app/models"

	"github.com/jinzhu/gorm"
)

// DBMigrate is a database migration
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&models.Book{})
	return db
}

// Create creates a new book record in the database.
// If a record with same ID exists return false.
func Create(db *gorm.DB, b models.Book) (bool, []error) {
	if !db.NewRecord(b) {
		return false, nil
	}
	errs := db.Create(&b).GetErrors()
	return true, errs
}

// ReadOneBook gets the book by the ID.
func ReadOneBook(db *gorm.DB, id int) (*models.Book, []error) {
	var b models.Book
	errs := db.First(&b, id).GetErrors()
	return &b, errs
}

// ReadAllBooks gets all books from the database.
func ReadAllBooks(db *gorm.DB) (*models.Books, []error) {
	var bs models.Books
	errs := db.Find(&bs).GetErrors()
	return &bs, errs
}

// Update changes values of the Book.
func UpdateBook(db *gorm.DB, id int, changes models.Book) (models.Book, []error) {
	var b models.Book
	errs := db.First(&b, id).GetErrors()
	if len(errs) != 0 {
		return models.Book{}, errs
	}
	return b, db.Model(&b).Updates(changes).GetErrors()
}

// DeleteBook finds a book and removes it.
func DeleteBook(db *gorm.DB, id int) (models.Book, bool, []error) {
	var b models.Book
	errs := db.First(&b, id).GetErrors()
	if len(errs) != 0 {
		return models.Book{}, false, errs
	}
	errs = db.Delete(&b).GetErrors()
	if len(errs) != 0 {
		return models.Book{}, true, errs
	}
	return b, true, nil
}
