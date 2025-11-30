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

func (repo *AttributeRepo) UpdateAttributes(
	ctx context.Context,
	bookID uuid.UUID,
	attributes map[string][]string,
) error {
	bookAttributeTable := model.BookAttributeTable

	tx, err := repo.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}

	defer func() {
		err := tx.Rollback(ctx)
		if err != nil && !errors.Is(err, sql.ErrTxDone) && !errors.Is(err, pgx.ErrTxClosed) {
			repo.Logger.ErrorContext(
				ctx, "rollback UpdateAttributes tx",
				slog.Any("err", err),
			)
		}
	}()

	deleteAttrsQuery, deleteAttrsArgs := squirrel.
		Delete(bookAttributeTable.Name()).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			bookAttributeTable.ColumnBookID(): bookID,
		}).
		MustSql()

	_, err = tx.Exec(ctx, deleteAttrsQuery, deleteAttrsArgs...)
	if err != nil {
		return fmt.Errorf("delete old attributes: %w", err)
	}

	builder := squirrel.Insert(bookAttributeTable.Name()).
		PlaceholderFormat(squirrel.Dollar).
		Columns(
			bookAttributeTable.ColumnBookID(),
			bookAttributeTable.ColumnAttr(),
			bookAttributeTable.ColumnValue(),
		)

	for code, values := range attributes {
		for _, value := range values {
			builder = builder.Values(bookID, code, value)
		}
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
