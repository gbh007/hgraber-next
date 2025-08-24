package page

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *PageRepo) RemoveDeletedPagesByHashes(ctx context.Context, hashes []core.FileHash) error {
	batch := &pgx.Batch{}

	resultCount := 0

	for _, hash := range hashes {
		builder := squirrel.Delete("deleted_pages").
			PlaceholderFormat(squirrel.Dollar).
			Where(squirrel.Eq{
				"md5_sum":    hash.Md5Sum,
				"sha256_sum": hash.Sha256Sum,
				"\"size\"":   hash.Size,
			})

		query, args, err := builder.ToSql()
		if err != nil {
			return fmt.Errorf("build query: %w", err)
		}

		repo.SquirrelDebugLog(ctx, query, args)
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
