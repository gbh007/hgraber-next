package page

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *PageRepo) PageSizeByAuthor(ctx context.Context) (map[string]core.SizeWithCount, error) {
	builder := squirrel.Select("COUNT(*)", "ba.value", "SUM(f.size)").
		PlaceholderFormat(squirrel.Dollar).
		From("book_attributes ba").
		InnerJoin("pages p ON ba.book_id = p.book_id").
		InnerJoin("files f ON f.id = p.file_id").
		Where(squirrel.Eq{
			"ba.attr": "author",
		}).
		GroupBy("ba.value")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	repo.SquirrelDebugLog(ctx, query, args)

	out := make(map[string]core.SizeWithCount, 100) //nolint:mnd // оптимизация

	rows, err := repo.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var (
			count sql.NullInt64
			size  sql.NullInt64
			name  string
		)

		err = rows.Scan(&count, &name, &size)
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		out[name] = core.SizeWithCount{
			Count: count.Int64,
			Size:  size.Int64,
		}
	}

	return out, nil
}
