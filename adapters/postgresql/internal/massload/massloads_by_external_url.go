package massload

import (
	"context"
	"fmt"
	"net/url"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

func (repo *MassloadRepo) MassloadsByExternalLink(
	ctx context.Context,
	u url.URL,
) ([]massloadmodel.Massload, error) {
	builder := squirrel.Select(model.MassloadColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("massloads").
		Where(
			squirrel.Expr(
				"EXISTS (SELECT FROM massload_external_links WHERE massload_id = id AND url = ?)",
				u.String(),
			),
		)

	query, args := builder.MustSql()

	result := make([]massloadmodel.Massload, 0)

	rows, err := repo.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		ml := massloadmodel.Massload{}

		err := rows.Scan(model.MassloadScanner(&ml))
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		result = append(result, ml)
	}

	return result, nil
}
