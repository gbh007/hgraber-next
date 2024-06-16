package postgresql

import (
	"context"
	"fmt"

	"hgnext/internal/entities"
)

func (d *Database) SystemSize(ctx context.Context) (entities.SystemSizeInfo, error) {
	systemSize := entities.SystemSizeInfo{}

	err := d.db.GetContext(ctx, &systemSize.BookCount, `SELECT COUNT(*) FROM books;`)
	if err != nil {
		return entities.SystemSizeInfo{}, fmt.Errorf("get book count : %w", err)
	}

	err = d.db.GetContext(ctx, &systemSize.BookUnparsedCount, `SELECT COUNT(*) FROM books WHERE (name IS NULL OR page_count IS NULL OR attributes_parsed = FALSE) AND origin_url IS NOT NULL;`)
	if err != nil {
		return entities.SystemSizeInfo{}, fmt.Errorf("get book unparsed count: %w", err)
	}

	err = d.db.GetContext(ctx, &systemSize.PageCount, `SELECT COUNT(*) FROM pages;`)
	if err != nil {
		return entities.SystemSizeInfo{}, fmt.Errorf("get page count: %w", err)
	}

	err = d.db.GetContext(ctx, &systemSize.PageUnloadedCount, `SELECT COUNT(*) FROM pages WHERE downloaded = FALSE;`)
	if err != nil {
		return entities.SystemSizeInfo{}, fmt.Errorf("get unloaded page count: %w", err)
	}

	err = d.db.GetContext(ctx, &systemSize.PageFileSize, `SELECT SUM(f."size") FROM pages AS p LEFT JOIN files AS f ON p.file_id = f.id WHERE f."size" IS NOT NULL;`)
	if err != nil {
		return entities.SystemSizeInfo{}, fmt.Errorf("get page file size: %w", err)
	}

	return systemSize, nil
}

func (d *Database) BookCount(ctx context.Context) (int, error) {
	var c int

	err := d.db.GetContext(ctx, &c, `SELECT COUNT(*) FROM books;`)
	if err != nil {
		return 0, err
	}

	return c, nil
}
