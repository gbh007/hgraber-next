package attribute

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *AttributeRepo) AttributeColors(ctx context.Context) ([]core.AttributeColor, error) {
	builder := squirrel.Select(
		"attr",
		"value",
		"text_color",
		"background_color",
		"created_at",
	).
		From("attribute_colors").
		PlaceholderFormat(squirrel.Dollar)

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

	result := make([]core.AttributeColor, 0, 10) //nolint:mnd // оптимизация

	for rows.Next() {
		color := core.AttributeColor{}

		err = rows.Scan(
			&color.Code,
			&color.Value,
			&color.TextColor,
			&color.BackgroundColor,
			&color.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		result = append(result, color)
	}

	return result, nil
}
