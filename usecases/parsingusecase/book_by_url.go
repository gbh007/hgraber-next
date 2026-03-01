package parsingusecase

import (
	"context"
	"fmt"
	"net/url"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (uc *UseCase) BookByURL(ctx context.Context, u url.URL) (core.BookContainer, error) {
	ids, err := uc.storage.GetBookIDsByURL(ctx, u)
	if err != nil {
		return core.BookContainer{}, fmt.Errorf("get books by url: %w", err)
	}

	if len(ids) == 0 {
		return core.BookContainer{}, core.ErrBookNotFound
	}

	firstBook := core.BookContainer{}

	for _, id := range ids {
		book, err := uc.bookAdapter.BookRaw(ctx, id)
		if err != nil {
			return core.BookContainer{}, fmt.Errorf("get book by id (%s): %w", id.String(), err)
		}

		// Предпочитаем отдавать загруженную и не пересобранную книгу
		if book.IsLoaded() && !book.Book.IsRebuild {
			return book, nil
		}

		// Если нет загруженных книг, то вернем первую оригинальную
		if firstBook.Book.ID == uuid.Nil && !book.Book.IsRebuild {
			firstBook = book
		}
	}

	if firstBook.Book.ID == uuid.Nil {
		return core.BookContainer{}, core.ErrBookNotFound
	}

	return firstBook, nil
}
