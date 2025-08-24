package massload

import (
	"context"
	"fmt"
	"net/url"

	"github.com/Masterminds/squirrel"
)

func (repo *MassloadRepo) DeleteMassloadExternalLink(ctx context.Context, id int, u url.URL) error {
	builder := squirrel.Delete("massload_external_links").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			"massload_id": id,
			"url":         u.String(),
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
