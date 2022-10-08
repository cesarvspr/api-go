package models

import (
	"time"

	"github.com/google/uuid"
)

// Book struct to describe book object.
type Book struct {
	ID         uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
	Title      string    `db:"title" json:"title" validate:"required,lte=255"`
	Author     string    `db:"author" json:"author" validate:"required,lte=255"`
	BookStatus bool      `db:"book_status" json:"book_status" validate:"required"`
	BookAttrs  BookAttrs `db:"book_attrs" json:"book_attrs" validate:"required,dive"`
}

// BookAttrs struct to describe book attributes.
type BookAttrs struct {
	Picture     string `json:"picture"`
	Description string `json:"description"`
	Rating      int    `json:"rating" validate:"min=1,max=10"`
}

// Value make the BookAttrs struct implement the driver.Valuer interface.
// This method simply returns the JSON-encoded representation of the struct.
func (b Book) Value() map[string]interface{} {
	return map[string]interface{}{
		"id":          b.ID,
		"created_at":  b.CreatedAt,
		"updated_at":  b.UpdatedAt,
		"title":       b.Title,
		"author":      b.Author,
		"book_status": b.BookStatus,
		"book_attrs":  b.BookAttrs,
	}
}
