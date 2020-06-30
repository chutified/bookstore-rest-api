package handlers

import (
	"strconv"
	"tommychu/workdir/027_api-example-v2/app/models"
	services "tommychu/workdir/027_api-example-v2/app/services/db"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GetAllBooks returns all books in the database.
func GetAllBooks(c *gin.Context) {
	db := c.Value("db").(*gorm.DB)

	// get books
	books, errs := services.ReadAllBooks(db)
	if len(errs) != 0 {
		c.JSON(500, errToJSON(errs...))
		return
	}

	// success
	c.JSON(200, books)
}

// GetBook returns a book with a specific ID.
func GetBook(c *gin.Context) {
	db := c.Value("db").(*gorm.DB)

	// get id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, errToJSON(err))
		return
	}

	// get the book
	book, errs, err := services.ReadBook(db, id)
	if err != nil {
		c.JSON(400, errToJSON(err))
		return
	}
	if len(errs) != 0 {
		c.JSON(500, errToJSON(errs...))
		return
	}

	// success
	c.JSON(200, book)
}

// NewBook creates a new book.
func NewBook(c *gin.Context) {
	db := c.Value("db").(*gorm.DB)

	// get the book
	var book models.Book
	err := c.BindJSON(&book)
	if err != nil {
		c.JSON(400, errToJSON(err))
		return
	}

	// create book
	b, errs, err := services.CreateBook(db, book)
	if err != nil {
		c.JSON(400, errToJSON(err))
		return
	}
	if len(errs) != 0 {
		c.JSON(400, errToJSON(errs...))
		return
	}

	// success
	c.JSON(200, b)
}

// UpdateBook changes book values.
func UpdateBook(c *gin.Context) {
	db := c.Value("db").(*gorm.DB)

	// get id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, errToJSON(err))
		return
	}

	// get book
	var book models.Book
	err = c.BindJSON(&book)
	if err != nil {
		c.JSON(400, errToJSON(err))
		return
	}

	// update
	book, errs := services.UpdateBook(db, id, book)
	if len(errs) != 0 {
		c.JSON(400, errToJSON(errs...))
		return
	}

	// success
	c.JSON(200, book)
}

// RemoveBook deletes a book by ID.
func RemoveBook(c *gin.Context) {
	db := c.Value("db").(*gorm.DB)

	// get id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, errToJSON(err))
		return
	}

	// deletiton
	errs := services.DeleteBook(db, id)
	if len(errs) != 0 {
		c.JSON(400, errToJSON(errs...))
		return
	}

	// success
	c.JSON(200, gin.H{
		"delete_book_id": id,
	})
}

// RecoverBook recovers deleted book by its ID.
func RecoverBook(c *gin.Context) {
	db := c.Value("db").(*gorm.DB)

	// get id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, errToJSON(err))
		return
	}

	// recover book
	book, errs := services.RecoverBook(db, id)
	if len(errs) != 0 {
		c.JSON(400, errToJSON(errs...))
		return
	}

	// success
	c.JSON(200, book)
}
