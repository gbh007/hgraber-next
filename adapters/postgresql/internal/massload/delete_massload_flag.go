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
