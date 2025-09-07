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

	query, args := builder.MustSql()

	_, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
