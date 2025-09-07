package massload

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

func (repo *MassloadRepo) MassloadAttributes(ctx context.Context, id int) ([]massloadmodel.Attribute, error) {
	builder := squirrel.Select(model.MassloadAttributeColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("massload_attributes").
		Where(squirrel.Eq{
			"massload_id": id,
		}).
		OrderBy("created_at")

	query, args := builder.MustSql()

	result := make([]massloadmodel.Attribute, 0)

	rows, err := repo.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		attr := massloadmodel.Attribute{}

		err := rows.Scan(model.MassloadAttributeScanner(&attr))
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		result = append(result, attr)
	}

	return result, nil
}
