package models

import (
	"time"
)

// Model is a copy of the gorm.Model struct (for the purpose of documentation).
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// Book is a bookstore product struct.
type Book struct {
	Model
	SKU         *string  `json:"sku,omitempty" gorm:"not null;unique"`
	Title       *string  `json:"title,omitempty" gorm:"not null"`
	Author      *string  `json:"author,omitempty" gorm:"not null"`
	Description string   `json:"description,omitempty" gorm:"type:varchar(600)"`
	Price       *float64 `json:"price,omitempty" gorm:"not null"`
}
