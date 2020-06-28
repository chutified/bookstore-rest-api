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
// @Param sku body string true "Stock Keeping Unit" validate(required)
// @Param name body string true "The name of the book" validate(required)
// @Param author body string true "The author's name" validate(required)
// @Param description body string false "The book's description"
// @Param price body string true "The book's value" validate(required) default(0) minimum(0)
// @Produce json
// @Success 200 {string} Book "Created"
// @Failure 400 {string} string "input"
// @Failure 500 {string} string "JSON unmarshal"
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
	if err := b.ToJSON(w); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(200)
}

// GetBook godoc
// @Summary return a book
// @Description find a book by id and serve it
// @Tags books
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {string} Book "Read"
// @Failure 400 {string} string "input"
// @Failure 500 {string} string "JSON unmarshal"
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
	if err := b.ToJSON(w); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(200)
}

// GetAllBooks godoc
// @Summary return all books
// @Description find all books in the database and serve it
// @Tags books
// @Produce json
// @Success 200 {string} Books "Read"
// @Failure 400 {string} string "input"
// @Failure 500 {string} string "JSON unmarshal"
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
	if err := bs.ToJSON(w); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(200)
}

// UpdateBook godoc
// @Summary update a book
// @Description find a book and update it with new values
// @Tags books
// @Accept json
// @Param sku body string false "Stock Keeping Unit"
// @Param name body string false "The name of the book"
// @Param author body string false "The author's name"
// @Param description body string false "The book's description"
// @Param price body string false "The book's value" minimum(0)
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {string} Book "Updated"
// @Failure 400 {string} string "input"
// @Failure 500 {string} string "JSON unmarshal"
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
	if err := b.ToJSON(w); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(200)
}

// RemoveBook godoc
// @Summary remove a book
// @Description find a book and delete it
// @Tags books
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {string} Book "Deleted"
// @Failure 400 {string} string "input"
// @Failure 500 {string} string "JSON unmarshal"
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
	if err := b.ToJSON(w); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(200)
}
