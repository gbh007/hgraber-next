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
	Create(ctx context.Context, fileID uuid.UUID, body io.Reader) error
}

type storage interface {
	GetBooks(ctx context.Context, filter entities.BookFilter) ([]entities.BookFull, error)

	NewBook(ctx context.Context, book entities.Book) error
	UpdateBookPages(ctx context.Context, id uuid.UUID, pages []entities.Page) error
	SetLabel(ctx context.Context, label entities.BookLabel) error
	UpdateAttribute(ctx context.Context, id uuid.UUID, attrCode string, values []string) error
	NewFile(ctx context.Context, file entities.File) error

	DeleteBook(ctx context.Context, id uuid.UUID) error
}

type agentSystem interface {
	ExportArchive(ctx context.Context, agentID uuid.UUID, bookID uuid.UUID, bookName string, body io.Reader) error
}

type tmpStorage interface {
	AddToExport(books []entities.BookFullWithAgent)
	ExportList() []entities.BookFullWithAgent
}

type logger interface {
	Logger(ctx context.Context) *slog.Logger
}

type UseCase struct {
	logger logger

	storage     storage
	fileStorage fileStorage
	agentSystem agentSystem
	tmpStorage  tmpStorage
}

func New(
	logger logger,
	storage storage,
	fileStorage fileStorage,
	agentSystem agentSystem,
	tmpStorage tmpStorage,
) *UseCase {
	return &UseCase{
		logger:      logger,
		storage:     storage,
		fileStorage: fileStorage,
		agentSystem: agentSystem,
		tmpStorage:  tmpStorage,
	}
}
