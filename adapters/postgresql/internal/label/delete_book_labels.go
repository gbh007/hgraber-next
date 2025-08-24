package label

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (repo *LabelRepo) DeleteBookLabels(ctx context.Context, bookID uuid.UUID) error {
	builder := squirrel.Delete("book_labels").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			"book_id": bookID,
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
