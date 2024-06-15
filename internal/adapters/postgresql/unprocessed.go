package postgresql

import (
	"context"
	"fmt"

	"hgnext/internal/adapters/postgresql/internal/model"
	"hgnext/internal/entities"
)

// FIXME: добавить лимиты
func (d *Database) NotDownloadedPages(ctx context.Context) ([]entities.PageForDownload, error) {
	raw := make([]*model.PageForDownload, 0)

	err := d.db.SelectContext(ctx, &raw, `SELECT p.book_id, b.url AS book_url, p.url AS image_url, p.page_number, p.ext from books AS b INNER JOIN pages AS p ON b.id = p.book_id WHERE downloaded = FALSE;`)
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

	err := d.db.SelectContext(ctx, &raw, `SELECT * FROM books WHERE name IS NULL OR page_count IS NULL OR attributes_parsed = FALSE;`)
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
