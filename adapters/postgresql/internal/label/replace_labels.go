package label

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *LabelRepo) ReplaceLabels(ctx context.Context, bookID uuid.UUID, labels []core.BookLabel) error {
	builder := squirrel.Insert("book_labels").
		PlaceholderFormat(squirrel.Dollar).
		Columns(
			"book_id",
			"page_number",
			"name",
			"value",
			"create_at",
		).
		Suffix(`ON CONFLICT (book_id, page_number, name) DO UPDATE SET value = EXCLUDED.value`)

	for _, label := range labels {
		builder = builder.Values(
			bookID,
			label.PageNumber,
			label.Name,
			label.Value,
			label.CreateAt,
		)
	}

	query, args := builder.MustSql()

	tx, err := repo.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}

	defer func() {
		err := tx.Rollback(ctx)
		if err != nil && !errors.Is(err, sql.ErrTxDone) && !errors.Is(err, pgx.ErrTxClosed) {
			repo.Logger.ErrorContext(
				ctx, "rollback ReplaceLabels tx",
				slog.Any("err", err),
			)
		}
	}()

	_, err = tx.Exec(ctx, `DELETE FROM book_labels WHERE book_id = $1;`, bookID)
	if err != nil {
		return fmt.Errorf("delete old labels: %w", err)
	}

	_, err = tx.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}

	return nil
}
