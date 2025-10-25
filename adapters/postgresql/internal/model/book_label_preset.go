package model

import (
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/domain/core"
)

var BookLabelPresetTable BookLabelPreset

type BookLabelPreset struct{}

func (BookLabelPreset) Name() string {
	return "label_presets"
}

func (BookLabelPreset) ColumnName() string        { return "name" }
func (BookLabelPreset) ColumnDescription() string { return "description" }
func (BookLabelPreset) ColumnValues() string      { return "values" }
func (BookLabelPreset) ColumnCreatedAt() string   { return "created_at" }
func (BookLabelPreset) ColumnUpdatedAt() string   { return "updated_at" }

func (p BookLabelPreset) Columns() []string {
	return []string{
		p.ColumnName(),
		p.ColumnDescription(),
		p.ColumnValues(),
		p.ColumnCreatedAt(),
		p.ColumnUpdatedAt(),
	}
}

func (BookLabelPreset) Scanner(preset *core.BookLabelPreset) RowScanner {
	return func(rows pgx.Rows) error {
		var (
			updatedAt   sql.NullTime
			description sql.NullString
		)

		err := rows.Scan(
			&preset.Name,
			&description,
			&preset.Values,
			&preset.CreatedAt,
			&updatedAt,
		)
		if err != nil {
			return fmt.Errorf("scan to model: %w", err)
		}

		preset.Description = description.String
		preset.UpdatedAt = updatedAt.Time

		return nil
	}
}
