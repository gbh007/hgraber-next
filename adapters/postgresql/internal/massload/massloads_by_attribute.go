package massload

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

func (repo *MassloadRepo) MassloadsByAttribute(
	ctx context.Context,
	code, value string,
) ([]massloadmodel.Massload, error) {
	attrTable := model.MassloadAttributeTable
	table := model.MassloadTable

	subQuery, subArgs := squirrel.Select("1").
		From(attrTable.Name()).
		Where(squirrel.Expr(attrTable.ColumnMassloadID() + " = " + table.ColumnID())).
		Where(squirrel.Eq{
			attrTable.ColumnAttrCode():  code,
			attrTable.ColumnAttrValue(): value,
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
