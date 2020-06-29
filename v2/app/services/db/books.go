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

// ReadAllBooks finds all books in the database and returns them.
func ReadAllBooks(db *gorm.DB) ([]models.Book, []error) {
	var books []models.Book
	return books, db.Find(&books).GetErrors()
}

// ReadBook finds a books with a specific ID.
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

// CreateBook adds a new book into the database.
func CreateBook(db *gorm.DB, book models.Book) (models.Book, []error, error) {
	if !db.NewRecord(book) {
		return models.Book{}, nil, fmt.Errorf("this book already exists: %v", book)
	}
	return book, db.Create(&book).GetErrors(), nil
}

// UpdateBook updates chenged fields.
func UpdateBook(db *gorm.DB, id int, book models.Book) (models.Book, []error) {
	var changing models.Book
	errs := db.First(&changing, id).GetErrors()
	if len(errs) != 0 {
		return models.Book{}, errs
	}
	return book, db.Model(changing).Updates(book).GetErrors()
}

// DeleteBook marks the book as deleted.
func DeleteBook(db *gorm.DB, id int) []error {
	var book models.Book
	return db.Delete(&book, id).GetErrors()
}
