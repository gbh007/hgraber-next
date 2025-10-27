package file

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *FileRepo) FSFilesInfo(
	ctx context.Context,
	fsID uuid.UUID,
	onlyInvalidData, onlyDetached bool,
) (core.SizeWithCount, error) {
	fileTable := model.FileTable
	pageTable := model.PageTable

	builder := squirrel.Select(
		"COUNT(*)",
		"SUM("+fileTable.ColumnSize()+")",
	).
		PlaceholderFormat(squirrel.Dollar).
		From(fileTable.Name()).
		Where(squirrel.Eq{
			fileTable.ColumnFSID(): fsID,
		})

	if onlyInvalidData {
		builder = builder.Where(squirrel.Eq{
			fileTable.ColumnInvalidData(): true,
		})
	}

	if onlyDetached {
		builder = builder.Where(
			squirrel.Expr(
				`NOT EXISTS (SELECT 1 FROM ` +
					pageTable.Name() +
					" WHERE " +
					pageTable.Name() + "." + pageTable.ColumnFileID() +
					" = " +
					fileTable.Name() + "." + fileTable.ColumnID() + ")",
			),
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
