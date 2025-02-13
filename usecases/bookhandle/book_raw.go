package bookhandle

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (uc *UseCase) BookRaw(ctx context.Context, bookID uuid.UUID) (core.BookContainer, error) {
	bookFull, err := uc.bookRequester.BookOriginFull(ctx, bookID)
	if err != nil {
		return core.BookContainer{}, fmt.Errorf("book requester: %w", err)
	}

	return bookFull, nil
}
