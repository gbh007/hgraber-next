package model

import (
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/domain/parsing"
)

func URLMirrorColumns() []string {
	return []string{
		"id",
		"name",
		"prefixes",
		"description",
	}
}

func URLMirrorScanner(mirror *parsing.URLMirror) RowScanner {
	return func(rows pgx.Rows) error {
		var description sql.NullString

		err := rows.Scan(
			&mirror.ID,
			&mirror.Name,
			&mirror.Prefixes,
			&description,
		)
		if err != nil {
			return fmt.Errorf("scan to model: %w", err)
		}

		mirror.Description = description.String

		return nil
	}
}
