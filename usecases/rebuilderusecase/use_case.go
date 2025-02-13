package rebuilderusecase

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/domain/core"
)

type storage interface {
	GetBook(ctx context.Context, bookID uuid.UUID) (core.Book, error)
	NewBook(ctx context.Context, book core.Book) error
	UpdateBook(ctx context.Context, book core.Book) error
	DeleteBook(ctx context.Context, id uuid.UUID) error
	UpdateBookDeletion(ctx context.Context, book core.Book) error

	ReplaceLabels(ctx context.Context, bookID uuid.UUID, labels []core.BookLabel) error
	SetLabels(ctx context.Context, labels []core.BookLabel) error
	DeleteBookLabels(ctx context.Context, bookID uuid.UUID) error

	BookOriginAttributes(ctx context.Context, bookID uuid.UUID) (map[string][]string, error)
	UpdateOriginAttributes(ctx context.Context, bookID uuid.UUID, attributes map[string][]string) error
	UpdateAttributes(ctx context.Context, bookID uuid.UUID, attributes map[string][]string) error
	DeleteBookOriginAttributes(ctx context.Context, bookID uuid.UUID) error
	DeleteBookAttributes(ctx context.Context, bookID uuid.UUID) error

	BookPages(ctx context.Context, bookID uuid.UUID) ([]core.Page, error)
	NewBookPages(ctx context.Context, pages []core.Page) error
	BookPagesWithHash(ctx context.Context, bookID uuid.UUID) ([]core.PageWithHash, error)

	DeadHashesByMD5Sums(ctx context.Context, md5Sums []string) ([]core.DeadHash, error)

	DeletedPages(ctx context.Context, bookID uuid.UUID) ([]core.PageWithHash, error)
	FilesByMD5Sums(ctx context.Context, md5Sums []string) ([]core.File, error)
	RemoveDeletedPages(ctx context.Context, bookID uuid.UUID, pageNumbers []int) error
	BookPagesWithHashByMD5Sums(ctx context.Context, md5Sums []string) ([]core.PageWithHash, error)

	SetDeadHashes(ctx context.Context, hashes []core.DeadHash) error
	MarkPageAsDeleted(ctx context.Context, bookID uuid.UUID, pageNumber int) error
	MarkBookAsDeleted(ctx context.Context, bookID uuid.UUID) error
	BookPagesCount(ctx context.Context, bookID uuid.UUID) (int, error)
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
