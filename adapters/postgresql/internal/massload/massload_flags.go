package massload

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

func (repo *MassloadRepo) MassloadFlags(ctx context.Context) ([]massloadmodel.Flag, error) {
	table := model.MassloadFlagTable

	builder := squirrel.Select(table.Columns()...).
		PlaceholderFormat(squirrel.Dollar).
		From(table.Name()).
		OrderBy(
			table.ColumnOrderWeight()+" DESC",
			table.ColumnCreatedAt(),
			table.ColumnCode(),
		)

	query, args := builder.MustSql()

	result := make([]massloadmodel.Flag, 0)

	rows, err := repo.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		flag := massloadmodel.Flag{}

		err := rows.Scan(table.Scanner(&flag))
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		result = append(result, flag)
	}

	return result, nil
}
