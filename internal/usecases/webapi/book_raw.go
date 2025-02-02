package webapi

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/internal/entities"
)

func (uc *UseCase) BookRaw(ctx context.Context, bookID uuid.UUID) (entities.BookContainer, error) {
	bookFull, err := uc.bookRequester.BookOriginFull(ctx, bookID)
	if err != nil {
		return entities.BookContainer{}, fmt.Errorf("book requester: %w", err)
	}

	return bookFull, nil
}
