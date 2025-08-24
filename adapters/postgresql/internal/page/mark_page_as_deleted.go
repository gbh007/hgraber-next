package page

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (repo *PageRepo) MarkPageAsDeleted(ctx context.Context, bookID uuid.UUID, pageNumber int) error {
	tx, err := repo.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}

	defer func() {
		err := tx.Rollback(ctx)
		if err != nil && !errors.Is(err, sql.ErrTxDone) && !errors.Is(err, pgx.ErrTxClosed) {
			repo.Logger.ErrorContext(
				ctx, "rollback MarkPageAsDeleted tx",
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
    p.book_id = $1 AND p.page_number = $2;`, bookID, pageNumber)
	if err != nil {
		return fmt.Errorf("copy page: %w", err)
	}

	_, err = tx.Exec(ctx, `DELETE FROM pages WHERE book_id = $1 AND page_number = $2;`, bookID, pageNumber)
	if err != nil {
		return fmt.Errorf("delete page: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}

	return nil
}
