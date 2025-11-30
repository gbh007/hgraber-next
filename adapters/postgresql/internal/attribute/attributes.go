package attribute

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *AttributeRepo) Attributes(ctx context.Context) ([]core.Attribute, error) {
	attrTable := model.AttributeTable

	builder := squirrel.Select(attrTable.Columns()...).
		From(attrTable.Name()).
		PlaceholderFormat(squirrel.Dollar).
		OrderBy(attrTable.ColumnOrder())

	query, args := builder.MustSql()

	rows, err := repo.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	result := make([]core.Attribute, 0, core.PossibleAttributeCount)

	for rows.Next() {
		attribute := core.Attribute{}

		err = rows.Scan(attrTable.Scanner(&attribute))
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		result = append(result, attribute)
	}

	return result, nil
}
