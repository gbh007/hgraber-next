package postgresql

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

// FIXME: добавить лимиты
func (d *Database) NotDownloadedPages(ctx context.Context) ([]core.PageForDownload, error) {
	raw := make([]*model.PageForDownload, 0)

	err := d.db.SelectContext(ctx, &raw, `SELECT p.book_id, b.origin_url AS book_url, p.origin_url AS image_url, p.page_number, p.ext from books AS b INNER JOIN pages AS p ON b.id = p.book_id WHERE downloaded = FALSE;`)
	if err != nil {
		return nil, fmt.Errorf("select: %w", err)
	}

	out := make([]core.PageForDownload, len(raw))
	for i, v := range raw {
		out[i], err = v.ToEntity()
		if err != nil {
			return nil, fmt.Errorf("to entity (%s, %d): %w", v.BookID, v.PageNumber, err)
		}
	}

	return out, nil
}

// FIXME: добавить лимиты
func (d *Database) UnprocessedBooks(ctx context.Context) ([]core.Book, error) {
	raw := make([]*model.Book, 0)

	err := d.db.SelectContext(ctx, &raw, `SELECT * FROM books WHERE (name IS NULL OR page_count IS NULL OR attributes_parsed = FALSE) AND origin_url IS NOT NULL AND deleted = FALSE AND is_rebuild = FALSE;`)
	if err != nil {
		return nil, fmt.Errorf("select: %w", err)
	}

	out := make([]core.Book, len(raw))
	for i, v := range raw {
		out[i], err = v.ToEntity()
		if err != nil {
			return nil, fmt.Errorf("to entity (%s): %w", v.ID, err)
		}
	}

	return out, nil
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

	d.squirrelDebugLog(ctx, query, args)

	result := make([]core.File, 0)

	rows, err := d.pool.Query(ctx, query, args...)
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
