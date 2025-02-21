package model

import (
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/domain/core"
)

func AttributeColumns() []string {
	return []string{
		"code",
		"name",
		"plural_name",
		"\"order\"",
		"description",
	}
}

func AttributeScanner(attribute *core.Attribute) RowScanner {
	return func(rows pgx.Rows) error {
		description := sql.NullString{}

		err := rows.Scan(
			&attribute.Code,
			&attribute.Name,
			&attribute.PluralName,
			&attribute.Order,
			&description,
		)
		if err != nil {
			return fmt.Errorf("scan to model: %w", err)
		}

		attribute.Description = description.String

		return nil
	}
}
