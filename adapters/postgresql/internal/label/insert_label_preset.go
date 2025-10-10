package label

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *LabelRepo) InsertLabelPreset(ctx context.Context, preset core.BookLabelPreset) error {
	table := model.BookLabelPresetTable

	builder := squirrel.Insert(table.Name()).
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			table.ColumnName():        preset.Name,
			table.ColumnDescription(): model.StringToDB(preset.Description),
			table.ColumnValues():      preset.Values,
			table.ColumnCreatedAt():   preset.CreatedAt,
		})

	query, args := builder.MustSql()

	_, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
