package deduplicator

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"

	"hgnext/internal/entities"
)

type storage interface {
	DuplicatedFiles(ctx context.Context) ([]entities.File, error)
	ReplaceFile(ctx context.Context, oldFileID, newFileID uuid.UUID) error
}

type logger interface {
	Logger(ctx context.Context) *slog.Logger
}

type UseCase struct {
	logger logger
	tracer trace.Tracer

	storage storage
}

func New(
	logger logger,
	storage storage,
	tracer trace.Tracer,
) *UseCase {
	return &UseCase{
		logger:  logger,
		storage: storage,
		tracer:  tracer,
	}
}
