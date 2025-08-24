package book

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *BookRepo) MarkBookAsDeleted(ctx context.Context, bookID uuid.UUID) error {
	tx, err := repo.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}

	defer func() {
		err := tx.Rollback(ctx)
		if err != nil && !errors.Is(err, sql.ErrTxDone) && !errors.Is(err, pgx.ErrTxClosed) {
			repo.Logger.ErrorContext(
				ctx, "rollback MarkBookAsDeleted tx",
				slog.Any("err", err),
			)
		}
	}()

	_, err = tx.Exec(ctx, `INSERT INTO
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
    p.book_id = $1;`, bookID)
	if err != nil {
		return fmt.Errorf("copy pages: %w", err)
	}

	_, err = tx.Exec(ctx, `DELETE FROM pages WHERE book_id = $1;`, bookID)
	if err != nil {
		return fmt.Errorf("delete pages: %w", err)
	}

	_, err = tx.Exec(ctx, `DELETE FROM book_attributes WHERE book_id = $1;`, bookID)
	if err != nil {
		return fmt.Errorf("delete attributes: %w", err)
	}

	res, err := tx.Exec(
		ctx,
		`UPDATE books SET deleted_at = $2, deleted = $3 WHERE id = $1;`,
		bookID, time.Now().UTC(), true,
	)
	if err != nil {
		return fmt.Errorf("update book: %w", err)
	}

	if res.RowsAffected() < 1 {
		return core.BookNotFoundError
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}

	return nil
}
