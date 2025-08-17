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
			"name":        ml.Name,
			"description": model.StringToDB(ml.Description),
			"flags":       ml.Flags,
			"created_at":  ml.CreatedAt,
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
			"name":        ml.Name,
			"description": model.StringToDB(ml.Description),
			"flags":       ml.Flags,
			"updated_at":  model.TimeToDB(ml.UpdatedAt),
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

func (d *Database) Massloads(ctx context.Context, filter massloadmodel.Filter) ([]massloadmodel.Massload, error) {
	builder := squirrel.Select(model.MassloadColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("massloads")

	orderBySuffix := ""

	if filter.Desc {
		orderBySuffix = " DESC"
	} else {
		orderBySuffix = " ASC"
	}

	orderBy := []string{
		"id" + orderBySuffix,
	}

	switch filter.OrderBy {
	case massloadmodel.FilterOrderByID:
		orderBy = []string{
			"id" + orderBySuffix,
		}

	case massloadmodel.FilterOrderByName:
		orderBy = []string{
			"name" + orderBySuffix,
			"id" + orderBySuffix,
		}

	case massloadmodel.FilterOrderByPageSize:
		orderBy = []string{
			"page_size" + orderBySuffix + " NULLS LAST",
			"id" + orderBySuffix,
		}

	case massloadmodel.FilterOrderByFileSize:
		orderBy = []string{
			"file_size" + orderBySuffix + " NULLS LAST",
			"id" + orderBySuffix,
		}
	}

	builder = builder.OrderBy(orderBy...)

	if filter.Fields.Name != "" {
		builder = builder.Where(squirrel.ILike{"name": "%" + filter.Fields.Name + "%"})
	}

	if len(filter.Fields.Flags) > 0 {
		builder = builder.Where(squirrel.Expr("flags @> ?", filter.Fields.Flags)) // особенность библиотеки, необходимо использовать `?`
	}

	if filter.Fields.ExternalLink != "" {
		builder = builder.Where(squirrel.Expr("EXISTS (SELECT FROM massload_external_links WHERE massload_id = id AND url ILIKE ?)", "%"+filter.Fields.Name+"%")) // особенность библиотеки, необходимо использовать `?`
	}

	for _, attrFilter := range filter.Fields.Attributes {
		subBuilder := squirrel.Select("1").
			PlaceholderFormat(squirrel.Question). // Важно: либа не может переконвертить другой тип форматирования для подзапроса!
			From("massload_attributes").
			Where(squirrel.Eq{
				"attr_code": attrFilter.Code,
			}).
			Where(squirrel.Expr(`massload_id = id`))

		switch attrFilter.Type {
		case massloadmodel.FilterAttributeTypeLike:
			if len(attrFilter.Values) == 0 {
				continue
			}

			subBuilder = subBuilder.Where(squirrel.ILike{
				"attr_value": "%" + attrFilter.Values[0] + "%",
			})

		case massloadmodel.FilterAttributeTypeIn:
			if len(attrFilter.Values) == 0 {
				continue
			}

			subBuilder = subBuilder.Where(squirrel.Eq{
				"attr_value": attrFilter.Values,
			})

		default:
			continue
		}

		subQuery, subArgs, err := subBuilder.ToSql()
		if err != nil {
			return nil, fmt.Errorf("build attribute sub query: %w", err)
		}

		builder = builder.Where(squirrel.Expr(`EXISTS (`+subQuery+`)`, subArgs...))
	}

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

func (d *Database) MassloadFlags(ctx context.Context) ([]massloadmodel.Flag, error) {
	builder := squirrel.Select(model.MassloadFlagColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("massload_flags").
		OrderBy("created_at", "code")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	result := make([]massloadmodel.Flag, 0)

	rows, err := d.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query :%w", err)
	}

	defer rows.Close()

	for rows.Next() {
		flag := massloadmodel.Flag{}

		err := rows.Scan(model.MassloadFlagScanner(&flag))
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		result = append(result, flag)
	}

	return result, nil
}

func (d *Database) CreateMassloadExternalLink(ctx context.Context, id int, link massloadmodel.ExternalLink) error {
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

func (d *Database) MassloadExternalLinks(ctx context.Context, id int) ([]massloadmodel.ExternalLink, error) {
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

	result := make([]massloadmodel.ExternalLink, 0)

	rows, err := d.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query :%w", err)
	}

	defer rows.Close()

	for rows.Next() {
		link := massloadmodel.ExternalLink{}

		err := rows.Scan(model.MassloadExternalLinkScanner(&link))
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		result = append(result, link)
	}

	return result, nil
}

func (d *Database) CreateMassloadAttribute(ctx context.Context, id int, attr massloadmodel.Attribute) error {
	builder := squirrel.Insert("massload_attributes").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			"massload_id": id,
			"attr_code":   attr.Code,
			"attr_value":  attr.Value,
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

func (d *Database) UpdateMassloadAttributeSize(ctx context.Context, attr massloadmodel.Attribute) error {
	builder := squirrel.Update("massload_attributes").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			"page_size":  model.NilInt64ToDB(attr.PageSize),
			"file_size":  model.NilInt64ToDB(attr.FileSize),
			"updated_at": model.TimeToDB(attr.UpdatedAt),
		}).
		Where(squirrel.Eq{
			"attr_code":  attr.Code,
			"attr_value": attr.Value,
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

func (d *Database) DeleteMassloadAttribute(ctx context.Context, id int, attr massloadmodel.Attribute) error {
	builder := squirrel.Delete("massload_attributes").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			"massload_id": id,
			"attr_code":   attr.Code,
			"attr_value":  attr.Value,
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

func (d *Database) MassloadAttributes(ctx context.Context, id int) ([]massloadmodel.Attribute, error) {
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

	result := make([]massloadmodel.Attribute, 0)

	rows, err := d.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query :%w", err)
	}

	defer rows.Close()

	for rows.Next() {
		attr := massloadmodel.Attribute{}

		err := rows.Scan(model.MassloadAttributeScanner(&attr))
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		result = append(result, attr)
	}

	return result, nil
}

func (d *Database) MassloadsAttributes(ctx context.Context) ([]massloadmodel.Attribute, error) {
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

	result := make([]massloadmodel.Attribute, 0)

	rows, err := d.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query :%w", err)
	}

	defer rows.Close()

	for rows.Next() {
		attr := massloadmodel.Attribute{}

		err := rows.Scan(&attr.Code, &attr.Value)
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		result = append(result, attr)
	}

	return result, nil
}
