package file

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *FileRepo) UpdateFileInvalidData(ctx context.Context, fileID uuid.UUID, invalidData bool) error {
	builder := squirrel.Update("files").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			"invalid_data": invalidData,
		}).
		Where(squirrel.Eq{
			"id": fileID,
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
