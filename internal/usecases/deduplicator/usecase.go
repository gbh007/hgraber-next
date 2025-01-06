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

	GetBook(ctx context.Context, bookID uuid.UUID) (entities.Book, error)
	BookIDsByMD5(ctx context.Context, md5sums []string) ([]uuid.UUID, error)

	GetPage(ctx context.Context, id uuid.UUID, pageNumber int) (entities.Page, error)
	BookPagesCount(ctx context.Context, bookID uuid.UUID) (int, error)
	BookPageWithHash(ctx context.Context, bookID uuid.UUID, pageNumber int) (entities.PageWithHash, error)
	BookPagesWithHash(ctx context.Context, bookID uuid.UUID) ([]entities.PageWithHash, error)
	BookPagesWithHashByHash(ctx context.Context, hash entities.FileHash) ([]entities.PageWithHash, error)
	BookPagesCountByHash(ctx context.Context, hash entities.FileHash) (int64, error)

	DeadHashesByMD5Sums(ctx context.Context, md5Sums []string) ([]entities.DeadHash, error)
	SetDeadHash(ctx context.Context, hash entities.DeadHash) error
	DeleteDeadHash(ctx context.Context, hash entities.DeadHash) error

	DeletedPagesHashes(ctx context.Context) ([]entities.FileHash, error)
	MarkPageAsDeleted(ctx context.Context, bookID uuid.UUID, pageNumber int) error
	MarkBookAsDeleted(ctx context.Context, bookID uuid.UUID) error

	BookAttributes(ctx context.Context, bookID uuid.UUID) (map[string][]string, error)
	BookOriginAttributes(ctx context.Context, bookID uuid.UUID) (map[string][]string, error)
	Attributes(ctx context.Context) ([]entities.Attribute, error)
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
