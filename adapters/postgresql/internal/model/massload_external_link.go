package model

import (
	"database/sql"
	"fmt"
	"net/url"

	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

var MassloadExternalLinkTable MassloadExternalLink

type MassloadExternalLink struct{}

func (MassloadExternalLink) Name() string {
	return "massload_external_links"
}

func (MassloadExternalLink) ColumnMassloadID() string    { return "massload_id" }
func (MassloadExternalLink) ColumnURL() string           { return "url" }
func (MassloadExternalLink) ColumnBooksAhead() string    { return "books_ahead" }
func (MassloadExternalLink) ColumnNewBooks() string      { return "new_books" }
func (MassloadExternalLink) ColumnExistingBooks() string { return "existing_books" }
func (MassloadExternalLink) ColumnAutoCheck() string     { return "auto_check" }
func (MassloadExternalLink) ColumnCreatedAt() string     { return "created_at" }
func (MassloadExternalLink) ColumnUpdatedAt() string     { return "updated_at" }

func (m MassloadExternalLink) Columns() []string {
	return []string{
		m.ColumnURL(),
		m.ColumnBooksAhead(),
		m.ColumnNewBooks(),
		m.ColumnExistingBooks(),
		m.ColumnAutoCheck(),
		m.ColumnCreatedAt(),
		m.ColumnUpdatedAt(),
	}
}

func (MassloadExternalLink) Scanner(link *massloadmodel.ExternalLink) RowScanner {
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
