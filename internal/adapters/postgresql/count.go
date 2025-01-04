package postgresql

import (
	"context"
	"database/sql"
	"fmt"

	"hgnext/internal/entities"
)

func (d *Database) SystemSize(ctx context.Context) (entities.SystemSizeInfo, error) {
	systemSize := entities.SystemSizeInfo{}

	var err error

	err = d.db.GetContext(ctx, &systemSize.BookCount, `SELECT COUNT(*) FROM books;`)
	if err != nil {
		return entities.SystemSizeInfo{}, fmt.Errorf("get book count : %w", err)
	}

	err = d.db.GetContext(ctx, &systemSize.DownloadedBookCount, `SELECT COUNT(*) FROM books WHERE deleted = FALSE AND page_count IS NOT NULL AND NOT EXISTS (SELECT 1 FROM pages WHERE book_id = books.id AND pages.downloaded = FALSE);`)
	if err != nil {
		return entities.SystemSizeInfo{}, fmt.Errorf("get downloaded book count : %w", err)
	}

	err = d.db.GetContext(ctx, &systemSize.VerifiedBookCount, `SELECT COUNT(*) FROM books WHERE deleted = FALSE AND verified = TRUE AND page_count IS NOT NULL AND NOT EXISTS (SELECT 1 FROM pages WHERE book_id = books.id AND pages.downloaded = FALSE);`)
	if err != nil {
		return entities.SystemSizeInfo{}, fmt.Errorf("get book verified count : %w", err)
	}

	err = d.db.GetContext(ctx, &systemSize.RebuildedBookCount, `SELECT COUNT(*) FROM books WHERE is_rebuild = TRUE;`)
	if err != nil {
		return entities.SystemSizeInfo{}, fmt.Errorf("get book rebuilded count: %w", err)
	}

	err = d.db.GetContext(ctx, &systemSize.BookUnparsedCount, `SELECT COUNT(*) FROM books WHERE (name IS NULL OR page_count IS NULL OR attributes_parsed = FALSE) AND origin_url IS NOT NULL AND deleted = FALSE AND is_rebuild = FALSE;`)
	if err != nil {
		return entities.SystemSizeInfo{}, fmt.Errorf("get book unparsed count: %w", err)
	}

	err = d.db.GetContext(ctx, &systemSize.DeletedBookCount, `SELECT COUNT(*) FROM books WHERE deleted = TRUE;`)
	if err != nil {
		return entities.SystemSizeInfo{}, fmt.Errorf("get book deleted count: %w", err)
	}

	err = d.db.GetContext(ctx, &systemSize.DeadHashCount, `SELECT COUNT(*) FROM dead_hashes;`)
	if err != nil {
		return entities.SystemSizeInfo{}, fmt.Errorf("get dead hash count: %w", err)
	}

	err = d.db.GetContext(ctx, &systemSize.PageCount, `SELECT COUNT(*) FROM pages;`)
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

	err = d.db.GetContext(ctx, &systemSize.DeletedPageCount, `SELECT COUNT(*) FROM deleted_pages;`)
	if err != nil {
		return entities.SystemSizeInfo{}, fmt.Errorf("get deleted page count: %w", err)
	}

	size := sql.NullInt64{}

	err = d.db.GetContext(ctx, &size, `SELECT SUM(f."size") FROM pages AS p LEFT JOIN files AS f ON p.file_id = f.id WHERE f."size" IS NOT NULL;`)
	if err != nil {
		return entities.SystemSizeInfo{}, fmt.Errorf("get page file size: %w", err)
	}

	systemSize.PageFileSize = size.Int64

	size = sql.NullInt64{}

	err = d.db.GetContext(ctx, &size, `SELECT SUM("size") FROM files WHERE "size" IS NOT NULL;`)
	if err != nil {
		return entities.SystemSizeInfo{}, fmt.Errorf("get file size: %w", err)
	}

	systemSize.FileSize = size.Int64

	return systemSize, nil
}
