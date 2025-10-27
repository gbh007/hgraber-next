package file

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

// TODO: добавить лимиты
func (repo *FileRepo) GetUnHashedFiles(ctx context.Context) ([]core.File, error) {
	table := model.FileTable

	builder := squirrel.Select(table.Columns()...).
		PlaceholderFormat(squirrel.Dollar).
		From(table.Name()).
		Where(squirrel.Or{
			squirrel.Expr(table.ColumnMd5Sum() + " IS NULL"),
			squirrel.Expr(table.ColumnSha256Sum() + " IS NULL"),
			squirrel.Expr(table.ColumnSize() + " IS NULL"),
		})

	query, args := builder.MustSql()

	result := make([]core.File, 0)

	rows, err := repo.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

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
