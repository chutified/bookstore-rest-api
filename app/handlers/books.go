package handlers

import (
	"strconv"
	"tommychu/workdir/027_api-example-v2/app/models"
	services "tommychu/workdir/027_api-example-v2/app/services/db"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GetAllBooks lists books.
// @Summary List books
// @Description Get all books from the database that are not marked as deleted.
// @Produce json
// @Success 200 {array} models.Book "listed - ok"
// @Failure 500 {object} models.AppErrors "internal error"
// @Router /books [get]
func GetAllBooks(c *gin.Context) {
	db := c.Value("db").(*gorm.DB)

	// get books
	books, errs := services.ReadAllBooks(db)
	if len(errs) != 0 {
		c.JSON(500, handleErrs(errs...))
		return
	}

	// success
	c.JSON(200, books)
}

// GetBook get a book.
// @Summary Get a book
// @Description Get a book by its ID that is not marked as deleted.
// @Produce json
// @Param id path int true "book id"
// @Success 200 {object} models.Book "got - ok"
// @Failure 400 {object} models.AppErrors "bad request"
// @Failure 500 {object} models.AppErrors "internal error"
// @Router /books/{id} [get]
func GetBook(c *gin.Context) {
	db := c.Value("db").(*gorm.DB)

	// get id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, handleErrs(err))
		return
	}

	// get the book
	book, errs, err := services.ReadBook(db, id)
	if err != nil {
		c.JSON(400, handleErrs(err))
		return
	}
	if len(errs) != 0 {
		c.JSON(500, handleErrs(errs...))
		return
	}

	// success
	c.JSON(200, book)
}

// NewBook creates a book.
// @Summary Create book
// @Description Create a new book with unique SKU.
// @Accept json
// @Produce json
// @Param book body models.Book true "book struct"
// @Success 200 {object} models.Book "book created - ok"
// @Failure 400 {object} models.AppErrors "bad request"
// @Router /books [post]
func NewBook(c *gin.Context) {
	db := c.Value("db").(*gorm.DB)

	// get the book
	var book models.Book
	err := c.BindJSON(&book)
	if err != nil {
		c.JSON(400, handleErrs(err))
		return
	}

	// create book
	b, errs, err := services.CreateBook(db, book)
	if err != nil {
		c.JSON(400, handleErrs(err))
		return
	}
	if len(errs) != 0 {
		c.JSON(400, handleErrs(errs...))
		return
	}

	// success
	c.JSON(200, b)
}

// UpdateBook updates a book.
// @Summary Update book
// @Description Find a book by its ID and update it with changed fields.
// @Accept json
// @Produce json
// @Param id path int true "book id"
// @Param book body models.Book true "book struct"
// @Success 200 {object} models.Book "updated - ok"
// @Failure 400 {object} models.AppErrors "bad request"
// @Router /books/{id} [put]
func UpdateBook(c *gin.Context) {
	db := c.Value("db").(*gorm.DB)

	// get id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, handleErrs(err))
		return
	}

	// get book
	var book models.Book
	err = c.BindJSON(&book)
	if err != nil {
		c.JSON(400, handleErrs(err))
		return
	}

	// update
	book, errs := services.UpdateBook(db, id, book)
	if len(errs) != 0 {
		c.JSON(400, handleErrs(errs...))
		return
	}

	// success
	c.JSON(200, book)
}

// RemoveBook deletes a book.
// @Summary Delete book.
// @Description Find a book by its ID and deletes it.
// @Produce json
// @Param id path int true "book id"
// @Success 200 {string} string "deleted book id - ok"
// @Failure 400 {object} models.AppErrors "bad request"
// @Router /books/{id} [delete]
func RemoveBook(c *gin.Context) {
	db := c.Value("db").(*gorm.DB)

	// get id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, handleErrs(err))
		return
	}

	// deletiton
	errs := services.DeleteBook(db, id)
	if len(errs) != 0 {
		c.JSON(400, handleErrs(errs...))
		return
	}

	// success
	c.JSON(200, gin.H{
		"delete_book_id": id,
	})
}

// RecoverBook recovers a deleted book.
// @Summary Recover deleted book
// @Description Find a book by its ID and remove a deleted tag from it.
// @Produce json
// @Param id path int true "book id"
// @Success 200 {object} models.Book "recovered - ok"
// @Failure 400 {object} models.AppErrors "bad request"
// @Router /books/{id}/recover [post]
func RecoverBook(c *gin.Context) {
	db := c.Value("db").(*gorm.DB)

	// get id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, handleErrs(err))
		return
	}

	// recover book
	book, errs := services.RecoverBook(db, id)
	if len(errs) != 0 {
		c.JSON(400, handleErrs(errs...))
		return
	}

	// success
	c.JSON(200, book)
}
