//revive:disable:file-length-limit
package book

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

//nolint:gocognit,cyclop,funlen // будет исправлено позднее
func (repo *BookRepo) buildBooksFilter(
	_ context.Context,
	filter core.BookFilter,
	isCount bool,
) (string, []any, error) {
	var builder squirrel.SelectBuilder

	bookTable := model.BookTable.WithPrefix("b")
	pageTable := model.PageTable.WithPrefix("p")
	bookLabelTable := model.BookLabelTable.WithPrefix("l")
	bookAttributeTable := model.BookAttributeTable.WithPrefix("a")

	if isCount {
		builder = squirrel.Select("COUNT(" + bookTable.ColumnID() + ")")
	} else {
		builder = squirrel.Select(bookTable.ColumnID())
	}

	builder = builder.PlaceholderFormat(squirrel.Dollar).
		From(bookTable.NameAlter())

	if !isCount {
		if filter.Limit > 0 {
			builder = builder.Limit(uint64(filter.Limit))
		}

		if filter.Offset > 0 {
			builder = builder.Offset(uint64(filter.Offset))
		}

		var orderBySuffix string

		if filter.Desc {
			orderBySuffix = " DESC"
		} else {
			orderBySuffix = " ASC"
		}

		orderBy := []string{
			bookTable.ColumnCreateAt() + orderBySuffix,
			bookTable.ColumnID() + orderBySuffix,
		}

		switch filter.OrderBy {
		case core.BookFilterOrderByCreated:
			orderBy = []string{
				bookTable.ColumnCreateAt() + orderBySuffix,
				bookTable.ColumnID() + orderBySuffix,
			}

		case core.BookFilterOrderByName:
			orderBy = []string{
				bookTable.ColumnName() + orderBySuffix,
				bookTable.ColumnID() + orderBySuffix,
			}

		case core.BookFilterOrderByID:
			orderBy = []string{
				bookTable.ColumnID() + orderBySuffix,
			}

		case core.BookFilterOrderByPageCount:
			orderBy = []string{
				bookTable.ColumnPageCount() + orderBySuffix,
				bookTable.ColumnID() + orderBySuffix,
			}

		case core.BookFilterOrderByCalcPageCount:
			orderBy = []string{
				bookTable.ColumnCalcPageCount() + orderBySuffix + " NULLS LAST",
				bookTable.ColumnID() + orderBySuffix,
			}

		case core.BookFilterOrderByCalcFileCount:
			orderBy = []string{
				bookTable.ColumnCalcFileCount() + orderBySuffix + " NULLS LAST",
				bookTable.ColumnID() + orderBySuffix,
			}

		case core.BookFilterOrderByCalcDeadHashCount:
			orderBy = []string{
				bookTable.ColumnCalcDeadHashCount() + orderBySuffix + " NULLS LAST",
				bookTable.ColumnID() + orderBySuffix,
			}

		case core.BookFilterOrderByCalcPageSize:
			orderBy = []string{
				bookTable.ColumnCalcPageSize() + orderBySuffix + " NULLS LAST",
				bookTable.ColumnID() + orderBySuffix,
			}

		case core.BookFilterOrderByCalcFileSize:
			orderBy = []string{
				bookTable.ColumnCalcFileSize() + orderBySuffix + " NULLS LAST",
				bookTable.ColumnID() + orderBySuffix,
			}

		case core.BookFilterOrderByCalcDeadHashSize:
			orderBy = []string{
				bookTable.ColumnCalcDeadHashSize() + orderBySuffix + " NULLS LAST",
				bookTable.ColumnID() + orderBySuffix,
			}

		case core.BookFilterOrderByCalculatedAt:
			orderBy = []string{
				bookTable.ColumnCalculatedAt() + orderBySuffix + " NULLS LAST",
				bookTable.ColumnID() + orderBySuffix,
			}

		case core.BookFilterOrderByCalcAvgPageSize:
			orderBy = []string{
				bookTable.ColumnCalcAvgPageSize() + orderBySuffix + " NULLS LAST",
				bookTable.ColumnID() + orderBySuffix,
			}
		}

		builder = builder.OrderBy(orderBy...)
	}

	if !filter.From.IsZero() {
		builder = builder.Where(squirrel.GtOrEq{
			bookTable.ColumnCreateAt(): filter.From,
		})
	}

	if !filter.To.IsZero() {
		builder = builder.Where(squirrel.Lt{
			bookTable.ColumnCreateAt(): filter.To,
		})
	}

	switch filter.ShowDeleted {
	case core.BookFilterShowTypeOnly:
		builder = builder.Where(squirrel.Eq{bookTable.ColumnDeleted(): true})
	case core.BookFilterShowTypeExcept:
		builder = builder.Where(squirrel.Eq{bookTable.ColumnDeleted(): false})
	case core.BookFilterShowTypeAll:
		// Ничего не делаем
	}

	switch filter.ShowVerified {
	case core.BookFilterShowTypeOnly:
		builder = builder.Where(squirrel.Eq{bookTable.ColumnVerified(): true})
	case core.BookFilterShowTypeExcept:
		builder = builder.Where(squirrel.Eq{bookTable.ColumnVerified(): false})
	case core.BookFilterShowTypeAll:
		// Ничего не делаем
	}

	switch filter.ShowRebuilded {
	case core.BookFilterShowTypeOnly:
		builder = builder.Where(squirrel.Eq{bookTable.ColumnIsRebuild(): true})
	case core.BookFilterShowTypeExcept:
		builder = builder.Where(squirrel.Eq{bookTable.ColumnIsRebuild(): false})
	case core.BookFilterShowTypeAll:
		// Ничего не делаем
	}

	switch filter.ShowDownloaded {
	case core.BookFilterShowTypeOnly:
		builder = builder.Where(squirrel.And{
			squirrel.Eq{bookTable.ColumnAttributesParsed(): true},
			squirrel.NotEq{bookTable.ColumnName(): nil},
			squirrel.NotEq{bookTable.ColumnPageCount(): nil},
			squirrel.Expr("NOT EXISTS (SELECT 1 FROM " + pageTable.NameAlter() + " WHERE " + pageTable.ColumnDownloaded() + " = FALSE AND " + bookTable.ColumnID() + " = " + pageTable.ColumnBookID() + ")"), //nolint:golines,lll // будет исправлено позднее
		})
	case core.BookFilterShowTypeExcept:
		builder = builder.Where(squirrel.Or{
			squirrel.Eq{bookTable.ColumnAttributesParsed(): false},
			squirrel.Eq{bookTable.ColumnName(): nil},
			squirrel.Eq{bookTable.ColumnPageCount(): nil},
			squirrel.Expr("EXISTS (SELECT 1 FROM " + pageTable.NameAlter() + " WHERE " + pageTable.ColumnDownloaded() + " = FALSE AND " + bookTable.ColumnID() + " = " + pageTable.ColumnBookID() + ")"), //nolint:golines,lll // будет исправлено позднее
		})
	case core.BookFilterShowTypeAll:
		// Ничего не делаем
	}

	switch filter.ShowWithoutPages {
	case core.BookFilterShowTypeOnly:
		builder = builder.Where(
			squirrel.Expr("NOT EXISTS (SELECT 1 FROM " + pageTable.NameAlter() + " WHERE " + bookTable.ColumnID() + " = " + pageTable.ColumnBookID() + ")"), //nolint:golines,lll // будет исправлено позднее
		)
	case core.BookFilterShowTypeExcept:
		builder = builder.Where(
			squirrel.Expr("EXISTS (SELECT 1 FROM " + pageTable.NameAlter() + " WHERE " + bookTable.ColumnID() + " = " + pageTable.ColumnBookID() + ")"), //nolint:lll // будет исправлено позднее
		)
	case core.BookFilterShowTypeAll:
		// Ничего не делаем
	}

	switch filter.ShowWithoutPreview {
	case core.BookFilterShowTypeOnly:
		builder = builder.Where(
			squirrel.Expr(
				"NOT EXISTS (SELECT 1 FROM "+pageTable.NameAlter()+" WHERE "+bookTable.ColumnID()+" = "+pageTable.ColumnBookID()+" AND "+pageTable.ColumnPageNumber()+" = ?)", //nolint:lll // будет исправлено позднее
				core.PageNumberForPreview,
			), // особенность библиотеки, необходимо использовать `?`
		)
	case core.BookFilterShowTypeExcept:
		builder = builder.Where(
			squirrel.Expr(
				"EXISTS (SELECT 1 FROM "+pageTable.NameAlter()+" WHERE "+bookTable.ColumnID()+" = "+pageTable.ColumnBookID()+" AND "+pageTable.ColumnPageNumber()+" = ?)", //nolint:lll // будет исправлено позднее
				core.PageNumberForPreview,
			), // особенность библиотеки, необходимо использовать `?`
		)
	case core.BookFilterShowTypeAll:
		// Ничего не делаем
	}

	if filter.Fields.Name != "" {
		builder = builder.Where(squirrel.ILike{bookTable.ColumnName(): "%" + filter.Fields.Name + "%"})
	}

	for _, attrFilter := range filter.Fields.Attributes {
		subBuilder := squirrel.Select("1").
			// Важно: либа не может переконвертить другой тип форматирования для подзапроса!
			PlaceholderFormat(squirrel.Question).
			From(bookAttributeTable.NameAlter()).
			Where(squirrel.Eq{
				bookAttributeTable.ColumnAttr(): attrFilter.Code,
			}).
			Where(squirrel.Expr(bookAttributeTable.ColumnBookID() + " = " + bookTable.ColumnID()))

		switch attrFilter.Type {
		case core.BookFilterAttributeTypeLike:
			if len(attrFilter.Values) == 0 {
				continue
			}

			subBuilder = subBuilder.Where(squirrel.ILike{
				bookAttributeTable.ColumnValue(): "%" + attrFilter.Values[0] + "%",
			})

		case core.BookFilterAttributeTypeIn:
			if len(attrFilter.Values) == 0 {
				continue
			}

			subBuilder = subBuilder.Where(squirrel.Eq{
				bookAttributeTable.ColumnValue(): attrFilter.Values,
			})

		case core.BookFilterAttributeTypeCountEq:
			if attrFilter.Count != 0 { // Случай когда нужно чтобы не было данных
				subBuilder = subBuilder.Having(squirrel.Eq{
					"COUNT(" + bookAttributeTable.ColumnValue() + ")": attrFilter.Count,
				}).
					GroupBy(bookAttributeTable.ColumnAttr())
			}

		case core.BookFilterAttributeTypeCountGt:
			subBuilder = subBuilder.Having(squirrel.Gt{
				"COUNT(" + bookAttributeTable.ColumnValue() + ")": attrFilter.Count,
			}).
				GroupBy(bookAttributeTable.ColumnAttr())

		case core.BookFilterAttributeTypeCountLt:
			if attrFilter.Count != 1 { // Случай когда нужно чтобы не было данных
				subBuilder = subBuilder.Having(squirrel.Lt{
					"COUNT(" + bookAttributeTable.ColumnValue() + ")": attrFilter.Count,
				}).
					GroupBy(bookAttributeTable.ColumnAttr())
			}
		case core.BookFilterAttributeTypeNone:
			continue

		default:
			continue
		}

		subQuery, subArgs, err := subBuilder.ToSql()
		if err != nil {
			return "", nil, fmt.Errorf("build attribute sub query: %w", err)
		}

		// Случай когда нужно чтобы не было данных
		if (attrFilter.Count == 0 && attrFilter.Type == core.BookFilterAttributeTypeCountEq) ||
			(attrFilter.Count == 1 && attrFilter.Type == core.BookFilterAttributeTypeCountLt) {
			builder = builder.Where(squirrel.Expr(`NOT EXISTS (`+subQuery+`)`, subArgs...))
		} else {
			builder = builder.Where(squirrel.Expr(`EXISTS (`+subQuery+`)`, subArgs...))
		}
	}

	for _, labelFilter := range filter.Fields.Labels {
		subBuilder := squirrel.Select("1").
			// Важно: либа не может переконвертить другой тип форматирования для подзапроса!
			PlaceholderFormat(squirrel.Question).
			From(bookLabelTable.NameAlter()).
			Where(squirrel.Eq{
				bookLabelTable.Name(): labelFilter.Name,
			}).
			Where(squirrel.Expr(bookLabelTable.ColumnBookID() + " = " + bookTable.ColumnID()))

		switch labelFilter.Type {
		case core.BookFilterLabelTypeLike:
			if len(labelFilter.Values) == 0 {
				continue
			}

			subBuilder = subBuilder.Where(squirrel.ILike{
				bookLabelTable.ColumnValue(): "%" + labelFilter.Values[0] + "%",
			})

		case core.BookFilterLabelTypeIn:
			if len(labelFilter.Values) == 0 {
				continue
			}

			subBuilder = subBuilder.Where(squirrel.Eq{
				bookLabelTable.ColumnValue(): labelFilter.Values,
			})

		case core.BookFilterLabelTypeCountEq:
			if labelFilter.Count != 0 { // Случай когда нужно чтобы не было данных
				subBuilder = subBuilder.Having(squirrel.Eq{
					"COUNT(" + bookLabelTable.ColumnValue() + ")": labelFilter.Count,
				}).
					GroupBy(bookLabelTable.Name())
			}

		case core.BookFilterLabelTypeCountGt:
			subBuilder = subBuilder.Having(squirrel.Gt{
				"COUNT(" + bookLabelTable.ColumnValue() + ")": labelFilter.Count,
			}).
				GroupBy(bookLabelTable.Name())

		case core.BookFilterLabelTypeCountLt:
			if labelFilter.Count != 1 { // Случай когда нужно чтобы не было данных
				subBuilder = subBuilder.Having(squirrel.Lt{
					"COUNT(" + bookLabelTable.ColumnValue() + ")": labelFilter.Count,
				}).
					GroupBy(bookLabelTable.Name())
			}
		case core.BookFilterLabelTypeNone:
			continue
		default:
			continue
		}

		subQuery, subArgs, err := subBuilder.ToSql()
		if err != nil {
			return "", nil, fmt.Errorf("build label sub query: %w", err)
		}

		// Случай когда нужно чтобы не было данных
		if (labelFilter.Count == 0 && labelFilter.Type == core.BookFilterLabelTypeCountEq) ||
			(labelFilter.Count == 1 && labelFilter.Type == core.BookFilterLabelTypeCountLt) {
			builder = builder.Where(squirrel.Expr(`NOT EXISTS (`+subQuery+`)`, subArgs...))
		} else {
			builder = builder.Where(squirrel.Expr(`EXISTS (`+subQuery+`)`, subArgs...))
		}
	}

	query, args := builder.MustSql()

	return query, args, nil
}
