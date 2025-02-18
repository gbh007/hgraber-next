package filesystemusecase

import (
	"context"
	"io"
	"log/slog"
	"net/url"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/domain/fsmodel"
)

type storage interface {
	FileStorage(ctx context.Context, id uuid.UUID) (fsmodel.FileStorageSystem, error)
	NewFileStorage(ctx context.Context, fs fsmodel.FileStorageSystem) error
	UpdateFileStorage(ctx context.Context, fs fsmodel.FileStorageSystem) error
	DeleteFileStorage(ctx context.Context, id uuid.UUID) error
	FSFilesInfo(ctx context.Context, fsID uuid.UUID, onlyInvalidData, onlyDetached bool) (core.SizeWithCount, error)

	File(ctx context.Context, id uuid.UUID) (core.File, error)
	UpdateFileFS(ctx context.Context, fileID uuid.UUID, fsID uuid.UUID) error
	UpdateFileInvalidData(ctx context.Context, fileID uuid.UUID, invalidData bool) error
	FileIDsByFilter(ctx context.Context, filter fsmodel.FileFilter) ([]uuid.UUID, error)

	GetPage(ctx context.Context, id uuid.UUID, pageNumber int) (core.Page, error)

	GetUnHashedFiles(ctx context.Context) ([]core.File, error)
	UpdateFileHash(ctx context.Context, id uuid.UUID, md5Sum, sha256Sum string, size int64) error
}

type fileStorage interface {
	FSList(ctx context.Context) ([]fsmodel.FSWithStatus, error)
	FSChange(ctx context.Context, fsID uuid.UUID, deleted bool) error
	Get(ctx context.Context, fileID uuid.UUID, fsID *uuid.UUID) (io.Reader, error)
	Create(ctx context.Context, fileID uuid.UUID, body io.Reader, fsID uuid.UUID) error
	Delete(ctx context.Context, fileID uuid.UUID, fsID *uuid.UUID) error
	HighwayFileURL(ctx context.Context, fileID uuid.UUID, ext string, fsID uuid.UUID) (url.URL, bool, error)
	State(ctx context.Context, includeFileIDs bool, includeFileSizes bool, fsID uuid.UUID) (fsmodel.FSState, error)
}

type tmpStorage interface {
	AddToValidate(ids []uuid.UUID)
	ValidateList() []uuid.UUID

	AddToFileTransfer(transfers []fsmodel.FileTransfer)
	FileTransferList() []fsmodel.FileTransfer
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

func (uc *UseCase) HighwayFileURL(ctx context.Context, fileID uuid.UUID, ext string, fsID uuid.UUID) (url.URL, bool, error) {
	return uc.fileStorage.HighwayFileURL(ctx, fileID, ext, fsID)
}
