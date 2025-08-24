package deadhash

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *DeadHashRepo) DeadHashesByMD5Sums(ctx context.Context, md5Sums []string) ([]core.DeadHash, error) {
	builder := squirrel.Select(
		"md5_sum",
		"sha256_sum",
		"size",
		"created_at",
	).
		PlaceholderFormat(squirrel.Dollar).
		From("dead_hashes").
		Where(squirrel.Eq{
			"md5_sum": md5Sums,
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	repo.SquirrelDebugLog(ctx, query, args)

	rows, err := repo.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	result := make([]core.DeadHash, 0, len(md5Sums))

	for rows.Next() {
		hash := core.DeadHash{}

		err = rows.Scan(
			&hash.Md5Sum,
			&hash.Sha256Sum,
			&hash.Size,
			&hash.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		result = append(result, hash)
	}

	return result, nil
}
