package attribute

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *AttributeRepo) UpdateAttributeColor(ctx context.Context, color core.AttributeColor) error {
	attrColorTable := model.AttributeColorTable

	builder := squirrel.Update(attrColorTable.Name()).
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			attrColorTable.ColumnTextColor():       color.TextColor,
			attrColorTable.ColumnBackgroundColor(): color.BackgroundColor,
		}).
		Where(squirrel.Eq{
			attrColorTable.ColumnAttr():  color.Code,
			attrColorTable.ColumnValue(): color.Value,
		})

	query, args := builder.MustSql()

	_, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
