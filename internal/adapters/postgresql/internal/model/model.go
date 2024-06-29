package model

import (
	"database/sql"
)

type Attribute struct {
	Code        string         `db:"code"`
	Name        string         `db:"name"`
	PluralName  string         `db:"plural_name"`
	Order       int            `db:"order"`
	Description sql.NullString `db:"description"`
}

type BookAttribute struct {
	BookID string `db:"book_id"`
	Attr   string `db:"attr"`
	Value  string `db:"value"`
}
