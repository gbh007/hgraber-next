package page

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
)

func (repo *PageRepo) RemoveDeletedPages(ctx context.Context, bookID uuid.UUID, pageNumbers []int) error {
	deletedPageTable := model.DeletedPageTable

	builder := squirrel.Delete(deletedPageTable.Name()).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			deletedPageTable.ColumnBookID():     bookID,
			deletedPageTable.ColumnPageNumber(): pageNumbers,
		})

	query, args := builder.MustSql()

	_, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
