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

	attributesInfo, err := uc.storage.Attributes(ctx)
	if err != nil {
		return entities.BookToWeb{}, fmt.Errorf("storage: get attributes info: %w", err)
	}

	attributesInfoMap := convertAttributes(attributesInfo)
	book := uc.bookConvert(bookFull, attributesInfoMap)

	return book, nil
}

func (uc *UseCase) BookRaw(ctx context.Context, bookID uuid.UUID) (entities.BookContainer, error) {
	bookFull, err := uc.bookRequester.BookOriginFull(ctx, bookID)
	if err != nil {
		return entities.BookContainer{}, fmt.Errorf("book requester: %w", err)
	}

	return bookFull, nil
}
