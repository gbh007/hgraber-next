package cleanup

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/internal/entities"
)

type storage interface {
	DetachedFiles(ctx context.Context) ([]entities.File, error)
	DeleteFile(ctx context.Context, id uuid.UUID) error
	FileIDsByFS(ctx context.Context, fsID uuid.UUID) ([]uuid.UUID, error)

	TruncateDeletedPages(ctx context.Context) error

	BookIDsWithDeletedRebuilds(ctx context.Context) ([]uuid.UUID, error)
	DeleteBooks(ctx context.Context, ids []uuid.UUID) error
}

type fileStorage interface {
	Delete(ctx context.Context, fileID uuid.UUID, fsID *uuid.UUID) error
	State(ctx context.Context, includeFileIDs, includeFileSizes bool, fsID uuid.UUID) (entities.FSState, error)
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
