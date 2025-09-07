package massload

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
)

func (repo *MassloadRepo) DeleteMassloadFlag(ctx context.Context, code string) error {
	builder := squirrel.Delete("massload_flags").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			"code": code,
		})

	query, args := builder.MustSql()

	_, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
