package urlmirror

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/parsing"
)

func (repo *URLMirrorRepo) NewMirror(ctx context.Context, mirror parsing.URLMirror) error {
	builder := squirrel.Insert("url_mirrors").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			"id":          mirror.ID,
			"name":        mirror.Name,
			"prefixes":    mirror.Prefixes,
			"description": model.StringToDB(mirror.Description),
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	repo.SquirrelDebugLog(ctx, query, args)

	_, err = repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
