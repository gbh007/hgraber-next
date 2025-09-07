package massload

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

func (repo *MassloadRepo) UpdateMassloadExternalLink(
	ctx context.Context,
	id int,
	link massloadmodel.ExternalLink,
) error {
	builder := squirrel.Update("massload_external_links").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			"auto_check": link.AutoCheck,
			"updated_at": link.UpdatedAt,
		}).
		Where(squirrel.Eq{
			"massload_id": id,
			"url":         link.URL.String(),
		})

	query, args := builder.MustSql()

	_, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
