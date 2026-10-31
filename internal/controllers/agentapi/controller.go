package agentapi

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"net/url"
	"time"

	"github.com/gbh007/hgraber-next/external/parserdto"
	"github.com/gbh007/hgraber-next/internal/domain/agentdto"
	"github.com/gbh007/hgraber-next/internal/domain/fsmodel"
	"github.com/gbh007/hgraber-next/internal/domain/serveragentdto"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
	"github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/otel/trace"

	"github.com/google/uuid"
)

type metricProvider interface {
	HTTPServerAddHandle(addr, operation string, status bool, d time.Duration)
	HTTPServerIncActive(addr, operation string)
	HTTPServerDecActive(addr, operation string)
	Registry() *prometheus.Registry
}

type ParsingUseCases interface {
	CheckBooks(ctx context.Context, urls []url.URL) ([]serveragentdto.AgentBookCheckResult, error)
	ParseBook(ctx context.Context, u url.URL) (serveragentdto.AgentBookDetails, error)
	DownloadPage(ctx context.Context, bookURL, imageURL url.URL) (io.Reader, error)
	MultiHandle(ctx context.Context, multiUrl url.URL) ([]serveragentdto.AgentBookCheckResult, error)

	HProxyBook(ctx context.Context, u url.URL, pageLimit *int) (parserdto.Book, error)
	HProxyList(ctx context.Context, u url.URL) (parserdto.BookList, error)
}

type ImportUseCases interface {
	Create(ctx context.Context, data agentdto.ImportData) error
}

type FileUseCases interface {
	Create(ctx context.Context, fileID uuid.UUID, body io.Reader) error
	Delete(ctx context.Context, fileID uuid.UUID) error
	Get(ctx context.Context, fileID uuid.UUID) (io.Reader, error)
	State(ctx context.Context, includeFileIDs, includeFileSizes bool) (fsmodel.FSState, error)
}

type HighwayUseCases interface {
	NewToken(ctx context.Context) (string, time.Time, error)
	ValidateToken(ctx context.Context, token string) error
	Get(ctx context.Context, fileID uuid.UUID) (io.Reader, error)
}

type config interface {
	GetAddr() string
	GetToken() string
	GetLogErrorHandler() bool
	GetDebug() bool
}

type Controller struct {
	startAt         time.Time
	logger          *slog.Logger
	tracer          trace.Tracer
	addr            string
	debug           bool
	logErrorHandler bool

	metricProvider metricProvider

	ogenServer *agentapi.Server

	importUseCase   ImportUseCases
	fileUseCase     FileUseCases
	parsingUseCases ParsingUseCases
	highwayUseCase  HighwayUseCases

	token          string
	parserCodes    []string
	enabledModules []string
}

func New(
	config config,
	startAt time.Time,
	logger *slog.Logger,
	tracer trace.Tracer,
	parsingUseCases ParsingUseCases,
	importUseCase ImportUseCases,
	fileUseCase FileUseCases,
	highwayUseCase HighwayUseCases,
	parserCodes []string,
	metricProvider metricProvider,
) (*Controller, error) {
	c := &Controller{
		startAt:         startAt,
		logger:          logger,
		tracer:          tracer,
		addr:            config.GetAddr(),
		debug:           config.GetDebug(),
		logErrorHandler: config.GetLogErrorHandler(),
		token:           config.GetToken(),
		metricProvider:  metricProvider,

		parserCodes:    parserCodes,
		enabledModules: make([]string, 0, 3),

		parsingUseCases: parsingUseCases,
		importUseCase:   importUseCase,
		fileUseCase:     fileUseCase,
		highwayUseCase:  highwayUseCase,
	}

	if c.parsingUseCases != nil {
		c.enabledModules = append(c.enabledModules, "parsing", "hproxy")
	}

	if c.importUseCase != nil {
		c.enabledModules = append(c.enabledModules, "import")
	}

	if c.fileUseCase != nil {
		c.enabledModules = append(c.enabledModules, "file system")
	}

	if c.highwayUseCase != nil {
		c.enabledModules = append(c.enabledModules, "highway")
	}

	ogenServer, err := agentapi.NewServer(
		c, c,
		agentapi.WithErrorHandler(c.methodErrorHandler),
		agentapi.WithMethodNotAllowed(methodNotAllowed),
		agentapi.WithNotFound(methodNotFound),
		agentapi.WithMiddleware(
			c.metricsMiddleware,
			c.simplePanicRecover,
		),
	)
	if err != nil {
		return nil, err
	}

	c.ogenServer = ogenServer

	return c, nil
}

var errorAccessForbidden = errors.New("access forbidden")

func (c *Controller) HandleHeaderAuth(ctx context.Context, operationName string, t agentapi.HeaderAuth) (context.Context, error) {
	if c.token == "" {
		return ctx, nil
	}

	if c.token != t.APIKey {
		return ctx, errorAccessForbidden
	}

	return ctx, nil
}
