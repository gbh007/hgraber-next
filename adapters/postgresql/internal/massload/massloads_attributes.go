package massload

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

func (repo *MassloadRepo) MassloadsAttributes(ctx context.Context) ([]massloadmodel.Attribute, error) {
	builder := squirrel.
		Select(
			"attr_code",
			"attr_value",
		).
		PlaceholderFormat(squirrel.Dollar).
		From("massload_attributes").
		GroupBy(
			"attr_code",
			"attr_value",
		)

	query, args := builder.MustSql()

	result := make([]massloadmodel.Attribute, 0)

	rows, err := repo.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		attr := massloadmodel.Attribute{}

		err := rows.Scan(&attr.Code, &attr.Value)
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		result = append(result, attr)
	}

	return result, nil
}
