package book

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (repo *BookRepo) BookIDsByMD5(ctx context.Context, md5Sums []string) ([]uuid.UUID, error) {
	builder := squirrel.Select("b.id").
		PlaceholderFormat(squirrel.Dollar).
		From("books b").
		InnerJoin("pages p ON p.book_id = b.id").
		InnerJoin("files f ON f.id = p.file_id").
		Where(squirrel.Eq{
			"f.md5_sum": md5Sums,
		}).
		GroupBy("b.id")

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
