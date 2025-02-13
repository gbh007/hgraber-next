package bookusecase

import (
	"context"
	"log/slog"
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/core"
)

type storage interface {
	VerifyBook(ctx context.Context, bookID uuid.UUID, verified bool, verifiedAt time.Time) error
	MarkBookAsDeleted(ctx context.Context, bookID uuid.UUID) error

	GetBook(ctx context.Context, bookID uuid.UUID) (core.Book, error)
	BookOriginAttributes(ctx context.Context, bookID uuid.UUID) (map[string][]string, error)
	BookPages(ctx context.Context, bookID uuid.UUID) ([]core.Page, error)
	Labels(ctx context.Context, bookID uuid.UUID) ([]core.BookLabel, error)
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
