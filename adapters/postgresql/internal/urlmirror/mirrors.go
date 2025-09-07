package urlmirror

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/parsing"
)

func (repo *URLMirrorRepo) Mirrors(ctx context.Context) ([]parsing.URLMirror, error) {
	builder := squirrel.Select(model.URLMirrorColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("url_mirrors").
		OrderBy("id")

	query, args := builder.MustSql()

	out := make([]parsing.URLMirror, 0, 10) //nolint:mnd // оптимизация

	rows, err := repo.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		mirror := parsing.URLMirror{}

		err := rows.Scan(model.URLMirrorScanner(&mirror))
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		out = append(out, mirror)
	}

	return out, nil
}
