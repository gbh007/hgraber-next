package attribute

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (repo *AttributeRepo) UpdateOriginAttributes(
	ctx context.Context,
	bookID uuid.UUID,
	attributes map[string][]string,
) error {
	tx, err := repo.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}

	defer func() {
		err := tx.Rollback(ctx)
		if err != nil && !errors.Is(err, sql.ErrTxDone) && !errors.Is(err, pgx.ErrTxClosed) {
			repo.Logger.ErrorContext(
				ctx, "rollback UpdateOriginAttributes tx",
				slog.Any("err", err),
			)
		}
	}()

	_, err = tx.Exec(ctx, `DELETE FROM book_origin_attributes WHERE book_id = $1;`, bookID)
	if err != nil {
		return fmt.Errorf("delete old attributes: %w", err)
	}

	builder := squirrel.Insert("book_origin_attributes").
		PlaceholderFormat(squirrel.Dollar).
		Columns(
			"book_id",
			"attr",
			"values",
		)

	for code, values := range attributes {
		builder = builder.Values(bookID, code, values)
	}

	query, args := builder.MustSql()

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
