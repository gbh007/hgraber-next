package model

import (
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/domain/core"
)

var AttributeTable = Attribute{baseTable: baseTable{name: "attributes"}}

type Attribute struct {
	baseTable
}

func (a Attribute) WithPrefix(pf string) Attribute {
	return Attribute{
		baseTable: a.withPrefix(pf),
	}
}

func (a Attribute) ColumnCode() string        { return a.prefix + "code" }
func (a Attribute) ColumnName() string        { return a.prefix + "name" }
func (a Attribute) ColumnPluralName() string  { return a.prefix + "plural_name" }
func (a Attribute) ColumnOrder() string       { return a.prefix + "\"order\"" }
func (a Attribute) ColumnDescription() string { return a.prefix + "description" }

func (a Attribute) Columns() []string {
	return []string{
		a.ColumnCode(),
		a.ColumnName(),
		a.ColumnPluralName(),
		a.ColumnOrder(),
		a.ColumnDescription(),
	}
}

func (Attribute) Scanner(attribute *core.Attribute) RowScanner {
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
