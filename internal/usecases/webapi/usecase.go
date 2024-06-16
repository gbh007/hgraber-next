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
	GetBookFull(ctx context.Context, bookID uuid.UUID) (entities.BookFull, error)
	GetBooks(ctx context.Context, filter entities.BookFilter) ([]entities.BookFull, error)
	BookCount(ctx context.Context) (int, error)
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
}

func New(
	logger *slog.Logger,
	workerManager workerManager,
	storage storage,
	fileStorage fileStorage,
) *UseCase {
	return &UseCase{
		logger:        logger,
		workerManager: workerManager,
		storage:       storage,
		fileStorage:   fileStorage,
	}
}
