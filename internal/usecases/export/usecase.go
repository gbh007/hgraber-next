package export

import (
	"context"
	"io"
	"log/slog"

	"github.com/google/uuid"

	"hgnext/internal/entities"
)

type fileStorage interface {
	Get(ctx context.Context, fileID uuid.UUID) (io.Reader, error)
}

type storage interface {
	GetBooks(ctx context.Context, filter entities.BookFilter) ([]entities.BookFull, error)
}

type agentSystem interface {
	ExportArchive(ctx context.Context, agentID uuid.UUID, bookID uuid.UUID, bookName string, body io.Reader) error
}

type UseCase struct {
	logger *slog.Logger

	storage     storage
	fileStorage fileStorage
	agentSystem agentSystem
}

func New(
	logger *slog.Logger,
	storage storage,
	fileStorage fileStorage,
	agentSystem agentSystem,
) *UseCase {
	return &UseCase{
		logger:      logger,
		storage:     storage,
		fileStorage: fileStorage,
		agentSystem: agentSystem,
	}
}
