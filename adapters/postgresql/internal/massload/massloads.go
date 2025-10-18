//revive:disable:file-length-limit
package massload

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

func (repo *MassloadRepo) Massloads(
	ctx context.Context,
	filter massloadmodel.Filter,
) ([]massloadmodel.Massload, error) {
	table := model.MassloadTable

	builder, err := repo.massloadsBuilder(filter)
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

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

//nolint:cyclop,funlen // будет исправлено позднее
func (repo *MassloadRepo) massloadsBuilder(
	filter massloadmodel.Filter,
) (squirrel.SelectBuilder, error) {
	attrTable := model.MassloadAttributeTable
	linkTable := model.MassloadExternalLinkTable
	table := model.MassloadTable

	builder := squirrel.Select(table.Columns()...).
		PlaceholderFormat(squirrel.Dollar).
		From(table.Name())

	var orderBySuffix string

	if filter.Desc {
		orderBySuffix = " DESC"
	} else {
		orderBySuffix = " ASC"
	}

	orderBy := []string{
		table.ColumnID() + orderBySuffix,
	}

	switch filter.OrderBy {
	case massloadmodel.FilterOrderByID:
		orderBy = []string{
			table.ColumnID() + orderBySuffix,
		}

	case massloadmodel.FilterOrderByName:
		orderBy = []string{
			table.ColumnName() + orderBySuffix,
			table.ColumnID() + orderBySuffix,
		}

	case massloadmodel.FilterOrderByPageSize:
		orderBy = []string{
			table.ColumnPageSize() + orderBySuffix + " NULLS LAST",
			table.ColumnID() + orderBySuffix,
		}

	case massloadmodel.FilterOrderByFileSize:
		orderBy = []string{
			table.ColumnFileSize() + orderBySuffix + " NULLS LAST",
			table.ColumnID() + orderBySuffix,
		}

	case massloadmodel.FilterOrderByPageCount:
		orderBy = []string{
			table.ColumnPageCount() + orderBySuffix + " NULLS LAST",
			table.ColumnID() + orderBySuffix,
		}

	case massloadmodel.FilterOrderByFileCount:
		orderBy = []string{
			table.ColumnFileCount() + orderBySuffix + " NULLS LAST",
			table.ColumnID() + orderBySuffix,
		}

	case massloadmodel.FilterOrderByBooksAhead:
		orderBy = []string{
			table.ColumnBooksAhead() + orderBySuffix + " NULLS LAST",
			table.ColumnID() + orderBySuffix,
		}

	case massloadmodel.FilterOrderByNewBooks:
		orderBy = []string{
			table.ColumnNewBooks() + orderBySuffix + " NULLS LAST",
			table.ColumnID() + orderBySuffix,
		}

	case massloadmodel.FilterOrderByExistingBooks:
		orderBy = []string{
			table.ColumnExistingBooks() + orderBySuffix + " NULLS LAST",
			table.ColumnID() + orderBySuffix,
		}

	case massloadmodel.FilterOrderByBooksInSystem:
		orderBy = []string{
			table.ColumnBooksInSystem() + orderBySuffix + " NULLS LAST",
			table.ColumnID() + orderBySuffix,
		}
	}

	builder = builder.OrderBy(orderBy...)

	if filter.Fields.Name != "" {
		builder = builder.Where(squirrel.ILike{table.ColumnName(): "%" + filter.Fields.Name + "%"})
	}

	if len(filter.Fields.Flags) > 0 {
		builder = builder.Where(
			squirrel.Expr(table.ColumnFlags()+" @> ?", filter.Fields.Flags),
		) // особенность библиотеки, необходимо использовать `?`
	}

	if len(filter.Fields.ExcludedFlags) > 0 {
		builder = builder.Where(
			squirrel.Expr("NOT "+table.ColumnFlags()+" && ?", filter.Fields.ExcludedFlags),
		) // особенность библиотеки, необходимо использовать `?`
	}

	if filter.Fields.ExternalLink != "" {
		subBuilder := squirrel.Select("1").
			// Важно: либа не может переконвертить другой тип форматирования для подзапроса!
			PlaceholderFormat(squirrel.Question).
			From(linkTable.Name()).
			Where(squirrel.Expr(linkTable.ColumnMassloadID() + " = " + table.ColumnID())).
			Where(squirrel.ILike{
				linkTable.ColumnURL(): "%" + filter.Fields.ExternalLink + "%",
			}).
			Prefix("EXISTS (").
			Suffix(")")

		subQuery, subArgs, err := subBuilder.ToSql()
		if err != nil {
			return squirrel.Select(), fmt.Errorf("build external link sub query: %w", err)
		}

		builder = builder.Where(subQuery, subArgs...)
	}

	for _, attrFilter := range filter.Fields.Attributes {
		subBuilder := squirrel.Select("1").
			// Важно: либа не может переконвертить другой тип форматирования для подзапроса!
			PlaceholderFormat(squirrel.Question).
			From(attrTable.Name()).
			Where(squirrel.Eq{
				attrTable.ColumnAttrCode(): attrFilter.Code,
			}).
			Where(squirrel.Expr(attrTable.ColumnMassloadID() + " = " + table.ColumnID())).
			Prefix("EXISTS (").
			Suffix(")")

		switch attrFilter.Type {
		case massloadmodel.FilterAttributeTypeLike:
			if len(attrFilter.Values) == 0 {
				continue
			}

			subBuilder = subBuilder.Where(squirrel.ILike{
				attrTable.ColumnAttrValue(): "%" + attrFilter.Values[0] + "%",
			})

		case massloadmodel.FilterAttributeTypeIn:
			if len(attrFilter.Values) == 0 {
				continue
			}

			subBuilder = subBuilder.Where(squirrel.Eq{
				attrTable.ColumnAttrValue(): attrFilter.Values,
			})

		case massloadmodel.FilterAttributeTypeNone:
			continue
		default:
			continue
		}

		subQuery, subArgs, err := subBuilder.ToSql()
		if err != nil {
			return squirrel.Select(), fmt.Errorf("build attribute sub query: %w", err)
		}

		builder = builder.Where(subQuery, subArgs...)
	}

	return builder, nil
}
