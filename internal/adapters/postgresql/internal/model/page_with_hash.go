package model

import (
	"database/sql"
	"fmt"
	"net/url"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"hgnext/internal/entities"
)

func PageWithHashColumns() []string {
	return []string{
		"p.book_id",
		"p.page_number",
		"p.ext",
		"p.origin_url",
		"p.downloaded",
		"p.file_id",
		"p.create_at",
		"p.load_at",
		"f.md5_sum",
		"f.sha256_sum",
		"f.size",
	}
}

func PageWithHashScanner(p *entities.PageWithHash) RowScanner {
	return func(rows pgx.Rows) error {
		var (
			originURL sql.NullString
			fileID    sql.NullString
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
			&fileID,
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

		if fileID.Valid {
			p.FileID, err = uuid.Parse(fileID.String)
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
