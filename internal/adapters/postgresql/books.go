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

		if filter.NewFirst {
			builder = builder.OrderBy("create_at DESC")
		} else {
			builder = builder.OrderBy("create_at ASC")
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

	if filter.Fields.Name != "" {
		builder = builder.Where(squirrel.ILike{"name": "%" + filter.Fields.Name + "%"})
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return "", nil, fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	return query, args, nil
}