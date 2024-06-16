package model

import (
	"database/sql"
	"time"
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

type BookLabel struct {
	BookID     string        `db:"book_id"`
	PageNumber sql.NullInt32 `db:"page_number"`
	Name       string        `db:"name"`
	Value      string        `db:"value"`
	CreateAt   time.Time     `db:"create_at"`
}
