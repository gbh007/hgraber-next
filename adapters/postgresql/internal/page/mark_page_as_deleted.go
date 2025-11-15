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
	pageTable := model.PageTable.WithPrefix("p")
	fileTable := model.FileTable.WithPrefix("f")
	deletedPageTable := model.DeletedPageTable

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

	insertQuery, insertArgs := squirrel.Insert(deletedPageTable.Name()).
		PlaceholderFormat(squirrel.Dollar).
		Select(
			squirrel.Select(
				pageTable.ColumnBookID()+" AS "+deletedPageTable.ColumnBookID(),
				pageTable.ColumnPageNumber()+" AS "+deletedPageTable.ColumnPageNumber(),
				pageTable.ColumnExt()+" AS "+deletedPageTable.ColumnExt(),
				pageTable.ColumnOriginURL()+" AS "+deletedPageTable.ColumnOriginURL(),
				fileTable.ColumnMd5Sum()+" AS "+deletedPageTable.ColumnMd5Sum(),
				fileTable.ColumnSha256Sum()+" AS "+deletedPageTable.ColumnSha256Sum(),
				fileTable.ColumnSize()+" AS "+deletedPageTable.ColumnSize(),
				pageTable.ColumnDownloaded()+" AS "+deletedPageTable.ColumnDownloaded(),
				pageTable.ColumnCreateAt()+" AS "+deletedPageTable.ColumnCreatedAt(),
				pageTable.ColumnLoadAt()+" AS "+deletedPageTable.ColumnLoadedAt(),
			).
				From(pageTable.NameAlter()).
				LeftJoin(model.JoinPageAndFile(pageTable, fileTable)).
				Where(squirrel.Eq{
					pageTable.ColumnBookID():     bookID,
					pageTable.ColumnPageNumber(): pageNumber,
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
