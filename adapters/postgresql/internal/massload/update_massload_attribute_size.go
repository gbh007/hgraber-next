package massload

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

func (repo *MassloadRepo) UpdateMassloadAttributeSize(ctx context.Context, attr massloadmodel.Attribute) error {
	builder := squirrel.Update("massload_attributes").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			"page_size":  model.NilInt64ToDB(attr.PageSize),
			"file_size":  model.NilInt64ToDB(attr.FileSize),
			"updated_at": model.TimeToDB(attr.UpdatedAt),
		}).
		Where(squirrel.Eq{
			"attr_code":  attr.Code,
			"attr_value": attr.Value,
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	repo.SquirrelDebugLog(ctx, query, args)

	_, err = repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
