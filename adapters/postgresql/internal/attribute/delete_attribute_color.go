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

	query, args := builder.MustSql()

	_, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
