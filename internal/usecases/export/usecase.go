package export

import (
	"context"
	"io"
	"log/slog"
	"net/url"

	"github.com/google/uuid"

	"hgnext/internal/entities"
)

type fileStorage interface {
	Get(ctx context.Context, fileID uuid.UUID, fsID *uuid.UUID) (io.Reader, error)
	Create(ctx context.Context, fileID uuid.UUID, body io.Reader, fsID uuid.UUID) error
	FSIDForDownload(ctx context.Context) (uuid.UUID, error)
}

type storage interface {
	NewBook(ctx context.Context, book entities.Book) error
	UpdateBookPages(ctx context.Context, id uuid.UUID, pages []entities.Page) error
	SetLabel(ctx context.Context, label entities.BookLabel) error
	UpdateAttributes(ctx context.Context, bookID uuid.UUID, attributes map[string][]string) error
	UpdateOriginAttributes(ctx context.Context, bookID uuid.UUID, attributes map[string][]string) error
	NewFile(ctx context.Context, file entities.File) error

	DeleteBook(ctx context.Context, id uuid.UUID) error

	GetBookIDsByURL(ctx context.Context, url url.URL) ([]uuid.UUID, error)

	MarkBookAsDeleted(ctx context.Context, bookID uuid.UUID) error
}

type agentSystem interface {
	ExportArchive(ctx context.Context, agentID uuid.UUID, data entities.AgentExportData) error
}

type tmpStorage interface {
	AddToExport(books []entities.BookFullWithAgent)
	ExportList() []entities.BookFullWithAgent
}

type bookRequester interface {
	Books(ctx context.Context, filter entities.BookFilter) ([]entities.BookContainer, error)
	BookOriginFull(ctx context.Context, bookID uuid.UUID) (entities.BookContainer, error)
}

type UseCase struct {
	logger *slog.Logger

	storage       storage
	fileStorage   fileStorage
	agentSystem   agentSystem
	tmpStorage    tmpStorage
	bookRequester bookRequester
}

func New(
	logger *slog.Logger,
	storage storage,
	fileStorage fileStorage,
	agentSystem agentSystem,
	tmpStorage tmpStorage,
	bookRequester bookRequester,
) *UseCase {
	return &UseCase{
		logger:        logger,
		storage:       storage,
		fileStorage:   fileStorage,
		agentSystem:   agentSystem,
		tmpStorage:    tmpStorage,
		bookRequester: bookRequester,
	}
}
