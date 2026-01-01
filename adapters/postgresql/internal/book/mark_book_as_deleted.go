package book

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *BookRepo) MarkBookAsDeleted(ctx context.Context, bookID uuid.UUID) error {
	pageTable := model.PageTable.WithPrefix("p")
	fileTable := model.FileTable.WithPrefix("f")
	deletedPageTable := model.DeletedPageTable
	bookTable := model.BookTable
	bookAttributeTable := model.BookAttributeTable

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
					pageTable.ColumnBookID(): bookID,
				}),
		).
		MustSql()

	_, err = tx.Exec(ctx, insertQuery, insertArgs...)
	if err != nil {
		return fmt.Errorf("copy pages: %w", err)
	}

	deletePagesQuery, deletePagesArgs := squirrel.
		Delete(pageTable.NameAlter()).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			pageTable.ColumnBookID(): bookID,
		}).
		MustSql()

	_, err = tx.Exec(ctx, deletePagesQuery, deletePagesArgs...)
	if err != nil {
		return fmt.Errorf("delete pages: %w", err)
	}

	deleteAttrsQuery, deleteAttrsArgs := squirrel.
		Delete(bookAttributeTable.Name()).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			bookAttributeTable.ColumnBookID(): bookID,
		}).
		MustSql()

	_, err = tx.Exec(ctx, deleteAttrsQuery, deleteAttrsArgs...)
	if err != nil {
		return fmt.Errorf("delete attributes: %w", err)
	}

	updateBookQuery, updateBookArgs := squirrel.Update(bookTable.Name()).
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			bookTable.ColumnDeleted():   true,
			bookTable.ColumnDeletedAt(): time.Now().UTC(),
		}).
		Where(squirrel.Eq{
			bookTable.ColumnID(): bookID,
		}).
		MustSql()

	res, err := tx.Exec(ctx, updateBookQuery, updateBookArgs...)
	if err != nil {
		return fmt.Errorf("update book: %w", err)
	}

	if res.RowsAffected() < 1 {
		return core.ErrBookNotFound
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}

	return nil
}
