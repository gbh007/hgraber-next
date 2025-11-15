package page

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *PageRepo) RemoveDeletedPagesByHashes(ctx context.Context, hashes []core.FileHash) error {
	deletedPageTable := model.DeletedPageTable
	batch := &pgx.Batch{}

	resultCount := 0

	for _, hash := range hashes {
		builder := squirrel.Delete(deletedPageTable.Name()).
			PlaceholderFormat(squirrel.Dollar).
			Where(squirrel.Eq{
				deletedPageTable.ColumnMd5Sum():    hash.Md5Sum,
				deletedPageTable.ColumnSha256Sum(): hash.Sha256Sum,
				deletedPageTable.ColumnSize():      hash.Size,
			})

		query, args := builder.MustSql()

		batch.Queue(query, args...)

		resultCount++
	}

	batchResult := repo.Pool.SendBatch(ctx, batch)

	defer func() {
		err := batchResult.Close()
		if err != nil {
			repo.Logger.ErrorContext(ctx, "close RemoveDeletedPagesByHashes batch", slog.String("err", err.Error()))
		}
	}()

	for range resultCount {
		_, err := batchResult.Exec()
		if err != nil {
			return fmt.Errorf("exec query: %w", err)
		}
	}

	return nil
}
