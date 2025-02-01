package webapi

import (
	"context"
	"io"

	"github.com/google/uuid"
)

func (uc *UseCase) File(ctx context.Context, fileID uuid.UUID, fsID *uuid.UUID) (io.Reader, error) {
	return uc.fileStorage.Get(ctx, fileID, fsID)
}
