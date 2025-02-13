package bff

import (
	"context"
	"log/slog"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/bff"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/domain/fsmodel"
)

type storage interface {
	BookIDs(ctx context.Context, filter core.BookFilter) ([]uuid.UUID, error)
	GetBook(ctx context.Context, bookID uuid.UUID) (core.Book, error)
	BookAttributes(ctx context.Context, bookID uuid.UUID) (map[string][]string, error)
	BookOriginAttributes(ctx context.Context, bookID uuid.UUID) (map[string][]string, error)
	BookPages(ctx context.Context, bookID uuid.UUID) ([]core.Page, error)
	Labels(ctx context.Context, bookID uuid.UUID) ([]core.BookLabel, error)

	BookIDsByMD5(ctx context.Context, md5sums []string) ([]uuid.UUID, error)
	BookPagesWithHash(ctx context.Context, bookID uuid.UUID) ([]core.PageWithHash, error)
	BookPageWithHash(ctx context.Context, bookID uuid.UUID, pageNumber int) (core.PageWithHash, error)
	BookPagesWithHashByMD5Sums(ctx context.Context, md5Sums []string) ([]core.PageWithHash, error)

	DeadHashesByMD5Sums(ctx context.Context, md5Sums []string) ([]core.DeadHash, error)

	BookCount(ctx context.Context, filter core.BookFilter) (int, error)

	Attributes(ctx context.Context) ([]core.Attribute, error)

	FileStorages(ctx context.Context) ([]fsmodel.FileStorageSystem, error)
}

type deduplicator interface {
	BookAttributesCompare(ctx context.Context, originID, targetID uuid.UUID, useOrigin bool) (bff.BookAttributesCompareResult, error)
	BookPagesCompare(ctx context.Context, originID, targetID uuid.UUID) (bff.BookPagesCompareResult, error)
}

type UseCase struct {
	logger *slog.Logger

	storage      storage
	deduplicator deduplicator
}

func New(
	logger *slog.Logger,
	storage storage,
	deduplicator deduplicator,
) *UseCase {
	return &UseCase{
		logger:       logger,
		storage:      storage,
		deduplicator: deduplicator,
	}
}
