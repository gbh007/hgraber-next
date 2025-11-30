package attribute

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *AttributeRepo) AttributeColor(ctx context.Context, code, value string) (core.AttributeColor, error) {
	attrColorTable := model.AttributeColorTable

	builder := squirrel.Select(attrColorTable.Columns()...).
		PlaceholderFormat(squirrel.Dollar).
		From(attrColorTable.Name()).
		Where(squirrel.Eq{
			attrColorTable.ColumnAttr():  code,
			attrColorTable.ColumnValue(): value,
		}).
		Limit(1)

	query, args := builder.MustSql()

	row := repo.Pool.QueryRow(ctx, query, args...)

	color := core.AttributeColor{}

	err := row.Scan(attrColorTable.Scanner(&color))
	if err != nil { // TODO: err no rows
		return core.AttributeColor{}, fmt.Errorf("scan row: %w", err)
	}

	return color, nil
}
