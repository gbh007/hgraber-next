package massload

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

func (repo *MassloadRepo) UpdateMassloadCounts(ctx context.Context, ml massloadmodel.Massload) error {
	builder := squirrel.Update("massloads").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			"books_ahead":    model.NilInt64ToDB(ml.BooksAhead),
			"new_books":      model.NilInt64ToDB(ml.NewBooks),
			"existing_books": model.NilInt64ToDB(ml.ExistingBooks),
			"updated_at":     model.TimeToDB(ml.UpdatedAt),
		}).
		Where(squirrel.Eq{
			"id": ml.ID,
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
