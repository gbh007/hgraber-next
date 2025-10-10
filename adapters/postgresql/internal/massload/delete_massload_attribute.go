package massload

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

func (repo *MassloadRepo) DeleteMassloadAttribute(ctx context.Context, id int, attr massloadmodel.Attribute) error {
	table := model.MassloadAttributeTable

	builder := squirrel.Delete(table.Name()).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			table.ColumnMassloadID(): id,
			table.ColumnAttrCode():   attr.Code,
			table.ColumnAttrValue():  attr.Value,
		})

	query, args := builder.MustSql()

	_, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
