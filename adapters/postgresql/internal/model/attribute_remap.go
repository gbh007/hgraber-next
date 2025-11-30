package model

import (
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/domain/core"
)

var AttributeRemapTable = AttributeRemap{baseTable: baseTable{name: "attribute_remaps"}}

type AttributeRemap struct {
	baseTable
}

func (ar AttributeRemap) WithPrefix(pf string) AttributeRemap {
	return AttributeRemap{
		baseTable: ar.withPrefix(pf),
	}
}

func (ar AttributeRemap) ColumnAttr() string      { return ar.prefix + "attr" }
func (ar AttributeRemap) ColumnValue() string     { return ar.prefix + "value" }
func (ar AttributeRemap) ColumnToAttr() string    { return ar.prefix + "to_attr" }
func (ar AttributeRemap) ColumnToValue() string   { return ar.prefix + "to_value" }
func (ar AttributeRemap) ColumnCreatedAt() string { return ar.prefix + "created_at" }
func (ar AttributeRemap) ColumnUpdatedAt() string { return ar.prefix + "updated_at" }

func (ar AttributeRemap) Columns() []string {
	return []string{
		ar.ColumnAttr(),
		ar.ColumnValue(),
		ar.ColumnToAttr(),
		ar.ColumnToValue(),
		ar.ColumnCreatedAt(),
		ar.ColumnUpdatedAt(),
	}
}

func (AttributeRemap) Scanner(attribute *core.AttributeRemap) RowScanner {
	return func(rows pgx.Rows) error {
		var (
			toAttr    sql.NullString
			toValue   sql.NullString
			updatedAt sql.NullTime
		)

		err := rows.Scan(
			&attribute.Code,
			&attribute.Value,
			&toAttr,
			&toValue,
			&attribute.CreatedAt,
			&updatedAt,
		)
		if err != nil {
			return fmt.Errorf("scan to model: %w", err)
		}

		attribute.ToCode = toAttr.String
		attribute.ToValue = toValue.String
		attribute.UpdateAt = updatedAt.Time

		return nil
	}
}
