package massload

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

func (repo *MassloadRepo) Massload(ctx context.Context, id int) (massloadmodel.Massload, error) {
	builder := squirrel.Select(model.MassloadColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("massloads").
		Where(squirrel.Eq{
			"id": id,
		}).
		Limit(1)

	query, args := builder.MustSql()

	ml := massloadmodel.Massload{}

	row := repo.Pool.QueryRow(ctx, query, args...)

	err := row.Scan(model.MassloadScanner(&ml))
	if err != nil {
		return massloadmodel.Massload{}, fmt.Errorf("exec: %w", err)
	}

	return ml, nil
}
