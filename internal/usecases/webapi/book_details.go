package webapi

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"hgnext/internal/entities"
)

func (uc *UseCase) Book(ctx context.Context, bookID uuid.UUID) (entities.BookToWeb, error) {
	bookFull, err := uc.bookRequester.BookFull(ctx, bookID)
	if err != nil {
		return entities.BookToWeb{}, fmt.Errorf("book requester: %w", err)
	}

	book := uc.bookConvert(bookFull)

	return book, nil
}

func (uc *UseCase) BookRaw(ctx context.Context, bookID uuid.UUID) (entities.BookFull, error) {
	bookFull, err := uc.bookRequester.BookOriginFull(ctx, bookID)
	if err != nil {
		return entities.BookFull{}, fmt.Errorf("book requester: %w", err)
	}

	return bookFull, nil
}
