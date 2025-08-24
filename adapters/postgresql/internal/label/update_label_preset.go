package label

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *LabelRepo) UpdateLabelPreset(ctx context.Context, preset core.BookLabelPreset) error {
	builder := squirrel.Update("label_presets").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			"description": model.StringToDB(preset.Description),
			"values":      preset.Values,
			"updated_at":  model.TimeToDB(preset.UpdatedAt),
		}).
		Where(squirrel.Eq{
			"name": preset.Name,
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
