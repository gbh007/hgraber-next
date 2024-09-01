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

	BookIDsByMD5(ctx context.Context, md5sums []string) ([]uuid.UUID, error)
	BookPagesWithHash(ctx context.Context, bookID uuid.UUID) ([]entities.PageWithHash, error)
	GetBook(ctx context.Context, bookID uuid.UUID) (entities.Book, error)
}

type UseCase struct {
	logger *slog.Logger
	tracer trace.Tracer

	storage storage
}

func New(
	logger *slog.Logger,
	storage storage,
	tracer trace.Tracer,
) *UseCase {
	return &UseCase{
		logger:  logger,
		storage: storage,
		tracer:  tracer,
	}
}
