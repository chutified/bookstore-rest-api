package services

import (
	"fmt"
	"tommychu/workdir/027_api-example-v2/app/models"

	"github.com/jinzhu/gorm"
)

// DBMigrate is a database migration.
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&models.Book{})
	return db
}

// ReadAllBooks gets all books.
func ReadAllBooks(db *gorm.DB) ([]models.Book, []error) {
	var books []models.Book
	return books, db.Find(&books).GetErrors()
}

// ReadBook finds a one specific book.
func ReadBook(db *gorm.DB, id int) (models.Book, []error, error) {
	var book models.Book
	errs := db.First(&book, id).GetErrors()
	if book.Model.ID == 0 {
		return models.Book{}, nil, fmt.Errorf("book with id '%d' does not exists", id)
	}
	if len(errs) != 0 {
		return models.Book{}, errs, nil
	}
	return book, nil, nil
}

// CreateBook creates a new book.
func CreateBook(db *gorm.DB, book models.Book) (models.Book, []error, error) {
	if !db.NewRecord(book) {
		return models.Book{}, nil, fmt.Errorf("this book already exists: %v", book)
	}
	return book, db.Create(&book).GetErrors(), nil
}

// UpdateBook updates the book with th changed fields.
func UpdateBook(db *gorm.DB, id int, book models.Book) (models.Book, []error) {
	var changing models.Book
	errs := db.First(&changing, id).GetErrors()
	if len(errs) != 0 {
		return models.Book{}, errs
	}
	errs = db.Model(&changing).Updates(book).GetErrors()
	return changing, errs
}

// DeleteBook marks the book as deleted.
func DeleteBook(db *gorm.DB, id int) []error {
	var book models.Book
	return db.Delete(&book, id).GetErrors()
}

// RecoverBook removes the deleted timestamp from the book record.
func RecoverBook(db *gorm.DB, id int) (models.Book, []error) {
	var book models.Book
	errs := db.Unscoped().Where("id = ?", id).First(&book).Update("deleted_at", nil).GetErrors()
	if len(errs) != 0 {
		return models.Book{}, errs
	}
	return book, nil
}
