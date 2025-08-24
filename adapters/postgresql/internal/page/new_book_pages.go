package page

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

// TODO: отрефакторить на squirel
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

	// TODO: слить с аналогичным дейтвием, реализовать как приватную функцию которая принимает транзакцию.
	for _, v := range pages {
		_, err = tx.Exec(
			ctx,
			`INSERT INTO pages (book_id, page_number, ext, origin_url, create_at, downloaded, load_at, file_id) VALUES($1, $2, $3, $4, $5, $6, $7, $8);`, //nolint:lll // будет исправлено позднее
			v.BookID,
			v.PageNumber,
			v.Ext,
			model.URLToDB(v.OriginURL),
			v.CreateAt.UTC(),
			v.Downloaded,
			model.TimeToDB(v.LoadAt),
			model.UUIDToDB(v.FileID),
		)
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
