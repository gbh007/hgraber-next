package deduplicatorusecase

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/domain/core"
)

type storage interface {
	GetBook(ctx context.Context, bookID uuid.UUID) (core.Book, error)
	BookIDsByMD5(ctx context.Context, md5sums []string) ([]uuid.UUID, error)

	BookPagesCount(ctx context.Context, bookID uuid.UUID) (int, error)
	BookPageWithHash(ctx context.Context, bookID uuid.UUID, pageNumber int) (core.PageWithHash, error)
	BookPagesWithHash(ctx context.Context, bookID uuid.UUID) ([]core.PageWithHash, error)
	BookPagesWithHashByHash(ctx context.Context, hash core.FileHash) ([]core.PageWithHash, error)
	BookPagesCountByHash(ctx context.Context, hash core.FileHash) (int64, error)
	BookPagesWithHashByMD5Sums(ctx context.Context, md5Sums []string) ([]core.PageWithHash, error)

	DeadHashesByMD5Sums(ctx context.Context, md5Sums []string) ([]core.DeadHash, error)
	SetDeadHash(ctx context.Context, hash core.DeadHash) error
	SetDeadHashes(ctx context.Context, hashes []core.DeadHash) error
	DeleteDeadHash(ctx context.Context, hash core.DeadHash) error
	DeleteDeadHashes(ctx context.Context, hashes []core.DeadHash) error

	DeletedPagesHashes(ctx context.Context) ([]core.FileHash, error)
	MarkPageAsDeleted(ctx context.Context, bookID uuid.UUID, pageNumber int) error
	MarkBookAsDeleted(ctx context.Context, bookID uuid.UUID) error
	RemoveDeletedPagesByHashes(ctx context.Context, hashes []core.FileHash) error

	BookAttributes(ctx context.Context, bookID uuid.UUID) (map[string][]string, error)
	BookOriginAttributes(ctx context.Context, bookID uuid.UUID) (map[string][]string, error)
	Attributes(ctx context.Context) ([]core.Attribute, error)
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
