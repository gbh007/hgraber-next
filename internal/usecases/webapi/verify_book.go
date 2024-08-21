package webapi

import (
	"context"

	"github.com/google/uuid"
)

func (uc *UseCase) VerifyBook(ctx context.Context, bookID uuid.UUID) error {
	return uc.storage.VerifyBook(ctx, bookID)
}
