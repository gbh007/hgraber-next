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
)

func (repo *PageRepo) MarkPageAsDeleted(ctx context.Context, bookID uuid.UUID, pageNumber int) error {
	pageTable := model.PageTable

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

	insertQuery, insertArgs := squirrel.Insert("deleted_pages").
		PlaceholderFormat(squirrel.Dollar).
		Select(
			squirrel.Select(
				"p."+pageTable.ColumnBookID(),
				"p."+pageTable.ColumnPageNumber(),
				"p."+pageTable.ColumnExt(),
				"p."+pageTable.ColumnOriginURL(),
				"f.md5_sum",
				"f.sha256_sum",
				"f.size",
				"p."+pageTable.ColumnDownloaded(),
				"p."+pageTable.ColumnCreateAt()+" AS created_at",
				"p."+pageTable.ColumnLoadAt()+" AS loaded_at",
			).
				From(pageTable.Name() + " p").
				LeftJoin("files f ON p." + pageTable.ColumnFileID() + " = f.id").
				Where(squirrel.Eq{
					"p." + pageTable.ColumnBookID():     bookID,
					"p." + pageTable.ColumnPageNumber(): pageNumber,
				}),
		).
		MustSql()

	_, err = tx.Exec(ctx, insertQuery, insertArgs...)
	if err != nil {
		return fmt.Errorf("copy page: %w", err)
	}

	deleteQuery, deleteArgs := squirrel.
		Delete(pageTable.Name()).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			pageTable.ColumnBookID():     bookID,
			pageTable.ColumnPageNumber(): pageNumber,
		}).
		MustSql()

	_, err = tx.Exec(ctx, deleteQuery, deleteArgs...)
	if err != nil {
		return fmt.Errorf("delete page: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}

	return nil
}
