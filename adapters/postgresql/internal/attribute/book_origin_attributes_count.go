package attribute

import (
	"context"
	"fmt"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *AttributeRepo) BookOriginAttributesCount(ctx context.Context) ([]core.AttributeVariant, error) {
	rows, err := repo.Pool.Query(
		ctx,
		`SELECT COUNT(*), attr, UNNEST(values) as value FROM book_origin_attributes GROUP BY attr, value;`,
	)
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
