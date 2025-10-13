package model

import (
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/domain/core"
)

var BookLabelTable BookLabel

type BookLabel struct{}

func (BookLabel) Name() string {
	return "book_labels"
}

func (BookLabel) ColumnBookID() string     { return "book_id" }
func (BookLabel) ColumnPageNumber() string { return "page_number" }
func (BookLabel) ColumnName() string       { return "name" }
func (BookLabel) ColumnValue() string      { return "value" }
func (BookLabel) ColumnCreateAt() string   { return "create_at" }

func (bl BookLabel) Columns() []string {
	return []string{
		bl.ColumnBookID(),
		bl.ColumnPageNumber(),
		bl.ColumnName(),
		bl.ColumnValue(),
		bl.ColumnCreateAt(),
	}
}

func (BookLabel) Scanner(label *core.BookLabel) RowScanner {
	return func(rows pgx.Rows) error {
		err := rows.Scan(
			&label.BookID,
			&label.PageNumber,
			&label.Name,
			&label.Value,
			&label.CreateAt,
		)
		if err != nil {
			return fmt.Errorf("scan to model: %w", err)
		}

		return nil
	}
}
