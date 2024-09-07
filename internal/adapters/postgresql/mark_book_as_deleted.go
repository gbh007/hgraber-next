package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"

	"hgnext/internal/entities"
)

func (d *Database) MarkBookAsDeleted(ctx context.Context, bookID uuid.UUID) error {
	tx, err := d.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}

	defer func() {
		err := tx.Rollback()
		if err != nil && !errors.Is(err, sql.ErrTxDone) {
			d.logger.ErrorContext(
				ctx, "rollback MarkBookAsDeleted tx",
				slog.Any("err", err),
			)
		}
	}()

	_, err = tx.ExecContext(ctx, `INSERT INTO
    deleted_pages
SELECT
    p.book_id,
    p.page_number,
    p.ext,
    p.origin_url,
    f.md5_sum,
    f.sha256_sum,
    f.size,
    p.downloaded,
    p.create_at AS created_at,
    p.load_at AS loaded_at
FROM pages p
    LEFT JOIN files f ON p.file_id = f.id
WHERE
    p.book_id = $1;`, bookID.String())
	if err != nil {
		return fmt.Errorf("copy pages: %w", err)
	}

	_, err = tx.ExecContext(ctx, `DELETE FROM pages WHERE book_id = $1;`, bookID.String())
	if err != nil {
		return fmt.Errorf("delete pages: %w", err)
	}

	_, err = tx.ExecContext(ctx, `DELETE FROM book_attributes WHERE book_id = $1;`, bookID.String())
	if err != nil {
		return fmt.Errorf("delete attributes: %w", err)
	}

	res, err := tx.ExecContext(
		ctx,
		`UPDATE books SET deleted_at = $2, deleted = $3 WHERE id = $1;`,
		bookID.String(), time.Now().UTC(), true,
	)
	if err != nil {
		return fmt.Errorf("update book: %w", err)
	}

	if !d.isApply(ctx, res) {
		return entities.BookNotFoundError
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}

	// Состояние размера изменилось, сбрасываем кеши.
	d.cachePageFileSize.Store(0)
	d.cacheFileSize.Store(0)

	return nil
}