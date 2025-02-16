package model

import (
	"database/sql"
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/core"
)

type Page struct {
	BookID     uuid.UUID      `db:"book_id"`
	PageNumber int            `db:"page_number"`
	Ext        string         `db:"ext"`
	OriginURL  sql.NullString `db:"origin_url"`
	CreateAt   time.Time      `db:"create_at"`
	Downloaded bool           `db:"downloaded"`
	LoadAt     sql.NullTime   `db:"load_at"`
	FileID     uuid.NullUUID  `db:"file_id"`
}

func (p Page) ToEntity() (core.Page, error) {
	var (
		originURL *url.URL
		err       error
	)

	if p.OriginURL.Valid {
		originURL, err = url.Parse(p.OriginURL.String)
		if err != nil {
			return core.Page{}, err
		}
	}

	return core.Page{
		BookID:     p.BookID,
		PageNumber: p.PageNumber,
		Ext:        p.Ext,
		OriginURL:  originURL,
		CreateAt:   p.CreateAt,
		Downloaded: p.Downloaded,
		LoadAt:     p.LoadAt.Time,
		FileID:     p.FileID.UUID,
	}, nil
}

type PageForDownload struct {
	BookID     uuid.UUID      `db:"book_id"`
	PageNumber int            `db:"page_number"`
	Ext        string         `db:"ext"`
	BookURL    sql.NullString `db:"book_url"`
	ImageURL   sql.NullString `db:"image_url"`
}

func (p PageForDownload) ToEntity() (core.PageForDownload, error) {
	var (
		bookURL *url.URL
		err     error
	)

	if p.BookURL.Valid {
		bookURL, err = url.Parse(p.BookURL.String)
		if err != nil {
			return core.PageForDownload{}, err
		}
	}

	var imageURL *url.URL

	if p.ImageURL.Valid {
		imageURL, err = url.Parse(p.ImageURL.String)
		if err != nil {
			return core.PageForDownload{}, err
		}
	}

	return core.PageForDownload{
		BookID:     p.BookID,
		PageNumber: p.PageNumber,
		Ext:        p.Ext,
		BookURL:    bookURL,
		ImageURL:   imageURL,
	}, nil
}
