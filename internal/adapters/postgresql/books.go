package postgresql

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"hgnext/internal/entities"
)

func (d *Database) BookCount(ctx context.Context) (int, error) {
	var c int

	// TODO: заменить на более оптимальную, если с ней не будет проблем
	err := d.db.GetContext(ctx, &c, `SELECT COUNT(*) FROM books WHERE deleted = FALSE;`)
	if err != nil {
		return 0, err
	}

	return c, nil
}

func (d *Database) BookIDs(ctx context.Context, filter entities.BookFilter) ([]uuid.UUID, error) {
	idsRaw := make([]string, 0)

	builder := squirrel.Select("id").
		PlaceholderFormat(squirrel.Dollar).
		From("books").
		Where(squirrel.Eq{"deleted": false})

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

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

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
