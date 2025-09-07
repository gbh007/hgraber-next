package massload

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

func (repo *MassloadRepo) DeleteMassloadAttribute(ctx context.Context, id int, attr massloadmodel.Attribute) error {
	builder := squirrel.Delete("massload_attributes").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			"massload_id": id,
			"attr_code":   attr.Code,
			"attr_value":  attr.Value,
		})

	query, args := builder.MustSql()

	_, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
