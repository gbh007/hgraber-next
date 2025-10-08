//revive:disable:file-length-limit
package model

import (
	"database/sql"
	"fmt"
	"net/url"

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

func MassloadExternalLinkColumns() []string {
	return []string{
		"url",
		"books_ahead",
		"new_books",
		"existing_books",
		"auto_check",
		"created_at",
		"updated_at",
	}
}

func MassloadExternalLinkScanner(link *massloadmodel.ExternalLink) RowScanner {
	return func(rows pgx.Rows) error {
		var (
			rawURL        string
			booksAhead    sql.NullInt64
			newBooks      sql.NullInt64
			existingBooks sql.NullInt64
			updatedAt     sql.NullTime
		)

		err := rows.Scan(
			&rawURL,
			&booksAhead,
			&newBooks,
			&existingBooks,
			&link.AutoCheck,
			&link.CreatedAt,
			&updatedAt,
		)
		if err != nil {
			return fmt.Errorf("scan to model: %w", err)
		}

		u, err := url.Parse(rawURL)
		if err != nil {
			return fmt.Errorf("parse url: %w", err)
		}

		link.URL = *u
		link.UpdatedAt = updatedAt.Time
		link.BooksAhead = NilInt64FromDB(booksAhead)
		link.NewBooks = NilInt64FromDB(newBooks)
		link.ExistingBooks = NilInt64FromDB(existingBooks)

		return nil
	}
}

func MassloadAttributeColumns() []string {
	return []string{
		"attr_code",
		"attr_value",
		"page_size",
		"file_size",
		"page_count",
		"file_count",
		"books_in_system",
		"created_at",
		"updated_at",
	}
}

func MassloadAttributeScanner(attr *massloadmodel.Attribute) RowScanner {
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

func MassloadFlagColumns() []string {
	return []string{
		"code",
		"name",
		"description",
		"order_weight",
		"text_color",
		"background_color",
		"created_at",
	}
}

func MassloadFlagScanner(flag *massloadmodel.Flag) RowScanner {
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
