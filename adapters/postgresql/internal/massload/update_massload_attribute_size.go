package massload

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

func (repo *MassloadRepo) UpdateMassloadAttributeSize(ctx context.Context, attr massloadmodel.Attribute) error {
	table := model.MassloadAttributeTable

	builder := squirrel.Update(table.Name()).
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			table.ColumnPageSize():      model.NilInt64ToDB(attr.PageSize),
			table.ColumnFileSize():      model.NilInt64ToDB(attr.FileSize),
			table.ColumnPageCount():     model.NilInt64ToDB(attr.PageCount),
			table.ColumnFileCount():     model.NilInt64ToDB(attr.FileCount),
			table.ColumnBooksInSystem(): model.NilInt64ToDB(attr.BookInSystem),
			table.ColumnUpdatedAt():     model.TimeToDB(attr.UpdatedAt),
		}).
		Where(squirrel.Eq{
			table.ColumnAttrCode():  attr.Code,
			table.ColumnAttrValue(): attr.Value,
		})

	query, args := builder.MustSql()

	_, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
