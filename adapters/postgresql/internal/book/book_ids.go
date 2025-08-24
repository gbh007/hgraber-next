package book

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *BookRepo) BookIDs(ctx context.Context, filter core.BookFilter) ([]uuid.UUID, error) {
	ids := make([]uuid.UUID, 0)

	query, args, err := repo.buildBooksFilter(ctx, filter, false)
	if err != nil {
		return nil, fmt.Errorf("build book filter: %w", err)
	}

	rows, err := repo.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var id uuid.UUID

		err := rows.Scan(&id)
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		ids = append(ids, id)
	}

	return ids, nil
}
