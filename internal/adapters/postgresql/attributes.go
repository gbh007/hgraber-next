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

	"hgnext/internal/adapters/postgresql/internal/model"
)

func (d *Database) bookAttributes(ctx context.Context, bookID uuid.UUID) ([]*model.BookAttribute, error) {
	raw := make([]*model.BookAttribute, 0)

	err := d.db.SelectContext(ctx, &raw, `SELECT * FROM book_attributes WHERE book_id = $1;`, bookID.String())
	if err != nil {
		return nil, err
	}

	return raw, nil
}

func (d *Database) BookAttributes(ctx context.Context, bookID uuid.UUID) (map[string][]string, error) {
	attributes, err := d.bookAttributes(ctx, bookID)
	if err != nil {
		return nil, fmt.Errorf("get attributes :%w", err)
	}

	out := make(map[string][]string, 7)

	for _, attribute := range attributes {
		out[attribute.Attr] = append(out[attribute.Attr], attribute.Value)
	}

	return out, nil
}

func (d *Database) BookOriginAttributes(ctx context.Context, bookID uuid.UUID) (map[string][]string, error) {
	rows, err := d.pool.Query(ctx, `SELECT attr, values FROM book_origin_attributes WHERE book_id = $1;`, bookID.String())
	if err != nil {
		return nil, fmt.Errorf("select rows: %w", err)
	}

	defer rows.Close()

	out := make(map[string][]string, 7)

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

func (d *Database) UpdateAttribute(ctx context.Context, id uuid.UUID, attrCode string, values []string) error {
	tx, err := d.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `DELETE FROM book_attributes WHERE book_id = $1 AND attr = $2;`, id.String(), attrCode)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			d.logger.ErrorContext(ctx, rollbackErr.Error())
		}

		return err
	}

	for _, v := range values {
		_, err = tx.ExecContext(
			ctx,
			`INSERT INTO book_attributes (book_id, attr, value) VALUES($1, $2, $3);`,
			id.String(), attrCode, v,
		)
		if err != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				d.logger.ErrorContext(ctx, rollbackErr.Error())
			}

			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) UpdateAttributes(ctx context.Context, bookID uuid.UUID, attributes map[string][]string) error {
	tx, err := d.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}

	defer func() {
		err := tx.Rollback()
		if err != nil && !errors.Is(err, sql.ErrTxDone) {
			d.logger.ErrorContext(
				ctx, "rollback UpdateAttributes tx",
				slog.Any("err", err),
			)
		}
	}()

	_, err = tx.ExecContext(ctx, `DELETE FROM book_attributes WHERE book_id = $1;`, bookID.String())
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
			builder = builder.Values(bookID.String(), code, value)
		}
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("build query: %w", err)
	}

	d.squirrelDebugLog(ctx, query, args)

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}

	return nil
}

func (d *Database) UpdateOriginAttributes(ctx context.Context, bookID uuid.UUID, attributes map[string][]string) error {
	tx, err := d.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}

	defer func() {
		err := tx.Rollback(ctx)
		if err != nil && !errors.Is(err, sql.ErrTxDone) {
			d.logger.ErrorContext(
				ctx, "rollback UpdateAttributes tx",
				slog.Any("err", err),
			)
		}
	}()

	_, err = tx.Exec(ctx, `DELETE FROM book_origin_attributes WHERE book_id = $1;`, bookID.String())
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
		builder = builder.Values(bookID.String(), code, values)
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
