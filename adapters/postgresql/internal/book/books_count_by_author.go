package book

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
)

func (repo *BookRepo) BooksCountByAuthor(ctx context.Context) (map[string]int64, error) {
	bookAttributeTable := model.BookAttributeTable

	builder := squirrel.Select("COUNT(*)", bookAttributeTable.ColumnValue()).
		PlaceholderFormat(squirrel.Dollar).
		From(bookAttributeTable.Name()).
		Where(squirrel.Eq{
			bookAttributeTable.ColumnAttr(): "author",
		}).
		GroupBy(bookAttributeTable.ColumnValue())

	query, args := builder.MustSql()

	out := make(map[string]int64, 100) //nolint:mnd // оптимизация

	rows, err := repo.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var (
			count sql.NullInt64
			name  string
		)

		err = rows.Scan(&count, &name)
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		out[name] = count.Int64
	}

	return out, nil
}
