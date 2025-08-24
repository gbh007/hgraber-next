package massload

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
)

func (repo *MassloadRepo) DeleteMassload(ctx context.Context, id int) error {
	builder := squirrel.Delete("massloads").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			"id": id,
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
