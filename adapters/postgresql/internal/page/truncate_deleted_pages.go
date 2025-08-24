package page

import (
	"context"
	"fmt"
)

func (repo *PageRepo) TruncateDeletedPages(ctx context.Context) error {
	_, err := repo.Pool.Exec(ctx, `TRUNCATE deleted_pages;`)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
