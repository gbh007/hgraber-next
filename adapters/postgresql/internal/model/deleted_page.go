package model

import (
	"database/sql"
	"fmt"
	"net/url"

	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/domain/core"
)

var DeletedPageTable = DeletedPage{baseTable: baseTable{name: "deleted_pages"}}

type DeletedPage struct {
	baseTable
}

func (p DeletedPage) WithPrefix(pf string) DeletedPage {
	return DeletedPage{
		baseTable: p.withPrefix(pf),
	}
}

func (p DeletedPage) ColumnBookID() string     { return "book_id" }
func (p DeletedPage) ColumnPageNumber() string { return "page_number" }
func (p DeletedPage) ColumnExt() string        { return "ext" }
func (p DeletedPage) ColumnOriginURL() string  { return "origin_url" }
func (p DeletedPage) ColumnDownloaded() string { return "downloaded" }
func (p DeletedPage) ColumnCreatedAt() string  { return "created_at" }
func (p DeletedPage) ColumnLoadedAt() string   { return "loaded_at" }
func (p DeletedPage) ColumnMd5Sum() string     { return "md5_sum" }
func (p DeletedPage) ColumnSha256Sum() string  { return "sha256_sum" }
func (p DeletedPage) ColumnSize() string       { return "\"size\"" }

func (p DeletedPage) ToPageWithHashColumns() []string {
	return []string{
		p.ColumnBookID(),
		p.ColumnPageNumber(),
		p.ColumnExt(),
		p.ColumnOriginURL(),
		p.ColumnDownloaded(),
		p.ColumnCreatedAt(),
		p.ColumnLoadedAt(),
		p.ColumnMd5Sum(),
		p.ColumnSha256Sum(),
		p.ColumnSize(),
	}
}

func (DeletedPage) ToPageWithHashScanner(p *core.PageWithHash) RowScanner {
	return func(rows pgx.Rows) error {
		var (
			originURL sql.NullString
			loadAt    sql.NullTime
			md5Sum    sql.NullString
			sha256Sum sql.NullString
			size      sql.NullInt64
		)

		err := rows.Scan(
			&p.BookID,
			&p.PageNumber,
			&p.Ext,
			&originURL,
			&p.Downloaded,
			&p.CreateAt,
			&loadAt,
			&md5Sum,
			&sha256Sum,
			&size,
		)
		if err != nil {
			return fmt.Errorf("scan to model: %w", err)
		}

		if originURL.Valid {
			p.OriginURL, err = url.Parse(originURL.String)
			if err != nil {
				return fmt.Errorf("convert to entity: %w", err)
			}
		}

		p.LoadAt = loadAt.Time
		p.Md5Sum = md5Sum.String
		p.Sha256Sum = sha256Sum.String
		p.Size = size.Int64

		return nil
	}
}
