package model

import (
	"database/sql"
	"fmt"
	"net/url"

	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/domain/core"
)

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

func DeletedPageToPageWithHashScanner(p *core.PageWithHash) RowScanner {
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
