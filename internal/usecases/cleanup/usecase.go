package cleanup

import (
	"context"
	"log/slog"

	"github.com/google/uuid"

	"hgnext/internal/entities"
)

type storage interface {
	DetachedFiles(ctx context.Context) ([]entities.File, error)
	DeleteFile(ctx context.Context, id uuid.UUID) error
}

type fileStorage interface {
	Delete(ctx context.Context, fileID uuid.UUID) error
}

type UseCase struct {
	logger *slog.Logger

	storage     storage
	fileStorage fileStorage
}

func New(
	logger *slog.Logger,
	storage storage,
	fileStorage fileStorage,
) *UseCase {
	return &UseCase{
		logger:      logger,
		storage:     storage,
		fileStorage: fileStorage,
	}
}
