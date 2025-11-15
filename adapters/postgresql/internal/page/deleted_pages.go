package page

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *PageRepo) DeletedPages(ctx context.Context, bookID uuid.UUID) ([]core.PageWithHash, error) {
	deletedPageTable := model.DeletedPageTable

	builder := squirrel.Select(deletedPageTable.ToPageWithHashColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From(deletedPageTable.Name()).
		Where(squirrel.Eq{
			deletedPageTable.ColumnBookID(): bookID,
		}).
		OrderBy(deletedPageTable.ColumnPageNumber())

	query, args := builder.MustSql()

	out := make([]core.PageWithHash, 0, 10) //nolint:mnd // оптимизация

	rows, err := repo.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		page := core.PageWithHash{}

		err := rows.Scan(deletedPageTable.ToPageWithHashScanner(&page))
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		out = append(out, page)
	}

	return out, nil
}
