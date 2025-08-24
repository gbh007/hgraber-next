package attribute

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
)

func (repo *AttributeRepo) DeleteAttributeColor(ctx context.Context, code, value string) error {
	builder := squirrel.Delete("attribute_colors").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			"attr":  code,
			"value": value,
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
