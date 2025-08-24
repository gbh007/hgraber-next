package book

import (
	"context"
	"fmt"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *BookRepo) BookCount(ctx context.Context, filter core.BookFilter) (int, error) {
	var c int

	query, args, err := repo.buildBooksFilter(ctx, filter, true)
	if err != nil {
		return 0, fmt.Errorf("build book filter: %w", err)
	}

	row := repo.Pool.QueryRow(ctx, query, args...)

	err = row.Scan(&c)
	if err != nil {
		return 0, fmt.Errorf("exec query: %w", err)
	}

	return c, nil
}
