package page

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (repo *PageRepo) BookPagesCount(ctx context.Context, bookID uuid.UUID) (int, error) {
	builder := squirrel.Select("COUNT(*)").
		PlaceholderFormat(squirrel.Dollar).
		From("pages").
		Where(squirrel.Eq{
			"book_id": bookID,
		})

	query, args := builder.MustSql()

	count := sql.NullInt64{}
	row := repo.Pool.QueryRow(ctx, query, args...)

	err := row.Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("get count: %w", err)
	}

	return int(count.Int64), nil
}
