package cleanupusecase

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/domain/fsmodel"
)

type storage interface {
	DetachedFiles(ctx context.Context) ([]core.File, error)
	DeleteFile(ctx context.Context, id uuid.UUID) error
	FileIDsByFilter(ctx context.Context, filter fsmodel.FileFilter) ([]uuid.UUID, error)

	TruncateDeletedPages(ctx context.Context) error

	BookIDsWithDeletedRebuilds(ctx context.Context) ([]uuid.UUID, error)
	DeleteBooks(ctx context.Context, ids []uuid.UUID) error

	BookPagesCountByHash(ctx context.Context, hash core.FileHash) (int64, error)
	SetDeadHashes(ctx context.Context, hashes []core.DeadHash) error
	DeletedPagesHashes(ctx context.Context) ([]core.FileHash, error)
	RemoveDeletedPagesByHashes(ctx context.Context, hashes []core.FileHash) error
}

type fileStorage interface {
	Delete(ctx context.Context, fileID uuid.UUID, fsID *uuid.UUID) error
	State(ctx context.Context, includeFileIDs, includeFileSizes bool, fsID uuid.UUID) (fsmodel.FSState, error)
}

type UseCase struct {
	logger *slog.Logger
	tracer trace.Tracer

	storage     storage
	fileStorage fileStorage
}

func New(
	logger *slog.Logger,
	tracer trace.Tracer,
	storage storage,
	fileStorage fileStorage,
) *UseCase {
	return &UseCase{
		logger:      logger,
		tracer:      tracer,
		storage:     storage,
		fileStorage: fileStorage,
	}
}
