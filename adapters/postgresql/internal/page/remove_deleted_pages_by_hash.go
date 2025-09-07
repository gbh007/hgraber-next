package page

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *PageRepo) RemoveDeletedPagesByHash(ctx context.Context, hash core.FileHash) error {
	builder := squirrel.Delete("deleted_pages").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			"md5_sum":    hash.Md5Sum,
			"sha256_sum": hash.Sha256Sum,
			"\"size\"":   hash.Size,
		})

	query, args := builder.MustSql()

	_, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
