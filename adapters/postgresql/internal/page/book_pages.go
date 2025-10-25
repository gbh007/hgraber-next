package page

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *PageRepo) BookPages(ctx context.Context, bookID uuid.UUID) ([]core.Page, error) {
	table := model.PageTable

	builder := squirrel.Select(table.Columns()...).
		PlaceholderFormat(squirrel.Dollar).
		From(table.Name()).
		Where(squirrel.Eq{
			table.ColumnBookID(): bookID,
		}).
		OrderBy(table.ColumnPageNumber())

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
