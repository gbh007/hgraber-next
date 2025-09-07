package massload

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

func (repo *MassloadRepo) CreateMassload(ctx context.Context, ml massloadmodel.Massload) (int, error) {
	builder := squirrel.Insert("massloads").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			"name":        ml.Name,
			"description": model.StringToDB(ml.Description),
			"flags":       ml.Flags,
			"created_at":  ml.CreatedAt,
		}).
		Suffix("RETURNING id")

	query, args := builder.MustSql()

	var id int

	row := repo.Pool.QueryRow(ctx, query, args...)

	err := row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("exec query: %w", err)
	}

	return id, nil
}
