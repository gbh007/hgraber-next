package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

// FIXME: добавить лимиты
func (d *Database) NotDownloadedPages(ctx context.Context) ([]core.PageForDownload, error) {
	builder := squirrel.Select(
		"p.book_id",
		"b.origin_url AS book_url",  // Примечание: ренейминг не нужен для pgx, но оставлен для наглядности.
		"p.origin_url AS image_url", // Примечание: ренейминг не нужен для pgx, но оставлен для наглядности.
		"p.page_number",
		"p.ext",
	).
		PlaceholderFormat(squirrel.Dollar).
		From("books AS b").
		InnerJoin("pages AS p ON b.id = p.book_id").
		Where(squirrel.Eq{
			"p.downloaded": false,
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	d.SquirrelDebugLog(ctx, query, args)

	result := make([]core.PageForDownload, 0)

	rows, err := d.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query :%w", err)
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

// FIXME: добавить лимиты
func (d *Database) UnprocessedBooks(ctx context.Context) ([]core.Book, error) {
	builder := squirrel.Select(model.BookColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("books").
		Where(squirrel.And{
			squirrel.Or{
				squirrel.Expr(`name IS NULL`),
				squirrel.Expr(`page_count IS NULL`),
				squirrel.Eq{
					"attributes_parsed": false,
				},
			},
			squirrel.Expr(`origin_url IS NOT NULL`),
			squirrel.Eq{
				"deleted":    false,
				"is_rebuild": false,
			},
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	d.SquirrelDebugLog(ctx, query, args)

	result := make([]core.Book, 0)

	rows, err := d.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query :%w", err)
	}

	defer rows.Close()

	for rows.Next() {
		book := core.Book{}

		err := rows.Scan(model.BookScanner(&book))
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		result = append(result, book)
	}

	return result, nil
}

// FIXME: добавить лимиты
func (d *Database) GetUnHashedFiles(ctx context.Context) ([]core.File, error) {
	builder := squirrel.Select(model.FileColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("files").
		Where(squirrel.Or{
			squirrel.Expr(`md5_sum IS NULL`),
			squirrel.Expr(`sha256_sum IS NULL`),
			squirrel.Expr(`"size" IS NULL`),
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("storage: build query: %w", err)
	}

	d.SquirrelDebugLog(ctx, query, args)

	result := make([]core.File, 0)

	rows, err := d.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query :%w", err)
	}

	defer rows.Close()

	for rows.Next() {
		file := core.File{}

		err := rows.Scan(model.FileScanner(&file))
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		result = append(result, file)
	}

	return result, nil
}
