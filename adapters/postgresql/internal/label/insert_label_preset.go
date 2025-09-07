package label

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *LabelRepo) InsertLabelPreset(ctx context.Context, preset core.BookLabelPreset) error {
	builder := squirrel.Insert("label_presets").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			"name":        preset.Name,
			"description": model.StringToDB(preset.Description),
			"values":      preset.Values,
			"created_at":  preset.CreatedAt,
		})

	query, args := builder.MustSql()

	_, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
