package massload

import (
	"context"
	"fmt"
	"net/url"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
)

func (repo *MassloadRepo) DeleteMassloadExternalLink(ctx context.Context, id int, u url.URL) error {
	table := model.MassloadExternalLinkTable

	builder := squirrel.Delete(table.Name()).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			table.ColumnMassloadID(): id,
			table.ColumnURL():        u.String(),
		})

	query, args := builder.MustSql()

	_, err := repo.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
