package book

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
)

func (repo *BookRepo) BooksCountByAuthor(ctx context.Context) (map[string]int64, error) {
	builder := squirrel.Select("COUNT(*)", "value").
		PlaceholderFormat(squirrel.Dollar).
		From("book_attributes").
		Where(squirrel.Eq{
			"attr": "author",
		}).
		GroupBy("value")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	repo.SquirrelDebugLog(ctx, query, args)

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
