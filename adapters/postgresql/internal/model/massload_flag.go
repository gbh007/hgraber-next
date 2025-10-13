package model

import (
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

var MassloadFlagTable MassloadFlag

type MassloadFlag struct{}

func (MassloadFlag) Name() string {
	return "massload_flags"
}

func (MassloadFlag) ColumnCode() string            { return "code" }
func (MassloadFlag) ColumnName() string            { return "name" }
func (MassloadFlag) ColumnDescription() string     { return "description" }
func (MassloadFlag) ColumnOrderWeight() string     { return "order_weight" }
func (MassloadFlag) ColumnTextColor() string       { return "text_color" }
func (MassloadFlag) ColumnBackgroundColor() string { return "background_color" }
func (MassloadFlag) ColumnCreatedAt() string       { return "created_at" }

func (mf MassloadFlag) Columns() []string {
	return []string{
		mf.ColumnCode(),
		mf.ColumnName(),
		mf.ColumnDescription(),
		mf.ColumnOrderWeight(),
		mf.ColumnTextColor(),
		mf.ColumnBackgroundColor(),
		mf.ColumnCreatedAt(),
	}
}

func (MassloadFlag) Scanner(flag *massloadmodel.Flag) RowScanner {
	return func(rows pgx.Rows) error {
		var description, textColor, backgroundColor sql.NullString

		err := rows.Scan(
			&flag.Code,
			&flag.Name,
			&description,
			&flag.OrderWeight,
			&textColor,
			&backgroundColor,
			&flag.CreatedAt,
		)
		if err != nil {
			return fmt.Errorf("scan to model: %w", err)
		}

		flag.Description = description.String
		flag.TextColor = textColor.String
		flag.BackgroundColor = backgroundColor.String

		return nil
	}
}
