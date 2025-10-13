package label

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *LabelRepo) SetLabel(ctx context.Context, label core.BookLabel) error {
	table := model.BookLabelTable

	builder := squirrel.Insert(table.Name()).
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			table.ColumnBookID():     label.BookID,
			table.ColumnPageNumber(): label.PageNumber,
			table.ColumnName():       label.Name,
			table.ColumnValue():      label.Value,
			table.ColumnCreateAt():   label.CreateAt,
		}).
		Suffix(fmt.Sprintf(
			`ON CONFLICT (%s, %s, %s) DO UPDATE SET %s = EXCLUDED.%s`,
			table.ColumnBookID(),
			table.ColumnPageNumber(),
			table.ColumnName(),
			table.ColumnValue(),
			table.ColumnValue(),
		))

	query, args := builder.MustSql()

	_, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
