package massload

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

func (repo *MassloadRepo) Massload(ctx context.Context, id int) (massloadmodel.Massload, error) {
	table := model.MassloadTable

	builder := squirrel.Select(table.Columns()...).
		PlaceholderFormat(squirrel.Dollar).
		From(table.Name()).
		Where(squirrel.Eq{
			table.ColumnID(): id,
		}).
		Limit(1)

	query, args := builder.MustSql()

	ml := massloadmodel.Massload{}

	row := repo.Pool.QueryRow(ctx, query, args...)

	err := row.Scan(table.Scanner(&ml))
	if err != nil {
		return massloadmodel.Massload{}, fmt.Errorf("exec: %w", err)
	}

	return ml, nil
}
