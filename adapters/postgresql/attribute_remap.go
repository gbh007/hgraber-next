package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (d *Database) InsertAttributeRemap(ctx context.Context, ar core.AttributeRemap) error {
	builder := squirrel.Insert("attribute_remaps").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			"attr":       ar.Code,
			"value":      ar.Value,
			"to_attr":    model.StringToDB(ar.ToCode),
			"to_value":   model.StringToDB(ar.ToValue),
			"created_at": ar.CreatedAt,
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	d.SquirrelDebugLog(ctx, query, args)

	_, err = d.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

func (d *Database) UpdateAttributeRemap(ctx context.Context, ar core.AttributeRemap) error {
	builder := squirrel.Update("attribute_remaps").
		PlaceholderFormat(squirrel.Dollar).
		SetMap(map[string]any{
			"to_attr":    model.StringToDB(ar.ToCode),
			"to_value":   model.StringToDB(ar.ToValue),
			"updated_at": model.TimeToDB(ar.UpdateAt),
		}).
		Where(squirrel.Eq{
			"attr":  ar.Code,
			"value": ar.Value,
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	d.SquirrelDebugLog(ctx, query, args)

	_, err = d.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

func (d *Database) DeleteAttributeRemap(ctx context.Context, code, value string) error {
	builder := squirrel.Delete("attribute_remaps").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{
			"attr":  code,
			"value": value,
		})

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	d.SquirrelDebugLog(ctx, query, args)

	_, err = d.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

func (d *Database) AttributeRemaps(ctx context.Context) ([]core.AttributeRemap, error) {
	builder := squirrel.Select(model.AttributeRemapColumns()...).
		From("attribute_remaps").
		PlaceholderFormat(squirrel.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	d.SquirrelDebugLog(ctx, query, args)

	rows, err := d.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	result := make([]core.AttributeRemap, 0)

	for rows.Next() {
		ar := core.AttributeRemap{}

		err = rows.Scan(model.AttributeRemapScanner(&ar))
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		result = append(result, ar)
	}

	return result, nil
}

func (d *Database) AttributeRemap(ctx context.Context, code, value string) (core.AttributeRemap, error) {
	builder := squirrel.Select(model.AttributeRemapColumns()...).
		PlaceholderFormat(squirrel.Dollar).
		From("attribute_remaps").
		Where(squirrel.Eq{
			"attr":  code,
			"value": value,
		}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return core.AttributeRemap{}, fmt.Errorf("build query: %w", err)
	}

	d.SquirrelDebugLog(ctx, query, args)

	row := d.Pool.QueryRow(ctx, query, args...)

	ar := core.AttributeRemap{}

	err = row.Scan(model.AttributeRemapScanner(&ar))
	if errors.Is(err, sql.ErrNoRows) {
		return core.AttributeRemap{}, core.AttributeRemapNotFoundError
	}

	if err != nil {
		return core.AttributeRemap{}, fmt.Errorf("scan row: %w", err)
	}

	return ar, nil
}
