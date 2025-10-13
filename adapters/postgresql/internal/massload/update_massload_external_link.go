package massload

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

func (repo *MassloadRepo) UpdateMassloadExternalLink(
	ctx context.Context,
	id int,
	link massloadmodel.ExternalLink,
) error {
	table := model.MassloadExternalLinkTable

	builder := squirrel.Update(table.Name()).
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			table.ColumnAutoCheck(): link.AutoCheck,
			table.ColumnUpdatedAt(): link.UpdatedAt,
		}).
		Where(squirrel.Eq{
			table.ColumnMassloadID(): id,
			table.ColumnURL():        link.URL.String(),
		})

	query, args := builder.MustSql()

	_, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
