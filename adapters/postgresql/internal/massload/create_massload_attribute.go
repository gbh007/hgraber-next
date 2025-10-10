package massload

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

func (repo *MassloadRepo) CreateMassloadAttribute(ctx context.Context, id int, attr massloadmodel.Attribute) error {
	table := model.MassloadAttributeTable

	builder := squirrel.Insert(table.Name()).
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			table.ColumnMassloadID(): id,
			table.ColumnAttrCode():   attr.Code,
			table.ColumnAttrValue():  attr.Value,
			table.ColumnCreatedAt():  attr.CreatedAt,
		})

	query, args := builder.MustSql()

	_, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
