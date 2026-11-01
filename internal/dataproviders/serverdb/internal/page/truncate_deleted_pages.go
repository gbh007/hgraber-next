package page

import (
	"context"
	"fmt"

	"github.com/gbh007/hgraber-next/internal/dataproviders/serverdb/internal/model"
)

func (repo *PageRepo) TruncateDeletedPages(ctx context.Context) error {
	deletedPageTable := model.DeletedPageTable

	_, err := repo.Pool.Exec(ctx, "TRUNCATE "+deletedPageTable.Name())
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
