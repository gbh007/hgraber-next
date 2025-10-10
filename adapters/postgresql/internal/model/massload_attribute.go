package model

import (
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

var MassloadAttributeTable MassloadAttribute

type MassloadAttribute struct{}

func (MassloadAttribute) Name() string {
	return "massload_attributes"
}

func (MassloadAttribute) ColumnMassloadID() string    { return "massload_id" }
func (MassloadAttribute) ColumnAttrCode() string      { return "attr_code" }
func (MassloadAttribute) ColumnAttrValue() string     { return "attr_value" }
func (MassloadAttribute) ColumnPageSize() string      { return "page_size" }
func (MassloadAttribute) ColumnFileSize() string      { return "file_size" }
func (MassloadAttribute) ColumnPageCount() string     { return "page_count" }
func (MassloadAttribute) ColumnFileCount() string     { return "file_count" }
func (MassloadAttribute) ColumnBooksInSystem() string { return "books_in_system" }
func (MassloadAttribute) ColumnCreatedAt() string     { return "created_at" }
func (MassloadAttribute) ColumnUpdatedAt() string     { return "updated_at" }

func (ma MassloadAttribute) Columns() []string {
	return []string{
		// ma.ColumnMassloadID(), // Пока не используется в модели логики
		ma.ColumnAttrCode(),
		ma.ColumnAttrValue(),
		ma.ColumnPageSize(),
		ma.ColumnFileSize(),
		ma.ColumnPageCount(),
		ma.ColumnFileCount(),
		ma.ColumnBooksInSystem(),
		ma.ColumnCreatedAt(),
		ma.ColumnUpdatedAt(),
	}
}

func (MassloadAttribute) Scanner(attr *massloadmodel.Attribute) RowScanner {
	return func(rows pgx.Rows) error {
		var (
			pageSize     sql.NullInt64
			fileSize     sql.NullInt64
			pageCount    sql.NullInt64
			fileCount    sql.NullInt64
			bookInSystem sql.NullInt64
			updatedAt    sql.NullTime
		)

		err := rows.Scan(
			&attr.Code,
			&attr.Value,
			&pageSize,
			&fileSize,
			&pageCount,
			&fileCount,
			&bookInSystem,
			&attr.CreatedAt,
			&updatedAt,
		)
		if err != nil {
			return fmt.Errorf("scan to model: %w", err)
		}

		attr.UpdatedAt = updatedAt.Time
		attr.PageSize = NilInt64FromDB(pageSize)
		attr.FileSize = NilInt64FromDB(fileSize)
		attr.PageCount = NilInt64FromDB(pageCount)
		attr.FileCount = NilInt64FromDB(fileCount)
		attr.BookInSystem = NilInt64FromDB(bookInSystem)

		return nil
	}
}
