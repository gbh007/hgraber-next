package massload

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

func (repo *MassloadRepo) UpdateMassloadSize(ctx context.Context, ml massloadmodel.Massload) error {
	builder := squirrel.Update("massloads").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			"page_size":       model.NilInt64ToDB(ml.PageSize),
			"file_size":       model.NilInt64ToDB(ml.FileSize),
			"page_count":      model.NilInt64ToDB(ml.PageCount),
			"file_count":      model.NilInt64ToDB(ml.FileCount),
			"books_in_system": model.NilInt64ToDB(ml.BookInSystem),
			"updated_at":      model.TimeToDB(ml.UpdatedAt),
		}).
		Where(squirrel.Eq{
			"id": ml.ID,
		})

	query, args := builder.MustSql()

	_, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
