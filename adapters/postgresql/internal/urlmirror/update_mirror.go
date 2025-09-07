package urlmirror

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/parsing"
)

func (repo *URLMirrorRepo) UpdateMirror(ctx context.Context, mirror parsing.URLMirror) error {
	builder := squirrel.Update("url_mirrors").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			"name":        mirror.Name,
			"prefixes":    mirror.Prefixes,
			"description": model.StringToDB(mirror.Description),
		}).
		Where(squirrel.Eq{
			"id": mirror.ID,
		})

	query, args := builder.MustSql()

	_, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
