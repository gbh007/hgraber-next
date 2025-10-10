package label

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *LabelRepo) LabelPresets(ctx context.Context) ([]core.BookLabelPreset, error) {
	table := model.BookLabelPresetTable

	builder := squirrel.Select(table.Columns()...).
		From(table.Name()).
		PlaceholderFormat(squirrel.Dollar)

	query, args := builder.MustSql()

	rows, err := repo.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	result := make([]core.BookLabelPreset, 0, 10) //nolint:mnd // оптимизация

	for rows.Next() {
		var preset core.BookLabelPreset

		err = rows.Scan(table.Scanner(&preset))
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		result = append(result, preset)
	}

	return result, nil
}
