package attribute

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *AttributeRepo) AttributeColors(ctx context.Context) ([]core.AttributeColor, error) {
	attrColorTable := model.AttributeColorTable

	builder := squirrel.Select(attrColorTable.Columns()...).
		From(attrColorTable.Name()).
		PlaceholderFormat(squirrel.Dollar)

	query, args := builder.MustSql()

	rows, err := repo.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	result := make([]core.AttributeColor, 0, 10) //nolint:mnd // оптимизация

	for rows.Next() {
		color := core.AttributeColor{}

		err = rows.Scan(attrColorTable.Scanner(&color))
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		result = append(result, color)
	}

	return result, nil
}
