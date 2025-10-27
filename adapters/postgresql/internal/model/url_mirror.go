package model

import (
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/domain/parsing"
)

var URLMirrorTable URLMirror

type URLMirror struct{}

func (URLMirror) Name() string {
	return "url_mirrors"
}

func (URLMirror) ColumnID() string          { return "id" }
func (URLMirror) ColumnName() string        { return "name" }
func (URLMirror) ColumnPrefixes() string    { return "prefixes" }
func (URLMirror) ColumnDescription() string { return "description" }

func (um URLMirror) Columns() []string {
	return []string{
		um.ColumnID(),
		um.ColumnName(),
		um.ColumnPrefixes(),
		um.ColumnDescription(),
	}
}

func (URLMirror) Scanner(mirror *parsing.URLMirror) RowScanner {
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
