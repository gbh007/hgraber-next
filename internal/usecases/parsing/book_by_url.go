package parsing

import (
	"context"
	"fmt"
	"net/url"

	"hgnext/internal/entities"
)

func (uc *UseCase) BookByURL(ctx context.Context, u url.URL) (entities.BookFull, error) {
	ids, err := uc.storage.GetBookIDsByURL(ctx, u)
	if err != nil {
		return entities.BookFull{}, fmt.Errorf("get books by url: %w", err)
	}

	if len(ids) == 0 {
		return entities.BookFull{}, entities.BookNotFoundError
	}

	firstBook := entities.BookFull{}

	for i, id := range ids {
		book, err := uc.bookRequester.BookFull(ctx, id)
		if err != nil {
			return entities.BookFull{}, fmt.Errorf("get book by id (%s): %w", id.String(), err)
		}

		// Предпочитаем отдавать загруженную книгу
		if book.IsLoaded() {
			return book, nil
		}

		// Если нет загруженных книг, то вернем первую
		if i == 0 {
			firstBook = book
		}
	}

	return firstBook, nil
}
