package page

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
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *PageRepo) UpdateBookPages(ctx context.Context, id uuid.UUID, pages []core.Page) error {
	table := model.PageTable

	tx, err := repo.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}

	defer func() {
		err := tx.Rollback(ctx)
		if err != nil && !errors.Is(err, sql.ErrTxDone) && !errors.Is(err, pgx.ErrTxClosed) {
			repo.Logger.ErrorContext(
				ctx, "rollback UpdateBookPages tx",
				slog.Any("err", err),
			)
		}
	}()

	deleteQuery, deleteArgs := squirrel.
		Delete(table.Name()).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			table.ColumnBookID(): id,
		}).
		MustSql()

	_, err = tx.Exec(ctx, deleteQuery, deleteArgs...)
	if err != nil {
		return fmt.Errorf("delete old pages: %w", err)
	}

	for _, v := range pages {
		err = repo.insertPage(ctx, tx, v)
		if err != nil {
			return fmt.Errorf("insert page %d: %w", v.PageNumber, err)
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}

	return nil
}
