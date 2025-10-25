package page

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *PageRepo) BookPagesWithHashByMD5Sums(ctx context.Context, md5Sums []string) ([]core.PageWithHash, error) {
	pageTable := model.PageTable

	builder := squirrel.Select(model.PageWithHashColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From(pageTable.Name() + " p").
		LeftJoin("files f ON p." + pageTable.ColumnFileID() + " = f.id").
		Where(squirrel.Eq{
			"f.md5_sum": md5Sums,
		})

	query, args := builder.MustSql()

	out := make([]core.PageWithHash, 0, 10) //nolint:mnd // оптимизация

	rows, err := repo.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		page := core.PageWithHash{}

		err := rows.Scan(model.PageWithHashScanner(&page))
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		out = append(out, page)
	}

	return out, nil
}
