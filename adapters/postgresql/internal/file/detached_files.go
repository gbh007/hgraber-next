package file

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *FileRepo) DetachedFiles(ctx context.Context) ([]core.File, error) {
	fileTable := model.FileTable.WithPrefix(model.FileTable.Name())
	pageTable := model.PageTable.WithPrefix(model.PageTable.Name())

	builder := squirrel.Select(fileTable.Columns()...).
		PlaceholderFormat(squirrel.Dollar).
		From(fileTable.NameAlter()).
		Where(squirrel.Expr(
			`NOT EXISTS (SELECT 1 FROM ` +
				pageTable.NameAlter() +
				" WHERE " +
				pageTable.ColumnFileID() +
				" = " +
				fileTable.ColumnID() +
				")",
		))

	query, args := builder.MustSql()

	result := make([]core.File, 0)

	rows, err := repo.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		file := core.File{}

		err := rows.Scan(fileTable.Scanner(&file))
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		result = append(result, file)
	}

	return result, nil
}
