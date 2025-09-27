package bookusecase

import (
	"context"

	"github.com/google/uuid"
)

func (uc *UseCase) DeleteBook(ctx context.Context, bookID uuid.UUID) error {
	return uc.storage.MarkBookAsDeleted(ctx, bookID) //nolint:wrapcheck // нет смысла в врапинге
}

func (uc *UseCase) DeletePage(ctx context.Context, bookID uuid.UUID, pageNumber int) error {
	return uc.storage.MarkPageAsDeleted(ctx, bookID, pageNumber) //nolint:wrapcheck // нет смысла в врапинге
}
