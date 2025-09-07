package attribute

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *AttributeRepo) AttributesPageSize(
	ctx context.Context,
	attrs map[string][]string,
) (core.SizeWithCount, error) {
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
		return core.SizeWithCount{}, errors.New("incorrect condition: empty attributes")
	}

	builder := squirrel.Select(`SUM(f."size")`, `COUNT(*)`).
		PlaceholderFormat(squirrel.Dollar).
		From(`files f`).
		InnerJoin(`pages p ON f.id = p.file_id`).
		InnerJoin(`book_attributes ba ON ba.book_id = p.book_id`).
		Where(whereCond)

	query, args := builder.MustSql()

	row := repo.Pool.QueryRow(ctx, query, args...)

	var size, count sql.NullInt64

	err := row.Scan(&size, &count)
	if err != nil {
		return core.SizeWithCount{}, fmt.Errorf("exec query: %w", err)
	}

	return core.SizeWithCount{
		Count: count.Int64,
		Size:  size.Int64,
	}, nil
}
