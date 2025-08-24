package attribute

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
)

func (repo *AttributeRepo) AttributesFileSize(ctx context.Context, attrs map[string][]string) (int64, error) {
	whereCond := squirrel.Or{}

	for code, values := range attrs {
		if len(values) == 0 {
			continue
		}

		whereCond = append(whereCond, squirrel.Eq{
			"ba.attr":  code,
			"ba.value": values,
		})
	}

	if len(whereCond) == 0 {
		return 0, errors.New("incorrect condition: empty attributes")
	}

	subBuilder := squirrel.Select(
		`f."size"`,
		`f.md5_sum`,
		`f.sha256_sum`,
	).
		// Важно: либа не может переконвертить другой тип форматирования для подзапроса!
		PlaceholderFormat(squirrel.Question).
		From(`files f`).
		InnerJoin(`pages p ON f.id = p.file_id`).
		InnerJoin(`book_attributes ba ON ba.book_id = p.book_id`).
		Where(whereCond).
		GroupBy(
			`f."size"`,
			`f.md5_sum`,
			`f.sha256_sum`,
		)

	builder := squirrel.Select(`sum(uf."size")`).
		PlaceholderFormat(squirrel.Dollar).
		FromSelect(subBuilder, "uf")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, fmt.Errorf("build query: %w", err)
	}

	repo.SquirrelDebugLog(ctx, query, args)

	row := repo.Pool.QueryRow(ctx, query, args...)

	var size sql.NullInt64

	err = row.Scan(&size)
	if err != nil {
		return 0, fmt.Errorf("exec query: %w", err)
	}

	return size.Int64, nil
}
