package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID               uuid.UUID      `db:"id"`
	Name             sql.NullString `db:"name"`
	OriginURL        sql.NullString `db:"origin_url"`
	PageCount        sql.NullInt32  `db:"page_count"`
	AttributesParsed bool           `db:"attributes_parsed"`
	CreateAt         time.Time      `db:"create_at"`
}



type Attribute struct {
	Code        string         `db:"code"`
	Name        string         `db:"name"`
	PluralName  string         `db:"plural_name"`
	Description sql.NullString `db:"description"`
}

type BookAttribute struct {
	BookID uuid.UUID `db:"book_id"`
	Attr   string    `db:"attr"`
	Value  string    `db:"value"`
}

type BookLabel struct {
	BookID     uuid.UUID     `db:"book_id"`
	PageNumber sql.NullInt32 `db:"page_number"`
	Name       string        `db:"name"`
	Value      string        `db:"value"`
	CreateAt   time.Time     `db:"create_at"`
}

type Agent struct {
	ID        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	Addr      string    `db:"addr"`
	Token     string    `db:"token"`
	CanParse  bool      `db:"can_parse"`
	CanExport bool      `db:"can_export"`
	Priority  int       `db:"priority"`
	CreateAt  time.Time `db:"create_at"`
}
