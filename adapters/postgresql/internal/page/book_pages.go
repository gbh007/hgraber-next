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
	builder := squirrel.Select(model.PageColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("pages").
		Where(squirrel.Eq{
			"book_id": bookID,
		}).
		OrderBy("page_number")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	repo.SquirrelDebugLog(ctx, query, args)

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
