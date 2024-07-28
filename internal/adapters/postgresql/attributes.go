package postgresql

import (
	"context"
	"fmt"

	"github.com/google/uuid"

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
