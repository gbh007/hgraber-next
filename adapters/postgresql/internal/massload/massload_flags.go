package massload

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

func (repo *MassloadRepo) MassloadFlags(ctx context.Context) ([]massloadmodel.Flag, error) {
	builder := squirrel.Select(model.MassloadFlagColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("massload_flags").
		OrderBy("order_weight DESC", "created_at", "code")

	query, args := builder.MustSql()

	result := make([]massloadmodel.Flag, 0)

	rows, err := repo.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		flag := massloadmodel.Flag{}

		err := rows.Scan(model.MassloadFlagScanner(&flag))
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		result = append(result, flag)
	}

	return result, nil
}
