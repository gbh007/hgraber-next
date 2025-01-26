package bff

import (
	"context"
	"log/slog"

	"github.com/google/uuid"

	"hgnext/internal/entities"
)

type storage interface {
	BookIDs(ctx context.Context, filter entities.BookFilter) ([]uuid.UUID, error)
	GetBook(ctx context.Context, bookID uuid.UUID) (entities.Book, error)
	BookAttributes(ctx context.Context, bookID uuid.UUID) (map[string][]string, error)
	BookOriginAttributes(ctx context.Context, bookID uuid.UUID) (map[string][]string, error)
	BookPages(ctx context.Context, bookID uuid.UUID) ([]entities.Page, error)
	Labels(ctx context.Context, bookID uuid.UUID) ([]entities.BookLabel, error)

	BookIDsByMD5(ctx context.Context, md5sums []string) ([]uuid.UUID, error)
	BookPagesWithHash(ctx context.Context, bookID uuid.UUID) ([]entities.PageWithHash, error)
	BookPagesWithHashByMD5Sums(ctx context.Context, md5Sums []string) ([]entities.PageWithHash, error)

	DeadHashesByMD5Sums(ctx context.Context, md5Sums []string) ([]entities.DeadHash, error)

	Attributes(ctx context.Context) ([]entities.Attribute, error)
}

type UseCase struct {
	logger *slog.Logger

	storage storage
}

func New(
	logger *slog.Logger,
	storage storage,
) *UseCase {
	return &UseCase{
		logger:  logger,
		storage: storage,
	}
}
