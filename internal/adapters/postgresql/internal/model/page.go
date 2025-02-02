package model

import (
	"database/sql"
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/internal/entities"
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

type PageForDownload struct {
	BookID     string         `db:"book_id"`
	PageNumber int            `db:"page_number"`
	Ext        string         `db:"ext"`
	BookURL    sql.NullString `db:"book_url"`
	ImageURL   sql.NullString `db:"image_url"`
}

func (p PageForDownload) ToEntity() (entities.PageForDownload, error) {
	bookID, err := uuid.Parse(p.BookID)
	if err != nil {
		return entities.PageForDownload{}, err
	}

	var bookURL *url.URL

	if p.BookURL.Valid {
		bookURL, err = url.Parse(p.BookURL.String)
		if err != nil {
			return entities.PageForDownload{}, err
		}
	}

	var imageURL *url.URL

	if p.ImageURL.Valid {
		imageURL, err = url.Parse(p.ImageURL.String)
		if err != nil {
			return entities.PageForDownload{}, err
		}
	}

	return entities.PageForDownload{
		BookID:     bookID,
		PageNumber: p.PageNumber,
		Ext:        p.Ext,
		BookURL:    bookURL,
		ImageURL:   imageURL,
	}, nil
}
