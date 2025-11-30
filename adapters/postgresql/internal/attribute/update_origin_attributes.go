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

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
)

func (repo *AttributeRepo) UpdateOriginAttributes(
	ctx context.Context,
	bookID uuid.UUID,
	attributes map[string][]string,
) error {
	bookOriginAttributeTable := model.BookOriginAttributeTable

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

	deleteAttrsQuery, deleteAttrsArgs := squirrel.
		Delete(bookOriginAttributeTable.Name()).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			bookOriginAttributeTable.ColumnBookID(): bookID,
		}).
		MustSql()

	_, err = tx.Exec(ctx, deleteAttrsQuery, deleteAttrsArgs...)
	if err != nil {
		return fmt.Errorf("delete old attributes: %w", err)
	}

	builder := squirrel.Insert(bookOriginAttributeTable.Name()).
		PlaceholderFormat(squirrel.Dollar).
		Columns(
			bookOriginAttributeTable.ColumnBookID(),
			bookOriginAttributeTable.ColumnAttr(),
			bookOriginAttributeTable.ColumnValues(),
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
