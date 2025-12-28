package file

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *FileRepo) DuplicatedFiles(ctx context.Context) ([]core.File, error) {
	table := model.FileTable.WithPrefix("f")
	subTable := model.FileTable

	query, args := squirrel.Select(table.Columns()...).
		PlaceholderFormat(squirrel.Dollar).
		FromSelect(
			squirrel.
				Select(
					"COUNT(*) AS c",
					subTable.ColumnMd5Sum(),
					subTable.ColumnSha256Sum(),
				).
				From(subTable.Name()).
				GroupBy(
					subTable.ColumnMd5Sum(),
					subTable.ColumnSha256Sum(),
				).Having(squirrel.Gt{
				"COUNT(*)": 1,
			}),
			"t",
		).
		InnerJoin(
			table.NameAlter() +
				" ON " + table.ColumnMd5Sum() + " = t." + subTable.ColumnMd5Sum() +
				" AND " + table.ColumnSha256Sum() + " = t." + subTable.ColumnSha256Sum(),
		).
		OrderBy(table.ColumnID()).
		MustSql()

	rows, err := repo.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec: %w", err)
	}

	result := make([]core.File, 0)

	defer rows.Close()

	for rows.Next() {
		file := core.File{}

		err := rows.Scan(table.Scanner(&file))
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		result = append(result, file)
	}

	return result, nil
}
