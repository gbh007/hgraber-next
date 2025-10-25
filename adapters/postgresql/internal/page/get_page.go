package page

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *PageRepo) GetPage(ctx context.Context, id uuid.UUID, pageNumber int) (core.Page, error) {
	table := model.PageTable

	builder := squirrel.Select(table.Columns()...).
		PlaceholderFormat(squirrel.Dollar).
		From(table.Name()).
		Where(squirrel.Eq{
			table.ColumnBookID():     id,
			table.ColumnPageNumber(): pageNumber,
		}).
		Limit(1)

	query, args := builder.MustSql()

	page := core.Page{}

	row := repo.Pool.QueryRow(ctx, query, args...)

	err := row.Scan(table.Scanner(&page))
	if err != nil {
		return core.Page{}, fmt.Errorf("exec query: %w", err)
	}

	return page, nil
}
