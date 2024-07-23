package cleanup

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"

	"hgnext/internal/entities"
)

type storage interface {
	DetachedFiles(ctx context.Context) ([]entities.File, error)
	DeleteFile(ctx context.Context, id uuid.UUID) error
	FileIDs(ctx context.Context) ([]uuid.UUID, error)
}

type fileStorage interface {
	Delete(ctx context.Context, fileID uuid.UUID) error
	IDs(ctx context.Context) ([]uuid.UUID, error)
}

type logger interface {
	Logger(ctx context.Context) *slog.Logger
}

type UseCase struct {
	logger logger
	tracer trace.Tracer

	storage     storage
	fileStorage fileStorage
}

func New(
	logger logger,
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
