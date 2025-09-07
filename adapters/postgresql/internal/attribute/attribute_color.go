package attribute

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *AttributeRepo) AttributeColor(ctx context.Context, code, value string) (core.AttributeColor, error) {
	builder := squirrel.Select(
		"attr",
		"value",
		"text_color",
		"background_color",
		"created_at",
	).
		PlaceholderFormat(squirrel.Dollar).
		From("attribute_colors").
		Where(squirrel.Eq{
			"attr":  code,
			"value": value,
		}).
		Limit(1)

	query, args := builder.MustSql()

	row := repo.Pool.QueryRow(ctx, query, args...)

	color := core.AttributeColor{}

	err := row.Scan(
		&color.Code,
		&color.Value,
		&color.TextColor,
		&color.BackgroundColor,
		&color.CreatedAt,
	)
	if err != nil { // TODO: err no rows
		return core.AttributeColor{}, fmt.Errorf("scan row: %w", err)
	}

	return color, nil
}
