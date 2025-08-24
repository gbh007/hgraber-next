package attribute

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
)

func (repo *AttributeRepo) AttributesPageSize(ctx context.Context, attrs map[string][]string) (int64, error) {
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

	builder := squirrel.Select(`sum(f."size")`).
		PlaceholderFormat(squirrel.Dollar).
		From(`files f`).
		InnerJoin(`pages p ON f.id = p.file_id`).
		InnerJoin(`book_attributes ba ON ba.book_id = p.book_id`).
		Where(whereCond)

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
