package deadhash

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *DeadHashRepo) SetDeadHash(ctx context.Context, hash core.DeadHash) error {
	builder := squirrel.Insert("dead_hashes").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			"md5_sum":    hash.Md5Sum,
			"sha256_sum": hash.Sha256Sum,
			"size":       hash.Size,
			"created_at": hash.CreatedAt,
		}).
		Suffix(`ON CONFLICT DO NOTHING`)

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	repo.SquirrelDebugLog(ctx, query, args)

	_, err = repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
