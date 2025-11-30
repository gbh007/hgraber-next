package model

import (
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/domain/core"
)

var AttributeColorTable AttributeColor

type AttributeColor struct{}

func (AttributeColor) Name() string {
	return "attribute_colors"
}

func (AttributeColor) ColumnAttr() string            { return "attr" }
func (AttributeColor) ColumnValue() string           { return "value" }
func (AttributeColor) ColumnTextColor() string       { return "text_color" }
func (AttributeColor) ColumnBackgroundColor() string { return "background_color" }
func (AttributeColor) ColumnCreatedAt() string       { return "created_at" }

func (p AttributeColor) Columns() []string {
	return []string{
		p.ColumnAttr(),
		p.ColumnValue(),
		p.ColumnTextColor(),
		p.ColumnBackgroundColor(),
		p.ColumnCreatedAt(),
	}
}

func (AttributeColor) Scanner(color *core.AttributeColor) RowScanner {
	return func(rows pgx.Rows) error {
		err := rows.Scan(
			&color.Code,
			&color.Value,
			&color.TextColor,
			&color.BackgroundColor,
			&color.CreatedAt,
		)
		if err != nil {
			return fmt.Errorf("scan to model: %w", err)
		}

		return nil
	}
}
