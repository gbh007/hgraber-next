package attribute

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *AttributeRepo) InsertAttributeColor(ctx context.Context, color core.AttributeColor) error {
	builder := squirrel.Insert("attribute_colors").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			"attr":             color.Code,
			"value":            color.Value,
			"text_color":       color.TextColor,
			"background_color": color.BackgroundColor,
			"created_at":       color.CreatedAt,
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
