package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/adapters/postgresql/internal/model"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (d *Database) BookAttributes(ctx context.Context, bookID uuid.UUID) (map[string][]string, error) {
	rows, err := d.pool.Query(ctx, `SELECT attr, value FROM book_attributes WHERE book_id = $1;`, bookID)
	if err != nil {
		return nil, fmt.Errorf("select rows: %w", err)
	}

	defer rows.Close()

	out := make(map[string][]string, core.PossibleAttributeCount)

	for rows.Next() {
		var (
			code  string
			value string
		)

		err = rows.Scan(&code, &value)
		if err != nil {
			return nil, fmt.Errorf("scan rows: %w", err)
		}

		out[code] = append(out[code], value)
	}

	return out, nil
}

func (d *Database) UpdateAttributes(ctx context.Context, bookID uuid.UUID, attributes map[string][]string) error {
	tx, err := d.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}

	defer func() {
		err := tx.Rollback(ctx)
		if err != nil && !errors.Is(err, sql.ErrTxDone) && !errors.Is(err, pgx.ErrTxClosed) {
			d.logger.ErrorContext(
				ctx, "rollback UpdateAttributes tx",
				slog.Any("err", err),
			)
		}
	}()

	_, err = tx.Exec(ctx, `DELETE FROM book_attributes WHERE book_id = $1;`, bookID)
	if err != nil {
		return fmt.Errorf("delete old attributes: %w", err)
	}

	builder := squirrel.Insert("book_attributes").
		PlaceholderFormat(squirrel.Dollar).
		Columns(
			"book_id",
			"attr",
			"value",
		)

	for code, values := range attributes {
		for _, value := range values {
			builder = builder.Values(bookID, code, value)
		}
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	_, err = tx.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}

	return nil
}

func (d *Database) DeleteBookAttributes(ctx context.Context, bookID uuid.UUID) error {
	_, err := d.pool.Exec(ctx, `DELETE FROM book_attributes WHERE book_id = $1;`, bookID)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

func (d *Database) AttributesCount(ctx context.Context) ([]core.AttributeVariant, error) {
	rows, err := d.pool.Query(ctx, `SELECT COUNT(*), attr, value FROM book_attributes GROUP BY attr, value;`)
	if err != nil {
		return nil, fmt.Errorf("get attributes count: %w", err)
	}

	defer rows.Close()

	result := make([]core.AttributeVariant, 0, 100) // Берем изначальный запас емкости побольше

	for rows.Next() {
		var (
			count int
			code  string
			value string
		)

		err := rows.Scan(&count, &code, &value)
		if err != nil {
			return nil, fmt.Errorf("get attributes count: scan row: %w", err)
		}

		result = append(result, core.AttributeVariant{
			Code:  code,
			Value: value,
			Count: count,
		})
	}

	return result, nil
}

func (d *Database) Attributes(ctx context.Context) ([]core.Attribute, error) {
	builder := squirrel.Select(model.AttributeColumns()...).
		From("attributes").
		PlaceholderFormat(squirrel.Dollar).
		OrderBy("\"order\"")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	rows, err := d.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	defer rows.Close()

	result := make([]core.Attribute, 0, core.PossibleAttributeCount)

	for rows.Next() {
		attribute := core.Attribute{}

		err = rows.Scan(model.AttributeScanner(&attribute))
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		result = append(result, attribute)
	}

	return result, nil
}
