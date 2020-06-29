package handlers

import (
	"log"
	"net/http"
	"strconv"
	"tommychu/workdir/026_api-example/app/models"
	"tommychu/workdir/026_api-example/app/services/dbservice"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// NewBook godoc
// @Summary create a new book
// @Description validate a new book and insert it into the database
// @Tags books
// @Accept json
// @Produce json
// @Param book body models.Book true "The book" validate(required)
// @Success 200 {object} models.Book "Book creates"
// @Failure 400 {string} string "Bad JSON"
// @Failure 500 {string} string "JSON unmarshal error"
// @Router /books [post]
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

// GetBook godoc
// @Summary return a book
// @Description find a book by id and serve it
// @Tags books
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} models.Book "Book retrived"
// @Failure 400 {string} string "Bad JSON"
// @Failure 500 {string} string "JSON unmarshal error"
// @Router /books/{id} [get]
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

// GetAllBooks godoc
// @Summary return all books
// @Description find all books in the database and serve it
// @Tags books
// @Produce json
// @Success 200 {object} models.Books "Books retrieved"
// @Failure 400 {string} string "Bad JSON"
// @Failure 500 {string} string "JSON unmarshal error"
// @Router /books [get]
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

// UpdateBook godoc
// @Summary update a book
// @Description find a book and update it with new values
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body models.Book true "The book" validate(required)
// @Success 200 {object} models.Book "Book updated"
// @Failure 400 {string} string "Bad JSON"
// @Failure 500 {string} string "JSON unmarshal error"
// @Router /books/{id} [put]
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
	if err := b.ToJSON(w); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// RemoveBook godoc
// @Summary remove a book
// @Description find a book and delete it
// @Tags books
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} models.Book "Book deleted"
// @Failure 400 {string} string "Bad JSON"
// @Failure 500 {string} string "JSON unmarshal error"
// @Router /books/{id} [delete]
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
	if err := b.ToJSON(w); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
