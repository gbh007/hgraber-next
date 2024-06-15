package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"

	"hgnext/internal/entities"
)

type Book struct {
	ID               string         `db:"id"`
	Name             sql.NullString `db:"name"`
	OriginURL        sql.NullString `db:"origin_url"`
	PageCount        sql.NullInt32  `db:"page_count"`
	AttributesParsed bool           `db:"attributes_parsed"`
	CreateAt         time.Time      `db:"create_at"`
}

func (b Book) ToEntity() (entities.Book, error) {
	id, err := uuid.Parse(b.ID)
	if err != nil {
		return entities.Book{}, err
	}

	return entities.Book{
		ID:               id,
		Name:             b.Name.String,
		OriginURL:        b.OriginURL.String,
		PageCount:        int(b.PageCount.Int32),
		AttributesParsed: b.AttributesParsed,
		CreateAt:         b.CreateAt,
	}, nil
}
