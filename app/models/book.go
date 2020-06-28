package models

import (
	"encoding/json"
	"io"

	"github.com/jinzhu/gorm"
)

// Book is a product.
type Book struct {
	gorm.Model
	SKU         string  `json:"sku" gorm:"not null;unique"`
	Name        string  `json:"name" gorm:"not null;unique"`
	Author      string  `json:"author" gorm:"not null"`
	Description string  `json:"description" gorm:"type:varchar(600)"`
	Price       float64 `json:"price" gorm:"not null;default:0"`
}

// Books is a slice of Books.
type Books []*Book

// FromJSON decodes Book from JSON format.
func (b *Book) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(&b)
}

// ToJSON encodes Book to JSON format
func (b *Book) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(&b)
}

// FromJSON decodes Books from JSON format.
func (bs *Books) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(&bs)
}

// ToJSON encodes Books to JSON format
func (bs *Books) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(&bs)
}
