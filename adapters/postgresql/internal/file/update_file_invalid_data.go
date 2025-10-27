package file

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *FileRepo) UpdateFileInvalidData(ctx context.Context, fileID uuid.UUID, invalidData bool) error {
	table := model.FileTable

	builder := squirrel.Update(table.Name()).
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			table.ColumnInvalidData(): invalidData,
		}).
		Where(squirrel.Eq{
			table.ColumnID(): fileID,
		})

	query, args := builder.MustSql()

	res, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	if res.RowsAffected() < 1 {
		return core.ErrFileNotFound
	}

	return nil
}
