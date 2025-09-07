package book

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *BookRepo) DeleteBooks(ctx context.Context, ids []uuid.UUID) error {
	builder := squirrel.Delete("books").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			"id": ids,
		})

	query, args := builder.MustSql()

	res, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	if res.RowsAffected() < 1 {
		return core.BookNotFoundError
	}

	return nil
}
