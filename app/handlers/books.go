package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"tommychu/workdir/026_api-example/app/models"
	"tommychu/workdir/026_api-example/app/services/dbservice"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// unWrap unwraps all errors into a string
func unWrap(errs []error) string {
	var result string
	for _, err := range errs {
		result = fmt.Sprintf("%s\n%s", result, err.Error())
	}
	return result
}

// NewBook handles a new book creation.
func NewBook(db *gorm.DB, l *log.Logger, w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json")

	// get the Book
	var b models.Book
	err := b.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Invalid request body: could not unmarshal JSON: "+err.Error(), 400)
		return
	}

	// insert
	ok, errs := dbservice.Create(db, b)
	if !ok {
		http.Error(w, "record with the same ID already exists", 400)
		return
	}
	if len(errs) != 0 {
		http.Error(w, unWrap(errs), 400)
		return
	}

	// success
	l.Printf("[NEW] %s (%s)\n", b.Author, b.Name)
	w.WriteHeader(200)
	if err := b.ToJSON(w); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// GetBook handles the one book with specific ID.
func GetBook(db *gorm.DB, l *log.Logger, w http.ResponseWriter, r *http.Request) {

	// get id
	strID := mux.Vars(r)["id"]
	id, err := strconv.Atoi(strID)
	if err != nil {
		http.Error(w, "invalid id", 400)
		return
	}

	// retrive the book
	b, errs := dbservice.ReadOneBook(db, id)
	if len(errs) != 0 {
		http.Error(w, unWrap(errs), 400)
		return
	}

	// success
	l.Printf("[SEARCHED] %s (%s)\n", b.Author, b.Name)
	w.WriteHeader(200)
	if err := b.ToJSON(w); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// GetAllBooks handles request for all books in the database.
func GetAllBooks(db *gorm.DB, l *log.Logger, w http.ResponseWriter, r *http.Request) {

	// get books
	bs, errs := dbservice.ReadAllBooks(db)
	if len(errs) != 0 {
		http.Error(w, unWrap(errs), 500)
		return
	}

	// success
	l.Printf("[SEARCHED] ALL\n")
	w.WriteHeader(200)
	if err := bs.ToJSON(w); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// UpdateBook updates the book.
// Only values listed in JSON will be changed.
func UpdateBook(db *gorm.DB, l *log.Logger, w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json")

	// get id
	strID := mux.Vars(r)["id"]
	id, err := strconv.Atoi(strID)
	if err != nil {
		http.Error(w, "invalid id", 400)
		return
	}

	// get changes
	var b models.Book
	err = b.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Invalid request body: could not unmarshal JSON", 400)
		return
	}

	b, errs := dbservice.UpdateBook(db, id, b)
	if len(errs) != 0 {
		http.Error(w, unWrap(errs), 500)
		return
	}
	// success
	l.Printf("[UPDATED] %s (%s)\n", b.Author, b.Name)
	w.WriteHeader(200)
}

// RemoveBook removes the book.
func RemoveBook(db *gorm.DB, l *log.Logger, w http.ResponseWriter, r *http.Request) {

	// get id
	strID := mux.Vars(r)["id"]
	id, err := strconv.Atoi(strID)
	if err != nil {
		http.Error(w, "invalid id", 400)
		return
	}

	// delete
	b, exist, errs := dbservice.DeleteBook(db, id)
	if !exist {
		http.Error(w, "book with this id does not exist", 400)
		return
	}
	if len(errs) != 0 {
		http.Error(w, unWrap(errs), 500)
		return
	}

	// success
	l.Printf("[DELETED] %s (%s)\n", b.Author, b.Name)
	w.WriteHeader(200)
}
