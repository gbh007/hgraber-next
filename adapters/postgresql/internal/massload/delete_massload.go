package massload

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
)

func (repo *MassloadRepo) DeleteMassload(ctx context.Context, id int) error {
	table := model.MassloadTable

	builder := squirrel.Delete(table.Name()).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			table.ColumnID(): id,
		})

	query, args := builder.MustSql()

	_, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
