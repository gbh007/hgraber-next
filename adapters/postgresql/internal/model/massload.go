//revive:disable:file-length-limit
package model

import (
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

func MassloadColumns() []string {
	return []string{
		"id",
		"name",
		"description",
		"flags",
		"page_size",
		"file_size",
		"page_count",
		"file_count",
		"books_ahead",
		"new_books",
		"existing_books",
		"books_in_system",
		"created_at",
		"updated_at",
	}
}

func MassloadScanner(ml *massloadmodel.Massload) RowScanner {
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
