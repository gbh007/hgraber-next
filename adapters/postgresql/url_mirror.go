package postgresql

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/parsing"
)

func (d *Database) NewMirror(ctx context.Context, mirror parsing.URLMirror) error {
	builder := squirrel.Insert("url_mirrors").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]interface{}{
			"id":          mirror.ID,
			"name":        mirror.Name,
			"prefixes":    mirror.Prefixes,
			"description": model.StringToDB(mirror.Description),
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	_, err = d.pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

func (d *Database) UpdateMirror(ctx context.Context, mirror parsing.URLMirror) error {
	builder := squirrel.Update("url_mirrors").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]interface{}{
			"name":        mirror.Name,
			"prefixes":    mirror.Prefixes,
			"description": model.StringToDB(mirror.Description),
		}).
		Where(squirrel.Eq{
			"id": mirror.ID,
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	_, err = d.pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

func (d *Database) DeleteMirror(ctx context.Context, id uuid.UUID) error {
	builder := squirrel.Delete("url_mirrors").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			"id": id,
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	_, err = d.pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

func (d *Database) Mirrors(ctx context.Context) ([]parsing.URLMirror, error) {
	builder := squirrel.Select(model.URLMirrorColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("url_mirrors").
		OrderBy("id")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	out := make([]parsing.URLMirror, 0, 10)

	rows, err := d.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query :%w", err)
	}

	defer rows.Close()

	for rows.Next() {
		mirror := parsing.URLMirror{}

		err := rows.Scan(model.URLMirrorScanner(&mirror))
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		out = append(out, mirror)
	}

	return out, nil
}

func (d *Database) Mirror(ctx context.Context, id uuid.UUID) (parsing.URLMirror, error) {
	builder := squirrel.Select(model.URLMirrorColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("url_mirrors").
		Where(squirrel.Eq{
			"id": id,
		}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return parsing.URLMirror{}, fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	row := d.pool.QueryRow(ctx, query, args...)

	mirror := parsing.URLMirror{}

	err = row.Scan(model.URLMirrorScanner(&mirror))
	if err != nil {
		return parsing.URLMirror{}, fmt.Errorf("scan: %w", err)
	}

	return mirror, nil
}
