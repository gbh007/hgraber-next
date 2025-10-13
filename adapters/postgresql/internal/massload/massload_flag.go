package massload

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

func (repo *MassloadRepo) MassloadFlag(ctx context.Context, code string) (massloadmodel.Flag, error) {
	table := model.MassloadFlagTable

	builder := squirrel.Select(table.Columns()...).
		PlaceholderFormat(squirrel.Dollar).
		From(table.Name()).
		Where(squirrel.Eq{
			table.ColumnCode(): code,
		}).
		Limit(1)

	query, args := builder.MustSql()

	flag := massloadmodel.Flag{}

	row := repo.Pool.QueryRow(ctx, query, args...)

	err := row.Scan(table.Scanner(&flag))
	if err != nil {
		return massloadmodel.Flag{}, fmt.Errorf("exec: %w", err)
	}

	return flag, nil
}
