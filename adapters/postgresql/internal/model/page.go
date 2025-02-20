package model

import (
	"database/sql"
	"fmt"
	"net/url"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/domain/core"
)

func PageColumns() []string {
	return []string{
		"book_id",
		"page_number",
		"ext",
		"origin_url",
		"create_at",
		"downloaded",
		"load_at",
		"file_id",
	}
}

func PageScanner(page *core.Page) RowScanner {
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
