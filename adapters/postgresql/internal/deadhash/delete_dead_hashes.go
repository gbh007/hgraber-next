package deadhash

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *DeadHashRepo) DeleteDeadHashes(ctx context.Context, hashes []core.DeadHash) error {
	table := model.DeadHashTable

	batch := &pgx.Batch{}
	resultCount := 0

	for _, hash := range hashes {
		builder := squirrel.Delete(table.Name()).
			PlaceholderFormat(squirrel.Dollar).
			Where(squirrel.Eq{
				table.ColumnMd5Sum():    hash.Md5Sum,
				table.ColumnSha256Sum(): hash.Sha256Sum,
				table.ColumnSize():      hash.Size,
			})

		query, args := builder.MustSql()
		batch.Queue(query, args...)

		resultCount++
	}

	batchResult := repo.Pool.SendBatch(ctx, batch)

	defer func() {
		err := batchResult.Close()
		if err != nil {
			repo.Logger.ErrorContext(ctx, "close DeleteDeadHashes batch", slog.String("err", err.Error()))
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
