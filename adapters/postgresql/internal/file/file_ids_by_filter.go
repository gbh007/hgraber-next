package file

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/fsmodel"
)

func (repo *FileRepo) FileIDsByFilter(ctx context.Context, filter fsmodel.FileFilter) ([]uuid.UUID, error) {
	fileTable := model.FileTable
	pageTable := model.PageTable

	builder := squirrel.Select(fileTable.ColumnID()).
		PlaceholderFormat(squirrel.Dollar).
		From(fileTable.Name())

	if filter.FSID != nil {
		builder = builder.Where(squirrel.Eq{
			fileTable.ColumnFSID(): *filter.FSID,
		})
	}

	if filter.BookID != nil || filter.PageNumber != nil {
		subBuilder := squirrel.Select("1").
			// Важно: либа не может переконвертить другой тип форматирования для подзапроса!
			PlaceholderFormat(squirrel.Question).
			From(pageTable.Name()).
			Where(squirrel.Expr(pageTable.ColumnFileID() + " = " + fileTable.Name() + "." + fileTable.ColumnID()))

		if filter.BookID != nil {
			subBuilder = subBuilder.Where(squirrel.Eq{
				pageTable.ColumnBookID(): *filter.BookID,
			})
		}

		if filter.PageNumber != nil {
			subBuilder = subBuilder.Where(squirrel.Eq{
				pageTable.ColumnPageNumber(): *filter.PageNumber,
			})
		}

		subQuery, subArgs, err := subBuilder.ToSql()
		if err != nil {
			return nil, fmt.Errorf("build pages sub query: %w", err)
		}

		builder = builder.Where(squirrel.Expr(`EXISTS (`+subQuery+`)`, subArgs...))
	}

	query, args := builder.MustSql()

	rows, err := repo.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec: %w", err)
	}

	defer rows.Close()

	ids := make([]uuid.UUID, 0, 10) //nolint:mnd // предварительная оптимизация

	for rows.Next() {
		var id uuid.UUID

		err = rows.Scan(&id)
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		ids = append(ids, id)
	}

	return ids, nil
}
