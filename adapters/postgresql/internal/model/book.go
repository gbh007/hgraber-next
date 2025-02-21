package model

import (
	"database/sql"
	"fmt"
	"net/url"

	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/domain/core"
)

func BookColumns() []string {
	return []string{
		"id",
		"name",
		"origin_url",
		"page_count",
		"attributes_parsed",
		"create_at",
		"deleted",
		"deleted_at",
		"verified",
		"verified_at",
		"is_rebuild",
	}
}

func BookScanner(book *core.Book) RowScanner {
	return func(rows pgx.Rows) error {
		var (
			name       sql.NullString
			originURL  sql.NullString
			pageCount  sql.NullInt32
			deletedAt  sql.NullTime
			verifiedAt sql.NullTime
		)

		err := rows.Scan(
			&book.ID,
			&name,
			&originURL,
			&pageCount,
			&book.AttributesParsed,
			&book.CreateAt,
			&book.Deleted,
			&deletedAt,
			&book.Verified,
			&verifiedAt,
			&book.IsRebuild,
		)
		if err != nil {
			return fmt.Errorf("scan to model: %w", err)
		}

		if originURL.Valid {
			book.OriginURL, err = url.Parse(originURL.String)
			if err != nil {
				return fmt.Errorf("parse origin url: %w", err)
			}
		}

		book.Name = name.String
		book.PageCount = int(pageCount.Int32)
		book.DeletedAt = deletedAt.Time
		book.VerifiedAt = verifiedAt.Time

		return nil
	}
}
