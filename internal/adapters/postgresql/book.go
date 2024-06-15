package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/url"

	"github.com/google/uuid"

	"hgnext/internal/adapters/postgresql/internal/model"
	"hgnext/internal/entities"
)

func (d *Database) NewBook(ctx context.Context, book entities.Book) error {
	_, err := d.db.ExecContext(
		ctx,
		`INSERT INTO books (id, name, origin_url, page_count, attributes_parsed, create_at) VALUES($1, $2, $3, $4, $5, $6);`,
		book.ID.String(), model.StringToDB(book.Name), model.StringToDB(book.OriginURL), model.Int32ToDB(book.PageCount), book.AttributesParsed, book.CreateAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) UpdateBookName(ctx context.Context, id uuid.UUID, name string) error {
	res, err := d.db.ExecContext(ctx, `UPDATE books SET name = $1 WHERE id = $2;`, name, id.String())
	if err != nil {
		return err
	}

	if !d.isApply(ctx, res) {
		return entities.BookNotFoundError
	}

	return nil
}

func (d *Database) GetBookIDsByURL(ctx context.Context, url url.URL) ([]uuid.UUID, error) {
	var idsRaw []string

	err := d.db.SelectContext(ctx, &idsRaw, `SELECT id FROM books WHERE origin_url = $1;`, url.String())
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("%w - %s", entities.BookNotFoundError, url.String())
	}

	if err != nil {
		return nil, err
	}

	ids := make([]uuid.UUID, len(idsRaw))

	for i, idRaw := range idsRaw {
		ids[i], err = uuid.Parse(idRaw)
		if err != nil {
			return nil, err
		}
	}

	return ids, nil
}

// FIXME: отказаться от подобной функции
func (d *Database) bookIDs(ctx context.Context, filter entities.BookFilter) ([]uuid.UUID, error) {
	idsRaw := make([]string, 0)

	query := `SELECT id FROM books ORDER BY id ASC LIMIT $1 OFFSET $2;`
	if filter.NewFirst {
		query = `SELECT id FROM books ORDER BY id DESC LIMIT $1 OFFSET $2;`
	}

	err := d.db.SelectContext(ctx, &idsRaw, query, filter.Limit, filter.Offset)
	if err != nil {
		return nil, err
	}

	ids := make([]uuid.UUID, len(idsRaw))

	for i, idRaw := range idsRaw {
		ids[i], err = uuid.Parse(idRaw)
		if err != nil {
			return nil, err
		}
	}

	return ids, nil
}

func (d *Database) GetBook(ctx context.Context, bookID uuid.UUID) (entities.BookFull, error) {
	raw := new(model.Book)

	err := d.db.GetContext(ctx, raw, `SELECT * FROM books WHERE id = $1 LIMIT 1;`, bookID)
	if errors.Is(err, sql.ErrNoRows) {
		return entities.BookFull{}, fmt.Errorf("%w - %d", entities.BookNotFoundError, bookID)
	}

	if err != nil {
		return entities.BookFull{}, err
	}

	b, err := raw.ToEntity()
	if err != nil {
		return entities.BookFull{}, err
	}

	out := entities.BookFull{
		Book: b,
	}

	attributes, err := d.getBookAttr(ctx, bookID)
	if err != nil {
		return entities.BookFull{}, err
	}

	for _, attribute := range attributes {
		out.Attributes[attribute.Attr] = append(out.Attributes[attribute.Attr], attribute.Value)
	}

	pages, err := d.getBookPages(ctx, bookID)
	if err != nil {
		return entities.BookFull{}, err
	}

	for _, pageRaw := range pages {
		page, err := pageRaw.ToEntity()
		if err != nil {
			return entities.BookFull{}, err
		}

		out.Pages = append(out.Pages, page)
	}

	return out, nil
}

func (d *Database) GetBooks(ctx context.Context, filter entities.BookFilter) ([]entities.BookFull, error) {
	out := make([]entities.BookFull, 0)

	ids, err := d.bookIDs(ctx, filter)
	if err != nil {
		return nil, err
	}

	for _, id := range ids {
		book, err := d.GetBook(ctx, id)
		if err != nil {
			return nil, err
		}

		out = append(out, book)
	}

	return out, nil
}
