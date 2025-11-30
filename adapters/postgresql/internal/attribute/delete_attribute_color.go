package attribute

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
)

func (repo *AttributeRepo) DeleteAttributeColor(ctx context.Context, code, value string) error {
	attrColorTable := model.AttributeColorTable

	builder := squirrel.Delete(attrColorTable.Name()).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			attrColorTable.ColumnAttr():  code,
			attrColorTable.ColumnValue(): value,
		})

	query, args := builder.MustSql()

	_, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
