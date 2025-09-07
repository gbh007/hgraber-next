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
			"page_size":       model.NilInt64ToDB(attr.PageSize),
			"file_size":       model.NilInt64ToDB(attr.FileSize),
			"page_count":      model.NilInt64ToDB(attr.PageCount),
			"file_count":      model.NilInt64ToDB(attr.FileCount),
			"books_in_system": model.NilInt64ToDB(attr.BookInSystem),
			"updated_at":      model.TimeToDB(attr.UpdatedAt),
		}).
		Where(squirrel.Eq{
			"attr_code":  attr.Code,
			"attr_value": attr.Value,
		})

	query, args := builder.MustSql()

	_, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
