package parsing

import (
	"context"
	"io"
	"log/slog"
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/domain/core"
)

type storage interface {
	GetBookIDsByURL(ctx context.Context, url url.URL) ([]uuid.UUID, error)
	GetBook(ctx context.Context, bookID uuid.UUID) (core.Book, error)

	NewBook(ctx context.Context, book core.Book) error

	UpdateBook(ctx context.Context, book core.Book) error
	UpdateAttributes(ctx context.Context, bookID uuid.UUID, attributes map[string][]string) error
	UpdateOriginAttributes(ctx context.Context, bookID uuid.UUID, attributes map[string][]string) error
	UpdateBookPages(ctx context.Context, id uuid.UUID, pages []core.Page) error

	UpdatePageDownloaded(ctx context.Context, id uuid.UUID, pageNumber int, downloaded bool, fileID uuid.UUID) error
	NewFile(ctx context.Context, file core.File) error

	NotDownloadedPages(ctx context.Context) ([]core.PageForDownload, error)
	UnprocessedBooks(ctx context.Context) ([]core.Book, error)

	Agents(ctx context.Context, filter core.AgentFilter) ([]core.Agent, error)

	PagesByURL(ctx context.Context, u url.URL) ([]core.Page, error)
}

type agentSystem interface {
	BookParse(ctx context.Context, agentID uuid.UUID, url url.URL) (agentmodel.AgentBookDetails, error)
	BooksCheck(ctx context.Context, agentID uuid.UUID, urls []url.URL) ([]agentmodel.AgentBookCheckResult, error)
	PageLoad(ctx context.Context, agentID uuid.UUID, url agentmodel.AgentPageURL) (io.Reader, error)
	PagesCheck(ctx context.Context, agentID uuid.UUID, urls []agentmodel.AgentPageURL) ([]agentmodel.AgentPageCheckResult, error)
	BooksCheckMultiple(ctx context.Context, agentID uuid.UUID, u url.URL) ([]agentmodel.AgentBookCheckResult, error)
}

type fileStorage interface {
	FSIDForDownload(ctx context.Context) (uuid.UUID, error)
	Create(ctx context.Context, fileID uuid.UUID, body io.Reader, fsID uuid.UUID) error
	Get(ctx context.Context, fileID uuid.UUID, fsID *uuid.UUID) (io.Reader, error)
}

type bookRequester interface {
	BookOriginFull(ctx context.Context, bookID uuid.UUID) (core.BookContainer, error)
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
