package model

import (
	"database/sql"
	"net/url"
	"time"

	"github.com/google/uuid"

	"hgnext/internal/entities"
)

// FIXME: проверить работу UUID
type Page struct {
	BookID     string         `db:"book_id"`
	PageNumber int            `db:"page_number"`
	Ext        string         `db:"ext"`
	OriginURL  sql.NullString `db:"origin_url"`
	CreateAt   time.Time      `db:"create_at"`
	Downloaded bool           `db:"downloaded"`
	LoadAt     sql.NullTime   `db:"load_at"`
	FileID     sql.NullString `db:"file_id"`
}

func (p Page) ToEntity() (entities.Page, error) {
	bookID, err := uuid.Parse(p.BookID)
	if err != nil {
		return entities.Page{}, err
	}

	fileID := uuid.Nil

	if p.FileID.Valid {
		fileID, err = uuid.Parse(p.FileID.String)
		if err != nil {
			return entities.Page{}, err
		}
	}

	var originURL *url.URL

	if p.OriginURL.Valid {
		originURL, err = url.Parse(p.OriginURL.String)
		if err != nil {
			return entities.Page{}, err
		}
	}

	return entities.Page{
		BookID:     bookID,
		PageNumber: p.PageNumber,
		Ext:        p.Ext,
		OriginURL:  originURL,
		CreateAt:   p.CreateAt,
		Downloaded: p.Downloaded,
		LoadAt:     p.LoadAt.Time,
		FileID:     fileID,
	}, nil
}
