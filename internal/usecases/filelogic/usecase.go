package filelogic

import (
	"context"
	"io"
	"log/slog"

	"github.com/google/uuid"

	"hgnext/internal/entities"
)

type fileStorage interface {
	Get(ctx context.Context, fileID uuid.UUID) (io.Reader, error)
}

type storage interface {
	GetUnHashedFiles(ctx context.Context) ([]entities.File, error)
	UpdateFileHash(ctx context.Context, id uuid.UUID, md5Sum, sha256Sum string, size int64) error
}

type logger interface {
	Logger(ctx context.Context) *slog.Logger
}

type UseCase struct {
	logger logger

	storage     storage
	fileStorage fileStorage
}

func New(
	logger logger,
	storage storage,
	fileStorage fileStorage,
) *UseCase {
	return &UseCase{
		logger:      logger,
		storage:     storage,
		fileStorage: fileStorage,
	}
}
