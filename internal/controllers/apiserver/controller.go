package apiserver

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/url"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"

	"hgnext/internal/entities"
	"hgnext/open_api/serverAPI"
)

type parseUseCases interface {
	NewBooks(ctx context.Context, urls []url.URL, autoVerify bool) (entities.FirstHandleMultipleResult, error)

	BooksExists(ctx context.Context, urls []url.URL) ([]entities.AgentBookCheckResult, error)
	PagesExists(ctx context.Context, urls []entities.AgentPageURL) ([]entities.AgentPageCheckResult, error)
	BookByURL(ctx context.Context, u url.URL) (entities.BookFull, error)
	PageBodyByURL(ctx context.Context, u url.URL) (io.Reader, error)

	NewBooksMulti(ctx context.Context, urls []url.URL, autoVerify bool) (entities.MultiHandleMultipleResult, error)
}

type webAPIUseCases interface {
	SystemInfo(ctx context.Context) (entities.SystemSizeInfoWithMonitor, error)

	File(ctx context.Context, fileID uuid.UUID) (io.Reader, error)
	PageBody(ctx context.Context, bookID uuid.UUID, pageNumber int) (io.Reader, error)

	Book(ctx context.Context, bookID uuid.UUID) (entities.BookToWeb, error)
	BookRaw(ctx context.Context, bookID uuid.UUID) (entities.BookFull, error)
	BookList(ctx context.Context, filter entities.BookFilter) (entities.BookListToWeb, error)

	VerifyBook(ctx context.Context, bookID uuid.UUID) error
	DeleteBook(ctx context.Context, bookID uuid.UUID) error

	SetWorkerConfig(ctx context.Context, counts map[string]int)

	AttributesCount(ctx context.Context) ([]entities.AttributeVariant, error)

	SetLabel(ctx context.Context, label entities.BookLabel) error
	DeleteLabel(ctx context.Context, label entities.BookLabel) error
	Labels(ctx context.Context, bookID uuid.UUID) ([]entities.BookLabel, error)
	CreateLabelPreset(ctx context.Context, preset entities.BookLabelPreset) error
	UpdateLabelPreset(ctx context.Context, preset entities.BookLabelPreset) error
	DeleteLabelPreset(ctx context.Context, name string) error
	LabelPresets(ctx context.Context) ([]entities.BookLabelPreset, error)
	LabelPreset(ctx context.Context, name string) (entities.BookLabelPreset, error)
}

type agentUseCases interface {
	NewAgent(ctx context.Context, agent entities.Agent) error
	DeleteAgent(ctx context.Context, id uuid.UUID) error
	Agents(ctx context.Context, filter entities.AgentFilter, includeStatus bool) ([]entities.AgentWithStatus, error)
}

type exportUseCases interface {
	Export(ctx context.Context, agentID uuid.UUID, filter entities.BookFilter, deleteAfter bool) error
	ExportBook(ctx context.Context, bookID uuid.UUID) (io.Reader, entities.BookFull, error)
	ImportArchive(ctx context.Context, body io.Reader, deduplicate bool, autoVerify bool) (uuid.UUID, error)
}

type deduplicateUseCases interface {
	ArchiveEntryPercentage(ctx context.Context, archiveBody io.Reader) ([]entities.DeduplicateArchiveResult, error)
	BookByPageEntryPercentage(ctx context.Context, originBookID uuid.UUID) ([]entities.DeduplicateBookResult, error)
	BookPagesCompare(ctx context.Context, originID, targetID uuid.UUID) (entities.BookPagesCompareResult, error)
	UniquePages(ctx context.Context, originBookID uuid.UUID) ([]entities.Page, error)
	BooksByPage(ctx context.Context, bookID uuid.UUID, pageNumber int) ([]entities.BookWithPreviewPage, error)
}

type taskUseCases interface {
	RunTask(ctx context.Context, code entities.TaskCode) error
	TaskResults(ctx context.Context) ([]*entities.TaskResult, error)
}

type config interface {
	GetAddr() string
	GetExternalAddr() string
	GetStaticDir() string
	GetToken() string
}

type Controller struct {
	logger    *slog.Logger
	tracer    trace.Tracer
	debug     bool
	staticDir string

	parseUseCases       parseUseCases
	webAPIUseCases      webAPIUseCases
	agentUseCases       agentUseCases
	exportUseCases      exportUseCases
	deduplicateUseCases deduplicateUseCases
	taskUseCases        taskUseCases

	ogenServer *serverAPI.Server

	serverAddr string

	externalServerScheme       string
	externalServerHostWithPort string
	token                      string
}

func New(
	logger *slog.Logger,
	tracer trace.Tracer,
	config config,
	parseUseCases parseUseCases,
	webAPIUseCases webAPIUseCases,
	agentUseCases agentUseCases,
	exportUseCases exportUseCases,
	deduplicateUseCases deduplicateUseCases,
	taskUseCases taskUseCases,
	debug bool,
) (*Controller, error) {
	u, err := url.Parse(config.GetExternalAddr())
	if err != nil {
		return nil, fmt.Errorf("parse external server addr: %w", err)
	}

	c := &Controller{
		logger:                     logger,
		tracer:                     tracer,
		serverAddr:                 config.GetAddr(),
		externalServerScheme:       u.Scheme,
		externalServerHostWithPort: u.Host,
		parseUseCases:              parseUseCases,
		webAPIUseCases:             webAPIUseCases,
		agentUseCases:              agentUseCases,
		exportUseCases:             exportUseCases,
		deduplicateUseCases:        deduplicateUseCases,
		taskUseCases:               taskUseCases,
		debug:                      debug,
		staticDir:                  config.GetStaticDir(),
		token:                      config.GetToken(),
	}

	ogenServer, err := serverAPI.NewServer(
		c, c,
		// serverAPI.WithErrorHandler(), // FIXME: реализовать
		// serverAPI.WithMethodNotAllowed(), // FIXME: реализовать
		// serverAPI.WithNotFound(), // FIXME: реализовать
	)
	if err != nil {
		return nil, fmt.Errorf("create ogen server: %w", err)
	}

	c.ogenServer = ogenServer

	return c, nil
}
