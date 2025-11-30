package attribute

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
)

func (repo *AttributeRepo) DeleteAttributeRemap(ctx context.Context, code, value string) error {
	attrRemapTable := model.AttributeRemapTable

	builder := squirrel.Delete(attrRemapTable.Name()).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			attrRemapTable.ColumnAttr():  code,
			attrRemapTable.ColumnValue(): value,
		})

	query, args := builder.MustSql()

	_, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
