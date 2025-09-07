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
	builder := squirrel.Select(model.PageColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("pages").
		Where(squirrel.Eq{
			"book_id":     id,
			"page_number": pageNumber,
		}).
		Limit(1)

	query, args := builder.MustSql()

	page := core.Page{}

	row := repo.Pool.QueryRow(ctx, query, args...)

	err := row.Scan(model.PageScanner(&page))
	if err != nil {
		return core.Page{}, fmt.Errorf("exec query: %w", err)
	}

	return page, nil
}
