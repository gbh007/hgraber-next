package book

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *BookRepo) BookSizes(ctx context.Context) (map[uuid.UUID]core.SizeWithCount, error) {
	builder := squirrel.Select("COUNT(*)", "p.book_id", "SUM(f.size)").
		PlaceholderFormat(squirrel.Dollar).
		From("pages p").
		InnerJoin("files f ON f.id = p.file_id").
		GroupBy("p.book_id")

	query, args := builder.MustSql()

	out := make(map[uuid.UUID]core.SizeWithCount, 100) //nolint:mnd // оптимизация

	rows, err := repo.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var (
			count  sql.NullInt64
			size   sql.NullInt64
			bookID uuid.UUID
		)

		err = rows.Scan(&count, &bookID, &size)
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		out[bookID] = core.SizeWithCount{
			Count: count.Int64,
			Size:  size.Int64,
		}
	}

	return out, nil
}
