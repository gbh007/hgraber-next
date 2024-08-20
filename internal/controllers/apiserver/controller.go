package apiserver

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/url"
	"time"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"

	"hgnext/internal/entities"
	"hgnext/open_api/serverAPI"
)

type parseUseCases interface {
	NewBooks(ctx context.Context, urls []url.URL) (entities.FirstHandleMultipleResult, error)

	BooksExists(ctx context.Context, urls []url.URL) ([]entities.AgentBookCheckResult, error)
	PagesExists(ctx context.Context, urls []entities.AgentPageURL) ([]entities.AgentPageCheckResult, error)
	BookByURL(ctx context.Context, u url.URL) (entities.BookFull, error)
	PageBodyByURL(ctx context.Context, u url.URL) (io.Reader, error)
}

type webAPIUseCases interface {
	SystemInfo(ctx context.Context) (entities.SystemSizeInfoWithMonitor, error)

	File(ctx context.Context, fileID uuid.UUID) (io.Reader, error)
	PageBody(ctx context.Context, bookID uuid.UUID, pageNumber int) (io.Reader, error)

	Book(ctx context.Context, bookID uuid.UUID) (entities.BookToWeb, error)
	BookRaw(ctx context.Context, bookID uuid.UUID) (entities.BookFull, error)
	BookList(ctx context.Context, filter entities.BookFilter) ([]entities.BookToWeb, []int, error)
}

type agentUseCases interface {
	NewAgent(ctx context.Context, agent entities.Agent) error
	DeleteAgent(ctx context.Context, id uuid.UUID) error
	Agents(ctx context.Context, canParse, canExport, includeStatus bool) ([]entities.AgentWithStatus, error)
}

type exportUseCases interface {
	Export(ctx context.Context, agentID uuid.UUID, from, to time.Time) error
	ExportBook(ctx context.Context, bookID uuid.UUID) (io.Reader, error)
	ImportArchive(ctx context.Context, body io.Reader) (uuid.UUID, error)
}

type deduplicateUseCases interface {
	DeduplicateFiles(ctx context.Context) (count int, size int64, err error)
}

type cleanupUseCases interface {
	RemoveDetachedFiles(ctx context.Context) (count int, size int64, err error)
	RemoveFilesInStoragesMismatch(ctx context.Context) (notInDBCount, notInFSCount int, err error)
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
	cleanupUseCases     cleanupUseCases

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
	cleanupUseCases cleanupUseCases,
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
		cleanupUseCases:            cleanupUseCases,
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

// FIXME: реализовать
func (c *Controller) APIBookDeletePost(ctx context.Context, req *serverAPI.APIBookDeletePostReq) (serverAPI.APIBookDeletePostRes, error) {
	return &serverAPI.APIBookDeletePostNoContent{}, nil
}

// FIXME: реализовать
func (c *Controller) APIBookVerifyPost(ctx context.Context, req *serverAPI.APIBookVerifyPostReq) (serverAPI.APIBookVerifyPostRes, error) {
	return &serverAPI.APIBookVerifyPostNoContent{}, nil
}
