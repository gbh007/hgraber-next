package page

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *PageRepo) BookPagesWithHash(ctx context.Context, bookID uuid.UUID) ([]core.PageWithHash, error) {
	builder := squirrel.Select(model.PageWithHashColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("pages p").
		LeftJoin("files f ON p.file_id = f.id").
		Where(squirrel.Eq{
			"p.book_id": bookID,
		}).
		OrderBy("p.page_number")

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
