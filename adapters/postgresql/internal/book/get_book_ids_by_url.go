package book

import (
	"context"
	"fmt"
	"net/url"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
)

func (repo *BookRepo) GetBookIDsByURL(ctx context.Context, u url.URL) ([]uuid.UUID, error) {
	bookTable := model.BookTable

	builder := squirrel.Select(bookTable.ColumnID()).
		PlaceholderFormat(squirrel.Dollar).
		From(bookTable.Name()).
		Where(squirrel.Eq{
			bookTable.ColumnOriginURL(): u.String(),
		})

	query, args := builder.MustSql()

	result := []uuid.UUID{}

	rows, err := repo.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		id := uuid.UUID{}

		err := rows.Scan(&id)
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		result = append(result, id)
	}

	return result, nil
}
