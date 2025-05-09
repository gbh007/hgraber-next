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

	"github.com/gbh007/hgraber-next/domain/core"
)

func (d *Database) BookOriginAttributes(ctx context.Context, bookID uuid.UUID) (map[string][]string, error) {
	rows, err := d.pool.Query(ctx, `SELECT attr, values FROM book_origin_attributes WHERE book_id = $1;`, bookID)
	if err != nil {
		return nil, fmt.Errorf("select rows: %w", err)
	}

	defer rows.Close()

	out := make(map[string][]string, core.PossibleAttributeCount)

	for rows.Next() {
		var (
			code   string
			values []string
		)

		err = rows.Scan(&code, &values)
		if err != nil {
			return nil, fmt.Errorf("scan rows: %w", err)
		}

		out[code] = values
	}

	return out, nil
}

func (d *Database) UpdateOriginAttributes(ctx context.Context, bookID uuid.UUID, attributes map[string][]string) error {
	tx, err := d.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}

	defer func() {
		err := tx.Rollback(ctx)
		if err != nil && !errors.Is(err, sql.ErrTxDone) && !errors.Is(err, pgx.ErrTxClosed) {
			d.logger.ErrorContext(
				ctx, "rollback UpdateOriginAttributes tx",
				slog.Any("err", err),
			)
		}
	}()

	_, err = tx.Exec(ctx, `DELETE FROM book_origin_attributes WHERE book_id = $1;`, bookID)
	if err != nil {
		return fmt.Errorf("delete old attributes: %w", err)
	}

	builder := squirrel.Insert("book_origin_attributes").
		PlaceholderFormat(squirrel.Dollar).
		Columns(
			"book_id",
			"attr",
			"values",
		)

	for code, values := range attributes {
		builder = builder.Values(bookID, code, values)
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

func (d *Database) DeleteBookOriginAttributes(ctx context.Context, bookID uuid.UUID) error {
	_, err := d.pool.Exec(ctx, `DELETE FROM book_origin_attributes WHERE book_id = $1;`, bookID)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

func (d *Database) BookOriginAttributesCount(ctx context.Context) ([]core.AttributeVariant, error) {
	rows, err := d.pool.Query(ctx, `SELECT COUNT(*), attr, UNNEST(values) as value FROM book_origin_attributes GROUP BY attr, value;`)
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
