package page

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
)

func (repo *PageRepo) BookPagesCount(ctx context.Context, bookID uuid.UUID) (int, error) {
	table := model.PageTable

	builder := squirrel.Select("COUNT(*)").
		PlaceholderFormat(squirrel.Dollar).
		From(table.Name()).
		Where(squirrel.Eq{
			table.ColumnBookID(): bookID,
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
