package page

import (
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *PageRepo) UpdatePageDownloaded(
	ctx context.Context,
	bookID uuid.UUID,
	pageNumber int,
	downloaded bool,
	fileID uuid.UUID,
) error {
	table := model.PageTable

	query, args := squirrel.Update(table.Name()).
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			table.ColumnDownloaded(): downloaded,
			table.ColumnLoadAt():     time.Now().UTC(),
			table.ColumnFileID():     model.UUIDToDB(fileID),
		}).
		Where(squirrel.Eq{
			table.ColumnBookID():     bookID,
			table.ColumnPageNumber(): pageNumber,
		}).
		MustSql()

	res, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	if res.RowsAffected() < 1 {
		return core.ErrPageNotFound
	}

	return nil
}
