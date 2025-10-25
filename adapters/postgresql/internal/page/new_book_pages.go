package page

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

func (repo *PageRepo) NewBookPages(ctx context.Context, pages []core.Page) error {
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

func (repo *PageRepo) insertPage(ctx context.Context, tx pgx.Tx, page core.Page) error {
	table := model.PageTable

	query, args := squirrel.Insert(table.Name()).
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			table.ColumnBookID():     page.BookID,
			table.ColumnPageNumber(): page.PageNumber,
			table.ColumnExt():        page.Ext,
			table.ColumnOriginURL():  model.URLToDB(page.OriginURL),
			table.ColumnCreateAt():   page.CreateAt.UTC(),
			table.ColumnDownloaded(): page.Downloaded,
			table.ColumnLoadAt():     model.TimeToDB(page.LoadAt),
			table.ColumnFileID():     model.UUIDToDB(page.FileID),
		}).
		MustSql()

	_, err := tx.Exec(ctx, query, args...)
	if err != nil {
		return err //nolint:wrapcheck // оставляем оригинальную ошибку
	}

	return nil
}
