package deadhash

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/pkg"
)

func (repo *DeadHashRepo) SetDeadHashes(ctx context.Context, hashes []core.DeadHash) error {
	batches := pkg.Batching(hashes, 5000) //nolint:mnd // отдельная константа не нужна

	for _, batch := range batches {
		builder := squirrel.Insert("dead_hashes").
			PlaceholderFormat(squirrel.Dollar).
			Columns(
				"md5_sum",
				"sha256_sum",
				"size",
				"created_at",
			).
			Suffix(`ON CONFLICT DO NOTHING`)

		for _, hash := range batch {
			builder = builder.Values(hash.Md5Sum,
				hash.Sha256Sum,
				hash.Size,
				hash.CreatedAt,
			)
		}

		query, args := builder.MustSql()

		_, err := repo.Pool.Exec(ctx, query, args...)
		if err != nil {
			return fmt.Errorf("exec query: %w", err)
		}
	}

	return nil
}
