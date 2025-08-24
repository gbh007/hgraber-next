package urlmirror

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/parsing"
)

func (repo *URLMirrorRepo) Mirror(ctx context.Context, id uuid.UUID) (parsing.URLMirror, error) {
	builder := squirrel.Select(model.URLMirrorColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("url_mirrors").
		Where(squirrel.Eq{
			"id": id,
		}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return parsing.URLMirror{}, fmt.Errorf("build query: %w", err)
	}

	repo.SquirrelDebugLog(ctx, query, args)

	row := repo.Pool.QueryRow(ctx, query, args...)

	mirror := parsing.URLMirror{}

	err = row.Scan(model.URLMirrorScanner(&mirror))
	if err != nil {
		return parsing.URLMirror{}, fmt.Errorf("scan: %w", err)
	}

	return mirror, nil
}
