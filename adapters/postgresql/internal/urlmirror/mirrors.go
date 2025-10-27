package urlmirror

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/parsing"
)

func (repo *URLMirrorRepo) Mirrors(ctx context.Context) ([]parsing.URLMirror, error) {
	table := model.URLMirrorTable

	builder := squirrel.Select(table.Columns()...).
		PlaceholderFormat(squirrel.Dollar).
		From(table.Name()).
		OrderBy(table.ColumnID())

	query, args := builder.MustSql()

	out := make([]parsing.URLMirror, 0, 10) //nolint:mnd // оптимизация

	rows, err := repo.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		mirror := parsing.URLMirror{}

		err := rows.Scan(table.Scanner(&mirror))
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		out = append(out, mirror)
	}

	return out, nil
}
