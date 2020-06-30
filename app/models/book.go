package models

import "github.com/jinzhu/gorm"

// Book stores the book attributes
type Book struct {
	gorm.Model
	SKU         *string  `json:"sku,omitempty" gorm:"not null;unique"`
	Title       *string  `json:"title,omitempty" gorm:"not null"`
	Author      *string  `json:"author,omitempty" gorm:"not null"`
	Description string   `json:"description,omitempty" gorm:"type:varchar(600)"`
	Price       *float64 `json:"price,omitempty" gorm:"not null"`
}
