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

	ogenServer *server.Server

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

// FIXME: реализовать
func (c *Controller) APIBookRawPost(ctx context.Context, req *server.APIBookRawPostReq) (server.APIBookRawPostRes, error) {
	return &server.APIBookRawPostInternalServerError{
		InnerCode: "unimplemented",
	}, nil
}

// FIXME: реализовать
func (c *Controller) APIPageBodyPost(ctx context.Context, req *server.APIPageBodyPostReq) (server.APIPageBodyPostRes, error) {
	return &server.APIPageBodyPostInternalServerError{
		InnerCode: "unimplemented",
	}, nil
}

// FIXME: реализовать
func (c *Controller) APIParsingBookExistsPost(ctx context.Context, req *server.APIParsingBookExistsPostReq) (server.APIParsingBookExistsPostRes, error) {
	return &server.APIParsingBookExistsPostInternalServerError{
		InnerCode: "unimplemented",
	}, nil
}

// FIXME: реализовать
func (c *Controller) APIParsingPageExistsPost(ctx context.Context, req *server.APIParsingPageExistsPostReq) (server.APIParsingPageExistsPostRes, error) {
	return &server.APIParsingPageExistsPostInternalServerError{
		InnerCode: "unimplemented",
	}, nil
}
