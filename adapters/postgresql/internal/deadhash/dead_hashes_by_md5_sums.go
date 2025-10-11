package deadhash

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *DeadHashRepo) DeadHashesByMD5Sums(ctx context.Context, md5Sums []string) ([]core.DeadHash, error) {
	table := model.DeadHashTable

	builder := squirrel.Select(table.Columns()...).
		PlaceholderFormat(squirrel.Dollar).
		From(table.Name()).
		Where(squirrel.Eq{
			table.ColumnMd5Sum(): md5Sums,
		})

	query, args := builder.MustSql()

	rows, err := repo.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	result := make([]core.DeadHash, 0, len(md5Sums))

	for rows.Next() {
		hash := core.DeadHash{}

		err = rows.Scan(table.Scanner(&hash))
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		result = append(result, hash)
	}

	return result, nil
}
