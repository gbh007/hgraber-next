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

	"hgnext/internal/controllers/apiserver/internal/server"
	"hgnext/internal/entities"
)

type parseUseCases interface {
	NewBooks(ctx context.Context, urls []url.URL) (entities.FirstHandleMultipleResult, error)
}

type webAPIUseCases interface {
	SystemInfo(ctx context.Context) (entities.SystemSizeInfoWithMonitor, error)

	File(ctx context.Context, fileID uuid.UUID) (io.Reader, error)

	Book(ctx context.Context, bookID uuid.UUID) (entities.BookToWeb, error)
	BookList(ctx context.Context, filter entities.BookFilter) ([]entities.BookToWeb, []int, error)
}

type agentUseCases interface {
	NewAgent(ctx context.Context, agent entities.Agent) error
	DeleteAgent(ctx context.Context, id uuid.UUID) error
	Agents(ctx context.Context, canParse, canExport, includeStatus bool) ([]entities.AgentWithStatus, error)
}

type exportUseCases interface {
	Export(ctx context.Context, agentID uuid.UUID, from, to time.Time) error
	ImportArchive(ctx context.Context, body io.Reader) (uuid.UUID, error)
}

type deduplicateUseCases interface {
	DeduplicateFiles(ctx context.Context) (count int, size int64, err error)
}

type cleanupUseCases interface {
	RemoveDetachedFiles(ctx context.Context) (count int, size int64, err error)
	RemoveFilesInStoragesMismatch(ctx context.Context) (notInDBCount, notInFSCount int, err error)
}

type logger interface {
	Logger(ctx context.Context) *slog.Logger
}

type Controller struct {
	logger    logger
	tracer    trace.Tracer
	debug     bool
	staticDir string

	parseUseCases       parseUseCases
	webAPIUseCases      webAPIUseCases
	agentUseCases       agentUseCases
	exportUseCases      exportUseCases
	deduplicateUseCases deduplicateUseCases
	cleanupUseCases     cleanupUseCases

	ogenServer *server.Server

	serverAddr string

	externalServerScheme       string
	externalServerHostWithPort string
	token                      string
}

func New(
	logger logger,
	tracer trace.Tracer,
	serverAddr string,
	externalServerAddr string,
	parseUseCases parseUseCases,
	webAPIUseCases webAPIUseCases,
	agentUseCases agentUseCases,
	exportUseCases exportUseCases,
	deduplicateUseCases deduplicateUseCases,
	cleanupUseCases cleanupUseCases,
	debug bool,
	staticDir string,
	token string,
) (*Controller, error) {
	u, err := url.Parse(externalServerAddr)
	if err != nil {
		return nil, fmt.Errorf("parse external server addr: %w", err)
	}

	c := &Controller{
		logger:                     logger,
		tracer:                     tracer,
		serverAddr:                 serverAddr,
		externalServerScheme:       u.Scheme,
		externalServerHostWithPort: u.Host,
		parseUseCases:              parseUseCases,
		webAPIUseCases:             webAPIUseCases,
		agentUseCases:              agentUseCases,
		exportUseCases:             exportUseCases,
		deduplicateUseCases:        deduplicateUseCases,
		cleanupUseCases:            cleanupUseCases,
		debug:                      debug,
		staticDir:                  staticDir,
		token:                      token,
	}

	ogenServer, err := server.NewServer(
		c, c,
		// server.WithErrorHandler(), // FIXME: реализовать
		// server.WithMethodNotAllowed(), // FIXME: реализовать
		// server.WithNotFound(), // FIXME: реализовать
	)
	if err != nil {
		return nil, fmt.Errorf("create ogen server: %w", err)
	}

	c.ogenServer = ogenServer

	return c, nil
}
