package model

import (
	"database/sql"
	"fmt"
	"net/url"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/internal/entities"
)

type DeletedPage struct {
	BookID     string         `db:"book_id"`
	PageNumber int            `db:"page_number"`
	Ext        string         `db:"ext"`
	OriginURL  sql.NullString `db:"origin_url"`
	Md5Sum     sql.NullString `db:"md5_sum"`
	Sha256Sum  sql.NullString `db:"sha256_sum"`
	Size       sql.NullInt64  `db:"size"`
	Downloaded bool           `db:"downloaded"`
	CreatedAt  time.Time      `db:"created_at"`
	LoadedAt   sql.NullTime   `db:"loaded_at"`
	DeletedAt  sql.NullTime   `db:"deleted_at"`
}

func DeletedPageToPageWithHashColumns() []string {
	return []string{
		"book_id",
		"page_number",
		"ext",
		"origin_url",
		"downloaded",
		"created_at",
		"loaded_at",
		"md5_sum",
		"sha256_sum",
		"\"size\"",
	}
}

func DeletedPageToPageWithHashScanner(p *entities.PageWithHash) RowScanner {
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
