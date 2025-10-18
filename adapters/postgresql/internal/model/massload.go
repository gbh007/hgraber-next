//revive:disable:file-length-limit
package model

import (
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

var MassloadTable Massload

type Massload struct{}

func (Massload) Name() string {
	return "massloads"
}

func (Massload) ColumnID() string            { return "id" }
func (Massload) ColumnName() string          { return "name" }
func (Massload) ColumnDescription() string   { return "description" }
func (Massload) ColumnFlags() string         { return "flags" }
func (Massload) ColumnPageSize() string      { return "page_size" }
func (Massload) ColumnFileSize() string      { return "file_size" }
func (Massload) ColumnPageCount() string     { return "page_count" }
func (Massload) ColumnFileCount() string     { return "file_count" }
func (Massload) ColumnBooksAhead() string    { return "books_ahead" }
func (Massload) ColumnNewBooks() string      { return "new_books" }
func (Massload) ColumnExistingBooks() string { return "existing_books" }
func (Massload) ColumnBooksInSystem() string { return "books_in_system" }
func (Massload) ColumnCreatedAt() string     { return "created_at" }
func (Massload) ColumnUpdatedAt() string     { return "updated_at" }

func (m Massload) Columns() []string {
	return []string{
		m.ColumnID(),
		m.ColumnName(),
		m.ColumnDescription(),
		m.ColumnFlags(),
		m.ColumnPageSize(),
		m.ColumnFileSize(),
		m.ColumnPageCount(),
		m.ColumnFileCount(),
		m.ColumnBooksAhead(),
		m.ColumnNewBooks(),
		m.ColumnExistingBooks(),
		m.ColumnBooksInSystem(),
		m.ColumnCreatedAt(),
		m.ColumnUpdatedAt(),
	}
}

func (Massload) Scanner(ml *massloadmodel.Massload) RowScanner {
	return func(rows pgx.Rows) error {
		var (
			description   sql.NullString
			pageSize      sql.NullInt64
			fileSize      sql.NullInt64
			pageCount     sql.NullInt64
			fileCount     sql.NullInt64
			booksAhead    sql.NullInt64
			newBooks      sql.NullInt64
			existingBooks sql.NullInt64
			bookInSystem  sql.NullInt64
			updatedAt     sql.NullTime
		)

		err := rows.Scan(
			&ml.ID,
			&ml.Name,
			&description,
			&ml.Flags,
			&pageSize,
			&fileSize,
			&pageCount,
			&fileCount,
			&booksAhead,
			&newBooks,
			&existingBooks,
			&bookInSystem,
			&ml.CreatedAt,
			&updatedAt,
		)
		if err != nil {
			return fmt.Errorf("scan to model: %w", err)
		}

		ml.UpdatedAt = updatedAt.Time
		ml.Description = description.String
		ml.PageSize = NilInt64FromDB(pageSize)
		ml.FileSize = NilInt64FromDB(fileSize)
		ml.PageCount = NilInt64FromDB(pageCount)
		ml.FileCount = NilInt64FromDB(fileCount)
		ml.BooksAhead = NilInt64FromDB(booksAhead)
		ml.NewBooks = NilInt64FromDB(newBooks)
		ml.ExistingBooks = NilInt64FromDB(existingBooks)
		ml.BookInSystem = NilInt64FromDB(bookInSystem)

		return nil
	}
}
