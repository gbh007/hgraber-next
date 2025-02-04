package postgresql

import (
	"context"
	"fmt"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/entities"
)

// FIXME: добавить лимиты
func (d *Database) NotDownloadedPages(ctx context.Context) ([]entities.PageForDownload, error) {
	raw := make([]*model.PageForDownload, 0)

	err := d.db.SelectContext(ctx, &raw, `SELECT p.book_id, b.origin_url AS book_url, p.origin_url AS image_url, p.page_number, p.ext from books AS b INNER JOIN pages AS p ON b.id = p.book_id WHERE downloaded = FALSE;`)
	if err != nil {
		return nil, fmt.Errorf("select: %w", err)
	}

	out := make([]entities.PageForDownload, len(raw))
	for i, v := range raw {
		out[i], err = v.ToEntity()
		if err != nil {
			return nil, fmt.Errorf("to entity (%s, %d): %w", v.BookID, v.PageNumber, err)
		}
	}

	return out, nil
}

// FIXME: добавить лимиты
func (d *Database) UnprocessedBooks(ctx context.Context) ([]entities.Book, error) {
	raw := make([]*model.Book, 0)

	err := d.db.SelectContext(ctx, &raw, `SELECT * FROM books WHERE (name IS NULL OR page_count IS NULL OR attributes_parsed = FALSE) AND origin_url IS NOT NULL AND deleted = FALSE AND is_rebuild = FALSE;`)
	if err != nil {
		return nil, fmt.Errorf("select: %w", err)
	}

	out := make([]entities.Book, len(raw))
	for i, v := range raw {
		out[i], err = v.ToEntity()
		if err != nil {
			return nil, fmt.Errorf("to entity (%s): %w", v.ID, err)
		}
	}

	return out, nil
}

// FIXME: добавить лимиты
func (d *Database) GetUnHashedFiles(ctx context.Context) ([]entities.File, error) {
	raw := make([]*model.File, 0)

	err := d.db.SelectContext(ctx, &raw, `SELECT * FROM files WHERE md5_sum IS NULL OR sha256_sum IS NULL OR "size" IS NULL;`)
	if err != nil {
		return nil, fmt.Errorf("exec: %w", err)
	}

	out := make([]entities.File, len(raw))
	for i, v := range raw {
		out[i], err = v.ToEntity()
		if err != nil {
			return nil, fmt.Errorf("convert %s: %w", v.ID, err)
		}
	}

	return out, nil
}
