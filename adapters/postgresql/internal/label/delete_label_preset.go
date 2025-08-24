package label

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
)

func (repo *LabelRepo) DeleteLabelPreset(ctx context.Context, name string) error {
	builder := squirrel.Delete("label_presets").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			"name": name,
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
