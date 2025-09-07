package massload

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

func (repo *MassloadRepo) MassloadExternalLinks(ctx context.Context, id int) ([]massloadmodel.ExternalLink, error) {
	builder := squirrel.Select(model.MassloadExternalLinkColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("massload_external_links").
		Where(squirrel.Eq{
			"massload_id": id,
		}).
		OrderBy("created_at")

	query, args := builder.MustSql()

	result := make([]massloadmodel.ExternalLink, 0)

	rows, err := repo.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		link := massloadmodel.ExternalLink{}

		err := rows.Scan(model.MassloadExternalLinkScanner(&link))
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		result = append(result, link)
	}

	return result, nil
}
