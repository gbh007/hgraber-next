package page

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (repo *PageRepo) RemoveDeletedPages(ctx context.Context, bookID uuid.UUID, pageNumbers []int) error {
	builder := squirrel.Delete("deleted_pages").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			"book_id":     bookID,
			"page_number": pageNumbers,
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
