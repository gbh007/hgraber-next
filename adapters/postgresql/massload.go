package postgresql

import (
	"context"
	"fmt"
	"net/url"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

func (d *Database) CreateMassload(ctx context.Context, ml massloadmodel.Massload) (int, error) {
	builder := squirrel.Insert("massloads").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			"name":            ml.Name,
			"description":     model.StringToDB(ml.Description),
			"is_deduplicated": ml.IsDeduplicated,
			"created_at":      ml.CreatedAt,
		}).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	var id int

	row := d.pool.QueryRow(ctx, query, args...)

	err = row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("exec query: %w", err)
	}

	return id, nil
}

func (d *Database) UpdateMassload(ctx context.Context, ml massloadmodel.Massload) error {
	builder := squirrel.Update("massloads").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			"name":            ml.Name,
			"description":     model.StringToDB(ml.Description),
			"is_deduplicated": ml.IsDeduplicated,
			"updated_at":      model.TimeToDB(ml.UpdatedAt),
		}).
		Where(squirrel.Eq{
			"id": ml.ID,
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

func (d *Database) UpdateMassloadSize(ctx context.Context, ml massloadmodel.Massload) error {
	builder := squirrel.Update("massloads").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			"page_size":  model.NilInt64ToDB(ml.PageSize),
			"file_size":  model.NilInt64ToDB(ml.FileSize),
			"updated_at": model.TimeToDB(ml.UpdatedAt),
		}).
		Where(squirrel.Eq{
			"id": ml.ID,
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

func (d *Database) Massload(ctx context.Context, id int) (massloadmodel.Massload, error) {
	builder := squirrel.Select(model.MassloadColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("massloads").
		Where(squirrel.Eq{
			"id": id,
		}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return massloadmodel.Massload{}, fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	ml := massloadmodel.Massload{}

	row := d.pool.QueryRow(ctx, query, args...)

	err = row.Scan(model.MassloadScanner(&ml))
	if err != nil {
		return massloadmodel.Massload{}, fmt.Errorf("exec: %w", err)
	}

	return ml, nil
}

func (d *Database) Massloads(ctx context.Context) ([]massloadmodel.Massload, error) {
	builder := squirrel.Select(model.MassloadColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("massloads").
		OrderBy("id")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	result := make([]massloadmodel.Massload, 0)

	rows, err := d.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query :%w", err)
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

func (d *Database) DeleteMassload(ctx context.Context, id int) error {
	builder := squirrel.Delete("massloads").
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

func (d *Database) CreateMassloadExternalLink(ctx context.Context, id int, link massloadmodel.MassloadExternalLink) error {
	builder := squirrel.Insert("massload_external_links").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			"massload_id": id,
			"url":         link.URL.String(),
			"created_at":  link.CreatedAt,
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

func (d *Database) DeleteMassloadExternalLink(ctx context.Context, id int, u url.URL) error {
	builder := squirrel.Delete("massload_external_links").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			"massload_id": id,
			"url":         u.String(),
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

func (d *Database) MassloadExternalLinks(ctx context.Context, id int) ([]massloadmodel.MassloadExternalLink, error) {
	builder := squirrel.Select(model.MassloadExternalLinkColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("massload_external_links").
		Where(squirrel.Eq{
			"massload_id": id,
		}).
		OrderBy("created_at")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	result := make([]massloadmodel.MassloadExternalLink, 0)

	rows, err := d.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query :%w", err)
	}

	defer rows.Close()

	for rows.Next() {
		link := massloadmodel.MassloadExternalLink{}

		err := rows.Scan(model.MassloadExternalLinkScanner(&link))
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		result = append(result, link)
	}

	return result, nil
}

func (d *Database) CreateMassloadAttribute(ctx context.Context, id int, attr massloadmodel.MassloadAttribute) error {
	builder := squirrel.Insert("massload_attributes").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			"massload_id": id,
			"attr_code":   attr.AttrCode,
			"attr_value":  attr.AttrValue,
			"created_at":  attr.CreatedAt,
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

func (d *Database) UpdateMassloadAttributeSize(ctx context.Context, attr massloadmodel.MassloadAttribute) error {
	builder := squirrel.Update("massload_attributes").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			"page_size":  model.NilInt64ToDB(attr.PageSize),
			"file_size":  model.NilInt64ToDB(attr.FileSize),
			"updated_at": model.TimeToDB(attr.UpdatedAt),
		}).
		Where(squirrel.Eq{
			"attr_code":  attr.AttrCode,
			"attr_value": attr.AttrValue,
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

func (d *Database) DeleteMassloadAttribute(ctx context.Context, id int, attr massloadmodel.MassloadAttribute) error {
	builder := squirrel.Delete("massload_attributes").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			"massload_id": id,
			"attr_code":   attr.AttrCode,
			"attr_value":  attr.AttrValue,
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

func (d *Database) MassloadAttributes(ctx context.Context, id int) ([]massloadmodel.MassloadAttribute, error) {
	builder := squirrel.Select(model.MassloadAttributeColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("massload_attributes").
		Where(squirrel.Eq{
			"massload_id": id,
		}).
		OrderBy("created_at")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	result := make([]massloadmodel.MassloadAttribute, 0)

	rows, err := d.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query :%w", err)
	}

	defer rows.Close()

	for rows.Next() {
		attr := massloadmodel.MassloadAttribute{}

		err := rows.Scan(model.MassloadAttributeScanner(&attr))
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		result = append(result, attr)
	}

	return result, nil
}

func (d *Database) MassloadsAttributes(ctx context.Context) ([]massloadmodel.MassloadAttribute, error) {
	builder := squirrel.
		Select(
			"attr_code",
			"attr_value",
		).
		PlaceholderFormat(squirrel.Dollar).
		From("massload_attributes").
		GroupBy(
			"attr_code",
			"attr_value",
		)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	result := make([]massloadmodel.MassloadAttribute, 0)

	rows, err := d.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query :%w", err)
	}

	defer rows.Close()

	for rows.Next() {
		attr := massloadmodel.MassloadAttribute{}

		err := rows.Scan(&attr.AttrCode, &attr.AttrValue)
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		result = append(result, attr)
	}

	return result, nil
}
