package book

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/domain/core"
)

//nolint:gocognit,cyclop,funlen // будет исправлено позднее
func (repo *BookRepo) buildBooksFilter(
	ctx context.Context,
	filter core.BookFilter,
	isCount bool,
) (string, []any, error) {
	var builder squirrel.SelectBuilder

	if isCount {
		builder = squirrel.Select("COUNT(id)")
	} else {
		builder = squirrel.Select("id")
	}

	builder = builder.PlaceholderFormat(squirrel.Dollar).
		From("books")

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
			"create_at" + orderBySuffix,
			"id" + orderBySuffix,
		}

		switch filter.OrderBy {
		case core.BookFilterOrderByCreated:
			orderBy = []string{
				"create_at" + orderBySuffix,
				"id" + orderBySuffix,
			}

		case core.BookFilterOrderByName:
			orderBy = []string{
				"name" + orderBySuffix,
				"id" + orderBySuffix,
			}

		case core.BookFilterOrderByID:
			orderBy = []string{
				"id" + orderBySuffix,
			}

		case core.BookFilterOrderByPageCount:
			orderBy = []string{
				"page_count" + orderBySuffix,
				"id" + orderBySuffix,
			}
		}

		builder = builder.OrderBy(orderBy...)
	}

	if !filter.From.IsZero() {
		builder = builder.Where(squirrel.GtOrEq{
			"create_at": filter.From,
		})
	}

	if !filter.To.IsZero() {
		builder = builder.Where(squirrel.Lt{
			"create_at": filter.To,
		})
	}

	switch filter.ShowDeleted {
	case core.BookFilterShowTypeOnly:
		builder = builder.Where(squirrel.Eq{"deleted": true})
	case core.BookFilterShowTypeExcept:
		builder = builder.Where(squirrel.Eq{"deleted": false})
	case core.BookFilterShowTypeAll:
		// Ничего не делаем
	}

	switch filter.ShowVerified {
	case core.BookFilterShowTypeOnly:
		builder = builder.Where(squirrel.Eq{"verified": true})
	case core.BookFilterShowTypeExcept:
		builder = builder.Where(squirrel.Eq{"verified": false})
	case core.BookFilterShowTypeAll:
		// Ничего не делаем
	}

	switch filter.ShowRebuilded {
	case core.BookFilterShowTypeOnly:
		builder = builder.Where(squirrel.Eq{"is_rebuild": true})
	case core.BookFilterShowTypeExcept:
		builder = builder.Where(squirrel.Eq{"is_rebuild": false})
	case core.BookFilterShowTypeAll:
		// Ничего не делаем
	}

	switch filter.ShowDownloaded {
	case core.BookFilterShowTypeOnly:
		builder = builder.Where(squirrel.And{
			squirrel.Eq{"attributes_parsed": true},
			squirrel.NotEq{"name": nil},
			squirrel.NotEq{"page_count": nil},
			squirrel.Expr(`NOT EXISTS (SELECT 1 FROM pages WHERE downloaded = FALSE AND books.id = pages.book_id)`),
		})
	case core.BookFilterShowTypeExcept:
		builder = builder.Where(squirrel.Or{
			squirrel.Eq{"attributes_parsed": false},
			squirrel.Eq{"name": nil},
			squirrel.Eq{"page_count": nil},
			squirrel.Expr(`EXISTS (SELECT 1 FROM pages WHERE downloaded = FALSE AND books.id = pages.book_id)`),
		})
	case core.BookFilterShowTypeAll:
		// Ничего не делаем
	}

	switch filter.ShowWithoutPages {
	case core.BookFilterShowTypeOnly:
		builder = builder.Where(
			squirrel.Expr(`NOT EXISTS (SELECT 1 FROM pages WHERE books.id = pages.book_id)`),
		)
	case core.BookFilterShowTypeExcept:
		builder = builder.Where(
			squirrel.Expr(`EXISTS (SELECT 1 FROM pages WHERE books.id = pages.book_id)`),
		)
	case core.BookFilterShowTypeAll:
		// Ничего не делаем
	}

	switch filter.ShowWithoutPreview {
	case core.BookFilterShowTypeOnly:
		builder = builder.Where(
			squirrel.Expr(
				`NOT EXISTS (SELECT 1 FROM pages WHERE books.id = pages.book_id AND pages.page_number = ?)`,
				core.PageNumberForPreview,
			), // особенность библиотеки, необходимо использовать `?`
		)
	case core.BookFilterShowTypeExcept:
		builder = builder.Where(
			squirrel.Expr(
				`EXISTS (SELECT 1 FROM pages WHERE books.id = pages.book_id AND pages.page_number = ?)`,
				core.PageNumberForPreview,
			), // особенность библиотеки, необходимо использовать `?`
		)
	case core.BookFilterShowTypeAll:
		// Ничего не делаем
	}

	if filter.Fields.Name != "" {
		builder = builder.Where(squirrel.ILike{"name": "%" + filter.Fields.Name + "%"})
	}

	for _, attrFilter := range filter.Fields.Attributes {
		subBuilder := squirrel.Select("1").
			// Важно: либа не может переконвертить другой тип форматирования для подзапроса!
			PlaceholderFormat(squirrel.Question).
			From("book_attributes").
			Where(squirrel.Eq{
				"attr": attrFilter.Code,
			}).
			Where(squirrel.Expr(`book_id = books.id`))

		switch attrFilter.Type {
		case core.BookFilterAttributeTypeLike:
			if len(attrFilter.Values) == 0 {
				continue
			}

			subBuilder = subBuilder.Where(squirrel.ILike{
				"value": "%" + attrFilter.Values[0] + "%",
			})

		case core.BookFilterAttributeTypeIn:
			if len(attrFilter.Values) == 0 {
				continue
			}

			subBuilder = subBuilder.Where(squirrel.Eq{
				"value": attrFilter.Values,
			})

		case core.BookFilterAttributeTypeCountEq:
			if attrFilter.Count != 0 { // Случай когда нужно чтобы не было данных
				subBuilder = subBuilder.Having(squirrel.Eq{
					"COUNT(value)": attrFilter.Count,
				}).
					GroupBy("attr")
			}

		case core.BookFilterAttributeTypeCountGt:
			subBuilder = subBuilder.Having(squirrel.Gt{
				"COUNT(value)": attrFilter.Count,
			}).
				GroupBy("attr")

		case core.BookFilterAttributeTypeCountLt:
			if attrFilter.Count != 1 { // Случай когда нужно чтобы не было данных
				subBuilder = subBuilder.Having(squirrel.Lt{
					"COUNT(value)": attrFilter.Count,
				}).
					GroupBy("attr")
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
			From("book_labels").
			Where(squirrel.Eq{
				"name": labelFilter.Name,
			}).
			Where(squirrel.Expr(`book_id = books.id`))

		switch labelFilter.Type {
		case core.BookFilterLabelTypeLike:
			if len(labelFilter.Values) == 0 {
				continue
			}

			subBuilder = subBuilder.Where(squirrel.ILike{
				"value": "%" + labelFilter.Values[0] + "%",
			})

		case core.BookFilterLabelTypeIn:
			if len(labelFilter.Values) == 0 {
				continue
			}

			subBuilder = subBuilder.Where(squirrel.Eq{
				"value": labelFilter.Values,
			})

		case core.BookFilterLabelTypeCountEq:
			if labelFilter.Count != 0 { // Случай когда нужно чтобы не было данных
				subBuilder = subBuilder.Having(squirrel.Eq{
					"COUNT(value)": labelFilter.Count,
				}).
					GroupBy("name")
			}

		case core.BookFilterLabelTypeCountGt:
			subBuilder = subBuilder.Having(squirrel.Gt{
				"COUNT(value)": labelFilter.Count,
			}).
				GroupBy("name")

		case core.BookFilterLabelTypeCountLt:
			if labelFilter.Count != 1 { // Случай когда нужно чтобы не было данных
				subBuilder = subBuilder.Having(squirrel.Lt{
					"COUNT(value)": labelFilter.Count,
				}).
					GroupBy("name")
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

	query, args, err := builder.ToSql()
	if err != nil {
		return "", nil, fmt.Errorf("build query: %w", err)
	}

	repo.SquirrelDebugLog(ctx, query, args)

	return query, args, nil
}

//nolint:revive // будет исправлено позднее
