package massload

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

func (repo *MassloadRepo) CreateMassloadExternalLink(
	ctx context.Context,
	id int,
	link massloadmodel.ExternalLink,
) error {
	builder := squirrel.Insert("massload_external_links").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			"massload_id": id,
			"url":         link.URL.String(),
			"auto_check":  link.AutoCheck,
			"created_at":  link.CreatedAt,
		})

	query, args := builder.MustSql()

	_, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
