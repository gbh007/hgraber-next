package postgresql

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"hgnext/internal/entities"
)

func (d *Database) BookCount(ctx context.Context, filter entities.BookFilter) (int, error) {
	var c int

	query, args, err := d.buildBooksFilter(ctx, filter, true)
	if err != nil {
		return 0, fmt.Errorf("build book filter: %w", err)
	}

	err = d.db.GetContext(ctx, &c, query, args...)
	if err != nil {
		return 0, fmt.Errorf("exec query: %w", err)
	}

	return c, nil
}

func (d *Database) BookIDs(ctx context.Context, filter entities.BookFilter) ([]uuid.UUID, error) {
	idsRaw := make([]string, 0)

	query, args, err := d.buildBooksFilter(ctx, filter, false)
	if err != nil {
		return nil, fmt.Errorf("build book filter: %w", err)
	}

	err = d.db.SelectContext(ctx, &idsRaw, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	ids := make([]uuid.UUID, len(idsRaw))

	for i, idRaw := range idsRaw {
		ids[i], err = uuid.Parse(idRaw)
		if err != nil {
			return nil, fmt.Errorf("parse uuid: %w", err)
		}
	}

	return ids, nil
}

func (d *Database) buildBooksFilter(ctx context.Context, filter entities.BookFilter, isCount bool) (string, []interface{}, error) {
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

		orderBy := "create_at"

		switch filter.OrderBy {
		case entities.BookFilterOrderByCreated:
			orderBy = "create_at"
		case entities.BookFilterOrderByName:
			orderBy = "name"
		case entities.BookFilterOrderByID:
			orderBy = "id"
		case entities.BookFilterOrderByPageCount:
			orderBy = "page_count"
		}

		if filter.Desc {
			builder = builder.OrderBy(orderBy + " DESC")
		} else {
			builder = builder.OrderBy(orderBy + " ASC")
		}
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
	case entities.BookFilterShowTypeOnly:
		builder = builder.Where(squirrel.Eq{"deleted": true})
	case entities.BookFilterShowTypeExcept:
		builder = builder.Where(squirrel.Eq{"deleted": false})
	}

	switch filter.ShowVerified {
	case entities.BookFilterShowTypeOnly:
		builder = builder.Where(squirrel.Eq{"verified": true})
	case entities.BookFilterShowTypeExcept:
		builder = builder.Where(squirrel.Eq{"verified": false})
	}

	switch filter.ShowDownloaded {
	case entities.BookFilterShowTypeOnly:
		builder = builder.Where(squirrel.And{
			squirrel.Eq{"attributes_parsed": true},
			squirrel.NotEq{"name": nil},
			squirrel.NotEq{"page_count": nil},
			squirrel.Expr(`NOT EXISTS (SELECT book_id FROM pages WHERE downloaded = FALSE AND books.id = pages.book_id)`),
		})
	case entities.BookFilterShowTypeExcept:
		builder = builder.Where(squirrel.Or{
			squirrel.Eq{"attributes_parsed": false},
			squirrel.Eq{"name": nil},
			squirrel.Eq{"page_count": nil},
			squirrel.Expr(`EXISTS (SELECT book_id FROM pages WHERE downloaded = FALSE AND books.id = pages.book_id)`),
		})
	}

	if filter.Fields.Name != "" {
		builder = builder.Where(squirrel.ILike{"name": "%" + filter.Fields.Name + "%"})
	}

	for _, attrFilter := range filter.Fields.Attributes {
		subBuilder := squirrel.Select("1").
			PlaceholderFormat(squirrel.Question). // Важно: либа не может переконвертить другой тип форматирования для подзапроса!
			From("book_attributes").
			Where(squirrel.Eq{
				"attr": attrFilter.Code,
			}).
			Where(squirrel.Expr(`book_id = books.id`))

		switch attrFilter.Type {
		case entities.BookFilterAttributeTypeLike:
			if len(attrFilter.Values) == 0 {
				continue
			}

			subBuilder = subBuilder.Where(squirrel.ILike{
				"value": "%" + attrFilter.Values[0] + "%",
			})

		case entities.BookFilterAttributeTypeIn:
			if len(attrFilter.Values) == 0 {
				continue
			}

			subBuilder = subBuilder.Where(squirrel.Eq{
				"value": attrFilter.Values,
			})

		case entities.BookFilterAttributeTypeCountEq:
			subBuilder = subBuilder.Having(squirrel.Eq{
				"COUNT(value)": attrFilter.Count,
			}).
				GroupBy("attr")

		case entities.BookFilterAttributeTypeCountGt:
			subBuilder = subBuilder.Having(squirrel.Gt{
				"COUNT(value)": attrFilter.Count,
			}).
				GroupBy("attr")

		case entities.BookFilterAttributeTypeCountLt:
			subBuilder = subBuilder.Having(squirrel.Lt{
				"COUNT(value)": attrFilter.Count,
			}).
				GroupBy("attr")

		default:
			continue
		}

		subQuery, subArgs, err := subBuilder.ToSql()
		if err != nil {
			return "", nil, fmt.Errorf("build attribute sub query: %w", err)
		}

		builder = builder.Where(squirrel.Expr(`EXISTS (`+subQuery+`)`, subArgs...))
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return "", nil, fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	return query, args, nil
}
