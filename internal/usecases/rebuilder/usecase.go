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
	UpdateBookDeletion(ctx context.Context, book entities.Book) error

	ReplaceLabels(ctx context.Context, bookID uuid.UUID, labels []entities.BookLabel) error
	SetLabels(ctx context.Context, labels []entities.BookLabel) error
	DeleteBookLabels(ctx context.Context, bookID uuid.UUID) error

	BookOriginAttributes(ctx context.Context, bookID uuid.UUID) (map[string][]string, error)
	UpdateOriginAttributes(ctx context.Context, bookID uuid.UUID, attributes map[string][]string) error
	UpdateAttributes(ctx context.Context, bookID uuid.UUID, attributes map[string][]string) error
	DeleteBookOriginAttributes(ctx context.Context, bookID uuid.UUID) error
	DeleteBookAttributes(ctx context.Context, bookID uuid.UUID) error

	BookPages(ctx context.Context, bookID uuid.UUID) ([]entities.Page, error)
	NewBookPages(ctx context.Context, pages []entities.Page) error
	BookPagesWithHash(ctx context.Context, bookID uuid.UUID) ([]entities.PageWithHash, error)

	DeadHashesByMD5Sums(ctx context.Context, md5Sums []string) ([]entities.DeadHash, error)

	DeletedPages(ctx context.Context, bookID uuid.UUID) ([]entities.PageWithHash, error)
	FilesByMD5Sums(ctx context.Context, md5Sums []string) ([]entities.File, error)
	RemoveDeletedPages(ctx context.Context, bookID uuid.UUID, pageNumbers []int) error
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
