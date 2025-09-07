package book

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (repo *BookRepo) BookIDsWithDeletedRebuilds(ctx context.Context) ([]uuid.UUID, error) {
	builder := squirrel.Select("id").
		PlaceholderFormat(squirrel.Dollar).
		From("books").
		Where(squirrel.Eq{
			"deleted":    true,
			"is_rebuild": true,
		})

	query, args := builder.MustSql()

	result := []uuid.UUID{}

	rows, err := repo.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		id := uuid.UUID{}

		err := rows.Scan(&id)
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		result = append(result, id)
	}

	return result, nil
}
