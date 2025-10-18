package massload

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

func (repo *MassloadRepo) UpdateMassloadSize(ctx context.Context, ml massloadmodel.Massload) error {
	table := model.MassloadTable

	builder := squirrel.Update(table.Name()).
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			table.ColumnPageSize():      model.NilInt64ToDB(ml.PageSize),
			table.ColumnFileSize():      model.NilInt64ToDB(ml.FileSize),
			table.ColumnPageCount():     model.NilInt64ToDB(ml.PageCount),
			table.ColumnFileCount():     model.NilInt64ToDB(ml.FileCount),
			table.ColumnBooksInSystem(): model.NilInt64ToDB(ml.BookInSystem),
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
