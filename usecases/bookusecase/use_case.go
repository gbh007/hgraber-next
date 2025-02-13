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
}

type bookRequester interface { // FIXME: сделать этот пакет реализацией этого метода
	BookOriginFull(ctx context.Context, bookID uuid.UUID) (core.BookContainer, error)
}

type UseCase struct {
	logger *slog.Logger

	storage       storage
	bookRequester bookRequester
}

func New(
	logger *slog.Logger,
	storage storage,
	bookRequester bookRequester,
) *UseCase {
	return &UseCase{
		logger:        logger,
		storage:       storage,
		bookRequester: bookRequester,
	}
}
