package model

import (
	"database/sql"
	"hgnext/internal/entities"
	"time"

	"github.com/google/uuid"
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

	return entities.Page{
		BookID:     bookID,
		PageNumber: p.PageNumber,
		Ext:        p.Ext,
		OriginURL:  p.OriginURL.String,
		CreateAt:   p.CreateAt,
		Downloaded: p.Downloaded,
		LoadAt:     p.LoadAt.Time,
		FileID:     fileID,
	}, nil
}
