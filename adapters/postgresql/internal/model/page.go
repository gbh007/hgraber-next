package model

import (
	"database/sql"
	"fmt"
	"net/url"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/domain/core"
)

var PageTable Page

type Page struct{}

func (Page) Name() string {
	return "pages"
}

func (Page) ColumnBookID() string     { return "book_id" }
func (Page) ColumnPageNumber() string { return "page_number" }
func (Page) ColumnExt() string        { return "ext" }
func (Page) ColumnOriginURL() string  { return "origin_url" }
func (Page) ColumnCreateAt() string   { return "create_at" }
func (Page) ColumnDownloaded() string { return "downloaded" }
func (Page) ColumnLoadAt() string     { return "load_at" }
func (Page) ColumnFileID() string     { return "file_id" }

func (p Page) Columns() []string {
	return []string{
		p.ColumnBookID(),
		p.ColumnPageNumber(),
		p.ColumnExt(),
		p.ColumnOriginURL(),
		p.ColumnCreateAt(),
		p.ColumnDownloaded(),
		p.ColumnLoadAt(),
		p.ColumnFileID(),
	}
}

func (Page) Scanner(page *core.Page) RowScanner {
	return func(rows pgx.Rows) error {
		var (
			originURL sql.NullString
			loadAt    sql.NullTime
			fileID    uuid.NullUUID
		)

		err := rows.Scan(
			&page.BookID,
			&page.PageNumber,
			&page.Ext,
			&originURL,
			&page.CreateAt,
			&page.Downloaded,
			&loadAt,
			&fileID,
		)
		if err != nil {
			return fmt.Errorf("scan to model: %w", err)
		}

		if originURL.Valid {
			page.OriginURL, err = url.Parse(originURL.String)
			if err != nil {
				return fmt.Errorf("parse origin url: %w", err)
			}
		}

		page.LoadAt = loadAt.Time
		page.FileID = fileID.UUID

		return nil
	}
}
