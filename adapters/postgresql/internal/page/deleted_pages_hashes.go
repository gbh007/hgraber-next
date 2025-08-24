package page

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *PageRepo) DeletedPagesHashes(ctx context.Context) ([]core.FileHash, error) {
	builder := squirrel.Select(
		"md5_sum",
		"sha256_sum",
		"size",
	).
		PlaceholderFormat(squirrel.Dollar).
		From("deleted_pages").
		Where(squirrel.And{
			squirrel.Expr(`md5_sum IS NOT NULL`),
			squirrel.Expr(`sha256_sum IS NOT NULL`),
			squirrel.Expr(`size IS NOT NULL`),
		}).
		GroupBy(
			"md5_sum",
			"sha256_sum",
			"size",
		)

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

	result := make([]core.FileHash, 0, 100) //nolint:mnd // оптимизация

	for rows.Next() {
		hash := core.FileHash{}

		err = rows.Scan(
			&hash.Md5Sum,
			&hash.Sha256Sum,
			&hash.Size,
		)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		result = append(result, hash)
	}

	return result, nil
}
