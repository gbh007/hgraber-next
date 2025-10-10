package label

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
)

func (repo *LabelRepo) DeleteLabelPreset(ctx context.Context, name string) error {
	table := model.BookLabelPresetTable

	builder := squirrel.Delete(table.Name()).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			table.ColumnName(): name,
		})

	query, args := builder.MustSql()

	_, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
