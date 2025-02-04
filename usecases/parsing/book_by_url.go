package parsing

import (
	"context"
	"fmt"
	"net/url"

	"github.com/gbh007/hgraber-next/entities"
)

func (uc *UseCase) BookByURL(ctx context.Context, u url.URL) (entities.BookContainer, error) {
	ids, err := uc.storage.GetBookIDsByURL(ctx, u)
	if err != nil {
		return entities.BookContainer{}, fmt.Errorf("get books by url: %w", err)
	}

	if len(ids) == 0 {
		return entities.BookContainer{}, entities.BookNotFoundError
	}

	firstBook := entities.BookContainer{}

	for i, id := range ids {
		book, err := uc.bookRequester.BookOriginFull(ctx, id)
		if err != nil {
			return entities.BookContainer{}, fmt.Errorf("get book by id (%s): %w", id.String(), err)
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
