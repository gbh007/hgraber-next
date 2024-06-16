package webapi

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"hgnext/internal/entities"
)

func (uc *UseCase) Book(ctx context.Context, bookID uuid.UUID) (entities.BookToWeb, error) {
	bookFull, err := uc.storage.GetBookFull(ctx, bookID)
	if err != nil {
		return entities.BookToWeb{}, fmt.Errorf("storage: %w", err)
	}

	book := uc.bookConvert(bookFull)

	return book, nil
}
