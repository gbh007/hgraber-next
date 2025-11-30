package attribute

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *AttributeRepo) InsertAttributeColor(ctx context.Context, color core.AttributeColor) error {
	attrColorTable := model.AttributeColorTable

	builder := squirrel.Insert(attrColorTable.Name()).
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			attrColorTable.ColumnAttr():            color.Code,
			attrColorTable.ColumnValue():           color.Value,
			attrColorTable.ColumnTextColor():       color.TextColor,
			attrColorTable.ColumnBackgroundColor(): color.BackgroundColor,
			attrColorTable.ColumnCreatedAt():       color.CreatedAt,
		})

	query, args := builder.MustSql()

	_, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
