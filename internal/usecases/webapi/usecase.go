package webapi

import (
	"context"
	"io"
	"log/slog"

	"github.com/google/uuid"

	"hgnext/internal/entities"
)

type storage interface {
	SystemSize(ctx context.Context) (entities.SystemSizeInfo, error)
	BookCount(ctx context.Context, filter entities.BookFilter) (int, error)
	GetPage(ctx context.Context, id uuid.UUID, pageNumber int) (entities.Page, error)
}

type bookRequester interface {
	Books(ctx context.Context, filter entities.BookFilter) ([]entities.BookFull, error)
	Book(ctx context.Context, bookID uuid.UUID) (entities.Book, error)
	BookFull(ctx context.Context, bookID uuid.UUID) (entities.BookFull, error)
	BookOriginFull(ctx context.Context, bookID uuid.UUID) (entities.BookFull, error)
}

type workerManager interface {
	Info() []entities.SystemWorkerStat
}

type fileStorage interface {
	Get(ctx context.Context, fileID uuid.UUID) (io.Reader, error)
}

type UseCase struct {
	logger *slog.Logger

	workerManager workerManager
	storage       storage
	fileStorage   fileStorage
	bookRequester bookRequester
}

func New(
	logger *slog.Logger,
	workerManager workerManager,
	storage storage,
	fileStorage fileStorage,
	bookRequester bookRequester,
) *UseCase {
	return &UseCase{
		logger:        logger,
		workerManager: workerManager,
		storage:       storage,
		fileStorage:   fileStorage,
		bookRequester: bookRequester,
	}
}
