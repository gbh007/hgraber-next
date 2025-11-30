package attribute

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *AttributeRepo) AttributesCount(ctx context.Context) ([]core.AttributeVariant, error) {
	bookAttributeTable := model.BookAttributeTable

	query, args := squirrel.Select(
		"COUNT(*)",
		bookAttributeTable.ColumnAttr(),
		bookAttributeTable.ColumnValue(),
	).
		PlaceholderFormat(squirrel.Dollar).
		From(bookAttributeTable.Name()).
		GroupBy(
			bookAttributeTable.ColumnAttr(),
			bookAttributeTable.ColumnValue(),
		).
		MustSql()

	rows, err := repo.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("get attributes count: %w", err)
	}

	defer rows.Close()

	result := make([]core.AttributeVariant, 0, 100) //nolint:mnd // оптимизация

	for rows.Next() {
		var (
			count int
			code  string
			value string
		)

		err := rows.Scan(&count, &code, &value)
		if err != nil {
			return nil, fmt.Errorf("get attributes count: scan row: %w", err)
		}

		result = append(result, core.AttributeVariant{
			Code:  code,
			Value: value,
			Count: count,
		})
	}

	return result, nil
}
