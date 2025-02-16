package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type Attribute struct {
	Code        string         `db:"code"`
	Name        string         `db:"name"`
	PluralName  string         `db:"plural_name"`
	Order       int            `db:"order"`
	Description sql.NullString `db:"description"`
}

type BookAttribute struct {
	BookID uuid.UUID `db:"book_id"`
	Attr   string    `db:"attr"`
	Value  string    `db:"value"`
}

type BookOriginalAttribute struct {
	BookID uuid.UUID `db:"book_id"`
	Attr   string    `db:"attr"`
	Value  []string  `db:"values"`
}
