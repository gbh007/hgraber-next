package massload

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

func (repo *MassloadRepo) UpdateMassloadCounts(ctx context.Context, ml massloadmodel.Massload) error {
	table := model.MassloadTable

	builder := squirrel.Update(table.Name()).
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			table.ColumnBooksAhead():    model.NilInt64ToDB(ml.BooksAhead),
			table.ColumnNewBooks():      model.NilInt64ToDB(ml.NewBooks),
			table.ColumnExistingBooks(): model.NilInt64ToDB(ml.ExistingBooks),
			table.ColumnUpdatedAt():     model.TimeToDB(ml.UpdatedAt),
		}).
		Where(squirrel.Eq{
			table.ColumnID(): ml.ID,
		})

	query, args := builder.MustSql()

	_, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
