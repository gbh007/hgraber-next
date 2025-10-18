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
	linkTable := model.MassloadExternalLinkTable
	table := model.MassloadTable

	subQuery, subArgs := squirrel.Select("1").
		From(linkTable.Name()).
		Where(squirrel.Expr(linkTable.ColumnMassloadID() + " = " + table.ColumnID())).
		Where(squirrel.Eq{
			linkTable.ColumnURL(): u.String(),
		}).
		Prefix("EXISTS (").
		Suffix(")").
		MustSql()

	builder := squirrel.Select(table.Columns()...).
		PlaceholderFormat(squirrel.Dollar).
		From(table.Name()).
		Where(subQuery, subArgs...)

	query, args := builder.MustSql()

	result := make([]massloadmodel.Massload, 0)

	rows, err := repo.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		ml := massloadmodel.Massload{}

		err := rows.Scan(table.Scanner(&ml))
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		result = append(result, ml)
	}

	return result, nil
}
