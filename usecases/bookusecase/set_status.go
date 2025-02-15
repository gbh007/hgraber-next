package bookusecase

import (
	"context"
	"time"

	"github.com/google/uuid"
)

func (uc *UseCase) VerifyBook(ctx context.Context, bookID uuid.UUID, verified bool) error {
	var verifiedAt time.Time

	if verified {
		verifiedAt = time.Now().UTC()
	}

	return uc.storage.VerifyBook(ctx, bookID, verified, verifiedAt)
}

func (uc *UseCase) SetBookRebuild(ctx context.Context, bookID uuid.UUID, value bool) error {
	return uc.storage.SetBookRebuild(ctx, bookID, value)
}
