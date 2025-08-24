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
	builder := squirrel.Select(model.FileColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("files").
		Where(squirrel.Or{
			squirrel.Expr(`md5_sum IS NULL`),
			squirrel.Expr(`sha256_sum IS NULL`),
			squirrel.Expr(`"size" IS NULL`),
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("storage: build query: %w", err)
	}

	repo.SquirrelDebugLog(ctx, query, args)

	result := make([]core.File, 0)

	rows, err := repo.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		file := core.File{}

		err := rows.Scan(model.FileScanner(&file))
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		result = append(result, file)
	}

	return result, nil
}
