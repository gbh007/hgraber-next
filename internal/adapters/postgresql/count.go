package postgresql

import (
	"context"
	"database/sql"
	"fmt"

	"hgnext/internal/entities"
)

func (d *Database) SystemSize(ctx context.Context) (entities.SystemSizeInfo, error) {
	systemSize := entities.SystemSizeInfo{}

	// TODO: удалить старую версию если не будет проблем
	// err := d.db.GetContext(ctx, &systemSize.BookCount, `SELECT COUNT(*) FROM books;`)
	err := d.db.GetContext(ctx, &systemSize.BookCount, `SELECT reltuples FROM pg_class WHERE relname = 'books';`)
	if err != nil {
		return entities.SystemSizeInfo{}, fmt.Errorf("get book count : %w", err)
	}

	err = d.db.GetContext(ctx, &systemSize.BookUnparsedCount, `SELECT COUNT(*) FROM books WHERE (name IS NULL OR page_count IS NULL OR attributes_parsed = FALSE) AND origin_url IS NOT NULL AND deleted = FALSE;`)
	if err != nil {
		return entities.SystemSizeInfo{}, fmt.Errorf("get book unparsed count: %w", err)
	}

	// TODO: удалить старую версию если не будет проблем
	// err = d.db.GetContext(ctx, &systemSize.PageCount, `SELECT COUNT(*) FROM pages;`)
	err = d.db.GetContext(ctx, &systemSize.PageCount, `SELECT reltuples FROM pg_class WHERE relname = 'pages';`)
	if err != nil {
		return entities.SystemSizeInfo{}, fmt.Errorf("get page count: %w", err)
	}

	err = d.db.GetContext(ctx, &systemSize.PageUnloadedCount, `SELECT COUNT(*) FROM pages WHERE downloaded = FALSE;`)
	if err != nil {
		return entities.SystemSizeInfo{}, fmt.Errorf("get unloaded page count: %w", err)
	}

	err = d.db.GetContext(ctx, &systemSize.PageWithoutBodyCount, `SELECT COUNT(*) FROM pages WHERE file_id IS NULL;`)
	if err != nil {
		return entities.SystemSizeInfo{}, fmt.Errorf("get page without body count: %w", err)
	}

	systemSize.PageFileSize = d.cachePageFileSize.Load()

	// Оптимизация запросов в БД
	if systemSize.PageFileSize == 0 {
		size := sql.NullInt64{}

		err = d.db.GetContext(ctx, &size, `SELECT SUM(f."size") FROM pages AS p LEFT JOIN files AS f ON p.file_id = f.id WHERE f."size" IS NOT NULL;`)
		if err != nil {
			return entities.SystemSizeInfo{}, fmt.Errorf("get page file size: %w", err)
		}

		systemSize.PageFileSize = size.Int64
		d.cachePageFileSize.Store(size.Int64)
	}

	systemSize.FileSize = d.cacheFileSize.Load()

	// Оптимизация запросов в БД
	if systemSize.FileSize == 0 {
		size := sql.NullInt64{}

		err = d.db.GetContext(ctx, &size, `SELECT SUM("size") FROM files WHERE "size" IS NOT NULL;`)
		if err != nil {
			return entities.SystemSizeInfo{}, fmt.Errorf("get file size: %w", err)
		}

		systemSize.FileSize = size.Int64
		d.cacheFileSize.Store(size.Int64)
	}

	return systemSize, nil
}
