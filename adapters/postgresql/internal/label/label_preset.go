package label

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *LabelRepo) LabelPreset(ctx context.Context, name string) (core.BookLabelPreset, error) {
	table := model.BookLabelPresetTable

	builder := squirrel.Select(table.Columns()...).
		PlaceholderFormat(squirrel.Dollar).
		From(table.Name()).
		Where(squirrel.Eq{
			table.ColumnName(): name,
		}).
		Limit(1)

	query, args := builder.MustSql()

	row := repo.Pool.QueryRow(ctx, query, args...)

	var preset core.BookLabelPreset

	err := row.Scan(table.Scanner(&preset))
	if err != nil { // TODO: err no rows
		return core.BookLabelPreset{}, fmt.Errorf("scan row: %w", err)
	}

	return preset, nil
}
