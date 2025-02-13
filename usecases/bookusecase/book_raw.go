package bookusecase

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (uc *UseCase) BookRaw(ctx context.Context, bookID uuid.UUID) (core.BookContainer, error) {
	b, err := uc.storage.GetBook(ctx, bookID)
	if err != nil {
		return core.BookContainer{}, fmt.Errorf("get book: %w", err)
	}

	out := core.BookContainer{
		Book: b,
	}

	out.Attributes, err = uc.storage.BookOriginAttributes(ctx, bookID)
	if err != nil {
		return core.BookContainer{}, fmt.Errorf("get attributes: %w", err)
	}

	out.Pages, err = uc.storage.BookPages(ctx, bookID)
	if err != nil {
		return core.BookContainer{}, fmt.Errorf("get pages: %w", err)
	}

	out.Labels, err = uc.storage.Labels(ctx, bookID)
	if err != nil {
		return core.BookContainer{}, fmt.Errorf("get labels: %w", err)
	}

	return out, nil
}
