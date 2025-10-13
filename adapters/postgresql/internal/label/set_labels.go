package label

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *LabelRepo) SetLabels(ctx context.Context, labels []core.BookLabel) error {
	table := model.BookLabelTable

	builder := squirrel.Insert(table.Name()).
		PlaceholderFormat(squirrel.Dollar).
		Columns(
			table.ColumnBookID(),
			table.ColumnPageNumber(),
			table.ColumnName(),
			table.ColumnValue(),
			table.ColumnCreateAt(),
		).
		Suffix(fmt.Sprintf(
			`ON CONFLICT (%s, %s, %s) DO UPDATE SET %s = EXCLUDED.%s`,
			table.ColumnBookID(),
			table.ColumnPageNumber(),
			table.ColumnName(),
			table.ColumnValue(),
			table.ColumnValue(),
		))

	for _, label := range labels {
		builder = builder.Values(
			label.BookID,
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
