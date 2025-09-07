package file

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *FileRepo) FSFilesInfo(
	ctx context.Context,
	fsID uuid.UUID,
	onlyInvalidData, onlyDetached bool,
) (core.SizeWithCount, error) {
	builder := squirrel.Select(
		"COUNT(*)",
		"SUM(\"size\")",
	).
		PlaceholderFormat(squirrel.Dollar).
		From("files").
		Where(squirrel.Eq{
			"fs_id": fsID,
		})

	if onlyInvalidData {
		builder = builder.Where(squirrel.Eq{
			"invalid_data": true,
		})
	}

	if onlyDetached {
		builder = builder.Where(
			squirrel.Expr(`NOT EXISTS (SELECT 1 FROM pages WHERE file_id = files.id)`),
		)
	}

	query, args := builder.MustSql()

	var count, size sql.NullInt64

	row := repo.Pool.QueryRow(ctx, query, args...)

	err := row.Scan(&count, &size)
	if err != nil {
		return core.SizeWithCount{}, fmt.Errorf("scan: %w", err)
	}

	return core.SizeWithCount{
		Count: count.Int64,
		Size:  size.Int64,
	}, nil
}
