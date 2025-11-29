package model

import (
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/domain/core"
)

var BookLabelTable = BookLabel{baseTable: baseTable{name: "book_labels"}}

type BookLabel struct {
	baseTable
}

func (bl BookLabel) WithPrefix(pf string) BookLabel {
	return BookLabel{
		baseTable: bl.withPrefix(pf),
	}
}

func (bl BookLabel) ColumnBookID() string     { return bl.prefix + "book_id" }
func (bl BookLabel) ColumnPageNumber() string { return bl.prefix + "page_number" }
func (bl BookLabel) ColumnName() string       { return bl.prefix + "name" }
func (bl BookLabel) ColumnValue() string      { return bl.prefix + "value" }
func (bl BookLabel) ColumnCreateAt() string   { return bl.prefix + "create_at" }

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
