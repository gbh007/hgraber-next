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
	builder := squirrel.Select(model.PageColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("pages").
		Where(squirrel.Eq{
			"origin_url": u.String(),
		}).
		OrderBy("book_id", "page_number")

	query, args := builder.MustSql()

	result := make([]core.Page, 0)

	rows, err := repo.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		page := core.Page{}

		err := rows.Scan(model.PageScanner(&page))
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		result = append(result, page)
	}

	return result, nil
}
