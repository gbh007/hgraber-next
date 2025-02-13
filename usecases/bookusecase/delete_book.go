package bookusecase

import (
	"context"

	"github.com/google/uuid"
)

func (uc *UseCase) DeleteBook(ctx context.Context, bookID uuid.UUID) error {
	return uc.storage.MarkBookAsDeleted(ctx, bookID)
}
