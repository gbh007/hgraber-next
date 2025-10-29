package book

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *BookRepo) DeleteBook(ctx context.Context, id uuid.UUID) error {
	bookTable := model.BookTable

	builder := squirrel.Delete(bookTable.Name()).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			bookTable.ColumnID(): id,
		})

	query, args := builder.MustSql()

	res, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	if res.RowsAffected() < 1 {
		return core.ErrBookNotFound
	}

	return nil
}
