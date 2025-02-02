package model

import (
	"database/sql"
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/internal/entities"
)

type Book struct {
	ID               string         `db:"id"`
	Name             sql.NullString `db:"name"`
	OriginURL        sql.NullString `db:"origin_url"`
	PageCount        sql.NullInt32  `db:"page_count"`
	AttributesParsed bool           `db:"attributes_parsed"`
	CreateAt         time.Time      `db:"create_at"`
	Deleted          bool           `db:"deleted"`
	DeletedAt        sql.NullTime   `db:"deleted_at"`
	Verified         bool           `db:"verified"`
	VerifiedAt       sql.NullTime   `db:"verified_at"`
	IsRebuild        bool           `db:"is_rebuild"`
}

func (b Book) ToEntity() (entities.Book, error) {
	id, err := uuid.Parse(b.ID)
	if err != nil {
		return entities.Book{}, err
	}

	var originURL *url.URL

	if b.OriginURL.Valid {
		originURL, err = url.Parse(b.OriginURL.String)
		if err != nil {
			return entities.Book{}, err
		}
	}

	return entities.Book{
		ID:               id,
		Name:             b.Name.String,
		OriginURL:        originURL,
		PageCount:        int(b.PageCount.Int32),
		AttributesParsed: b.AttributesParsed,
		CreateAt:         b.CreateAt,

		Deleted:    b.Deleted,
		DeletedAt:  b.DeletedAt.Time,
		Verified:   b.Verified,
		VerifiedAt: b.VerifiedAt.Time,

		IsRebuild: b.IsRebuild,
	}, nil
}
