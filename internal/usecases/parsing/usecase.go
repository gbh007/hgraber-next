package parsing

import (
	"context"
	"io"
	"log/slog"
	"net/url"
	"time"

	"github.com/google/uuid"

	"hgnext/internal/entities"
)

type storage interface {
	GetBookIDsByURL(ctx context.Context, url url.URL) ([]uuid.UUID, error)
	GetBook(ctx context.Context, bookID uuid.UUID) (entities.Book, error)

	NewBook(ctx context.Context, book entities.Book) error

	UpdateBook(ctx context.Context, book entities.Book) error
	UpdateAttributes(ctx context.Context, bookID uuid.UUID, attributes map[string][]string) error
	UpdateOriginAttributes(ctx context.Context, bookID uuid.UUID, attributes map[string][]string) error
	UpdateBookPages(ctx context.Context, id uuid.UUID, pages []entities.Page) error

	UpdatePageDownloaded(ctx context.Context, id uuid.UUID, pageNumber int, downloaded bool, fileID uuid.UUID) error
	NewFile(ctx context.Context, file entities.File) error

	NotDownloadedPages(ctx context.Context) ([]entities.PageForDownload, error)
	UnprocessedBooks(ctx context.Context) ([]entities.Book, error)

	Agents(ctx context.Context, filter entities.AgentFilter) ([]entities.Agent, error)

	PagesByURL(ctx context.Context, u url.URL) ([]entities.Page, error)
}

type agentSystem interface {
	BookParse(ctx context.Context, agentID uuid.UUID, url url.URL) (entities.AgentBookDetails, error)
	BooksCheck(ctx context.Context, agentID uuid.UUID, urls []url.URL) ([]entities.AgentBookCheckResult, error)
	PageLoad(ctx context.Context, agentID uuid.UUID, url entities.AgentPageURL) (io.Reader, error)
	PagesCheck(ctx context.Context, agentID uuid.UUID, urls []entities.AgentPageURL) ([]entities.AgentPageCheckResult, error)
	BooksCheckMultiple(ctx context.Context, agentID uuid.UUID, u url.URL) ([]entities.AgentBookCheckResult, error)
}

type fileStorage interface {
	Create(ctx context.Context, fileID uuid.UUID, body io.Reader) error
	Get(ctx context.Context, fileID uuid.UUID) (io.Reader, error)
}

type bookRequester interface {
	BookOriginFull(ctx context.Context, bookID uuid.UUID) (entities.BookFull, error)
}

type UseCase struct {
	logger *slog.Logger

	storage       storage
	agentSystem   agentSystem
	fileStorage   fileStorage
	bookRequester bookRequester

	parseBookTimeout time.Duration
}

func New(
	logger *slog.Logger,
	storage storage,
	agentSystem agentSystem,
	fileStorage fileStorage,
	bookRequester bookRequester,
	parseBookTimeout time.Duration,
) *UseCase {
	return &UseCase{
		logger:           logger,
		storage:          storage,
		agentSystem:      agentSystem,
		fileStorage:      fileStorage,
		parseBookTimeout: parseBookTimeout,
		bookRequester:    bookRequester,
	}
}
