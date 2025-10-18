package massload

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

func (repo *MassloadRepo) CreateMassload(ctx context.Context, ml massloadmodel.Massload) (int, error) {
	table := model.MassloadTable

	builder := squirrel.Insert(table.Name()).
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			table.ColumnName():        ml.Name,
			table.ColumnDescription(): model.StringToDB(ml.Description),
			table.ColumnFlags():       ml.Flags,
			table.ColumnCreatedAt():   ml.CreatedAt,
		}).
		Suffix("RETURNING " + table.ColumnID())

	query, args := builder.MustSql()

	var id int

	row := repo.Pool.QueryRow(ctx, query, args...)

	err := row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("exec query: %w", err)
	}

	return id, nil
}
