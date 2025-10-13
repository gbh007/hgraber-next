package massload

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

func (repo *MassloadRepo) CreateMassloadExternalLink(
	ctx context.Context,
	id int,
	link massloadmodel.ExternalLink,
) error {
	table := model.MassloadExternalLinkTable

	builder := squirrel.Insert(table.Name()).
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			table.ColumnMassloadID(): id,
			table.ColumnURL():        link.URL.String(),
			table.ColumnAutoCheck():  link.AutoCheck,
			table.ColumnCreatedAt():  link.CreatedAt,
		})

	query, args := builder.MustSql()

	_, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
