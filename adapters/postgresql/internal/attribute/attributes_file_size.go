package attribute

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (repo *AttributeRepo) AttributesFileSize(
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

	builder := squirrel.Select(`SUM(uf."size")`, `COUNT(*)`).
		PlaceholderFormat(squirrel.Dollar).
		FromSelect(subBuilder, "uf")

	query, args, err := builder.ToSql()
	if err != nil {
		return core.SizeWithCount{}, fmt.Errorf("build query: %w", err)
	}

	repo.SquirrelDebugLog(ctx, query, args)

	row := repo.Pool.QueryRow(ctx, query, args...)

	var size, count sql.NullInt64

	err = row.Scan(&size, &count)
	if err != nil {
		return core.SizeWithCount{}, fmt.Errorf("exec query: %w", err)
	}

	return core.SizeWithCount{
		Count: count.Int64,
		Size:  size.Int64,
	}, nil
}
