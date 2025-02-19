package model

import (
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/domain/core"
)

func BookLabelColumns() []string {
	return []string{
		"book_id",
		"page_number",
		"name",
		"value",
		"create_at",
	}
}

func BookLabelScanner(label *core.BookLabel) RowScanner {
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
