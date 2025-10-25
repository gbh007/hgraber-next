package page

import (
	"context"
	"fmt"
	"net/url"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *PageRepo) PagesByURL(ctx context.Context, u url.URL) ([]core.Page, error) {
	table := model.PageTable

	builder := squirrel.Select(table.Columns()...).
		PlaceholderFormat(squirrel.Dollar).
		From(table.Name()).
		Where(squirrel.Eq{
			table.ColumnOriginURL(): u.String(),
		}).
		OrderBy(table.ColumnBookID(), table.ColumnPageNumber())

	query, args := builder.MustSql()

	result := make([]core.Page, 0)

	rows, err := repo.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		page := core.Page{}

		err := rows.Scan(table.Scanner(&page))
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		result = append(result, page)
	}

	return result, nil
}
