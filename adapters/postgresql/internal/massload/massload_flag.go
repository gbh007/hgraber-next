package massload

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

func (repo *MassloadRepo) MassloadFlag(ctx context.Context, code string) (massloadmodel.Flag, error) {
	builder := squirrel.Select(model.MassloadFlagColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("massload_flags").
		Where(squirrel.Eq{
			"code": code,
		}).
		Limit(1)

	query, args := builder.MustSql()

	flag := massloadmodel.Flag{}

	row := repo.Pool.QueryRow(ctx, query, args...)

	err := row.Scan(model.MassloadFlagScanner(&flag))
	if err != nil {
		return massloadmodel.Flag{}, fmt.Errorf("exec: %w", err)
	}

	return flag, nil
}
