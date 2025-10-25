package page

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

// TODO: добавить лимиты
func (repo *PageRepo) NotDownloadedPages(ctx context.Context) ([]core.PageForDownload, error) {
	pageTable := model.PageTable

	//nolint:lll // будет исправлено позднее
	builder := squirrel.Select(
		"p."+pageTable.ColumnBookID(),
		"b.origin_url AS book_url",                       // Примечание: ренейминг не нужен для pgx, но оставлен для наглядности.
		"p."+pageTable.ColumnOriginURL()+" AS image_url", // Примечание: ренейминг не нужен для pgx, но оставлен для наглядности.
		"p."+pageTable.ColumnPageNumber(),
		"p."+pageTable.ColumnExt(),
	).
		PlaceholderFormat(squirrel.Dollar).
		From("books AS b").
		InnerJoin(pageTable.Name() + " AS p ON b.id = p." + pageTable.ColumnBookID()).
		Where(squirrel.Eq{
			"p." + pageTable.ColumnDownloaded(): false,
		})

	query, args := builder.MustSql()

	result := make([]core.PageForDownload, 0)

	rows, err := repo.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var (
			page     core.PageForDownload
			bookURL  sql.NullString
			imageURL sql.NullString
		)

		err := rows.Scan(
			&page.BookID,
			&bookURL,
			&imageURL,
			&page.PageNumber,
			&page.Ext,
		)
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		if bookURL.Valid {
			page.BookURL, err = url.Parse(bookURL.String)
			if err != nil {
				return nil, fmt.Errorf("parse book url (%s,%d): %w", page.BookID.String(), page.PageNumber, err)
			}
		}

		if imageURL.Valid {
			page.ImageURL, err = url.Parse(imageURL.String)
			if err != nil {
				return nil, fmt.Errorf("parse page url (%s,%d): %w", page.BookID.String(), page.PageNumber, err)
			}
		}

		result = append(result, page)
	}

	return result, nil
}
