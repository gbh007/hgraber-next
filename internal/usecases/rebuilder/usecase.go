package rebuilder

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"

	"hgnext/internal/entities"
)

type storage interface {
	GetBook(ctx context.Context, bookID uuid.UUID) (entities.Book, error)
	NewBook(ctx context.Context, book entities.Book) error
	UpdateBook(ctx context.Context, book entities.Book) error
	DeleteBook(ctx context.Context, id uuid.UUID) error

	ReplaceLabels(ctx context.Context, bookID uuid.UUID, labels []entities.BookLabel) error
	SetLabels(ctx context.Context, labels []entities.BookLabel) error

	BookOriginAttributes(ctx context.Context, bookID uuid.UUID) (map[string][]string, error)
	UpdateOriginAttributes(ctx context.Context, bookID uuid.UUID, attributes map[string][]string) error
	UpdateAttributes(ctx context.Context, bookID uuid.UUID, attributes map[string][]string) error

	BookPages(ctx context.Context, bookID uuid.UUID) ([]entities.Page, error)
	NewBookPages(ctx context.Context, pages []entities.Page) error
	BookPagesWithHash(ctx context.Context, bookID uuid.UUID) ([]entities.PageWithHash, error)

	DeadHashesByMD5Sums(ctx context.Context, md5Sums []string) ([]entities.DeadHash, error)
}

type UseCase struct {
	logger *slog.Logger
	tracer trace.Tracer

	storage storage
}

func New(
	logger *slog.Logger,
	tracer trace.Tracer,
	storage storage,
) *UseCase {
	return &UseCase{
		logger:  logger,
		tracer:  tracer,
		storage: storage,
	}
}
