package model

import (
	"database/sql"
	"fmt"
	"net/url"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/domain/core"
)

var PageTable = Page{baseTable: baseTable{name: "pages"}}

type Page struct {
	baseTable
}

func (p Page) WithPrefix(pf string) Page {
	return Page{
		baseTable: p.withPrefix(pf),
	}
}

func (p Page) ColumnBookID() string     { return p.prefix + "book_id" }
func (p Page) ColumnPageNumber() string { return p.prefix + "page_number" }
func (p Page) ColumnExt() string        { return p.prefix + "ext" }
func (p Page) ColumnOriginURL() string  { return p.prefix + "origin_url" }
func (p Page) ColumnCreateAt() string   { return p.prefix + "create_at" }
func (p Page) ColumnDownloaded() string { return p.prefix + "downloaded" }
func (p Page) ColumnLoadAt() string     { return p.prefix + "load_at" }
func (p Page) ColumnFileID() string     { return p.prefix + "file_id" }

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
