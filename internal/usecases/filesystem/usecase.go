package filesystem

import (
	"context"
	"io"
	"log/slog"

	"github.com/google/uuid"

	"hgnext/internal/entities"
)

type storage interface {
	FileStorage(ctx context.Context, id uuid.UUID) (entities.FileStorageSystem, error)
	NewFileStorage(ctx context.Context, fs entities.FileStorageSystem) error
	UpdateFileStorage(ctx context.Context, fs entities.FileStorageSystem) error
	DeleteFileStorage(ctx context.Context, id uuid.UUID) error
	FSFilesInfo(ctx context.Context, fsID uuid.UUID, onlyInvalidData, onlyDetached bool) (entities.FSFilesInfo, error)

	File(ctx context.Context, id uuid.UUID) (entities.File, error)
	UpdateFileFS(ctx context.Context, fileID uuid.UUID, fsID uuid.UUID) error
	FileIDsByFS(ctx context.Context, fsID uuid.UUID) ([]uuid.UUID, error)
	UpdateFileInvalidData(ctx context.Context, fileID uuid.UUID, invalidData bool) error

	FileIDsByFilter(ctx context.Context, filter entities.FileFilter) ([]uuid.UUID, error)
}

type fileStorage interface {
	FSList(ctx context.Context) ([]entities.FSWithStatus, error)
	FSChange(ctx context.Context, fsID uuid.UUID, deleted bool) error
	Get(ctx context.Context, fileID uuid.UUID, fsID *uuid.UUID) (io.Reader, error)
	Create(ctx context.Context, fileID uuid.UUID, body io.Reader, fsID uuid.UUID) error
	Delete(ctx context.Context, fileID uuid.UUID, fsID *uuid.UUID) error
}

type tmpStorage interface {
	AddToValidate(ids []uuid.UUID)
	ValidateList() []uuid.UUID

	AddToFileTransfer(transfers []entities.FileTransfer)
	FileTransferList() []entities.FileTransfer
}

type UseCase struct {
	logger      *slog.Logger
	storage     storage
	fileStorage fileStorage
	tmpStorage  tmpStorage
}

func New(
	logger *slog.Logger,
	storage storage,
	fileStorage fileStorage,
	tmpStorage tmpStorage,
) *UseCase {
	return &UseCase{
		logger:      logger,
		storage:     storage,
		fileStorage: fileStorage,
		tmpStorage:  tmpStorage,
	}
}
