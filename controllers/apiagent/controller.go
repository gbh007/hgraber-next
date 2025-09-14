package apiagent

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"net/url"
	"time"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

type parsingUseCases interface {
	BooksExists(ctx context.Context, urls []url.URL) ([]agentmodel.AgentBookCheckResult, error)
	PagesExists(ctx context.Context, urls []agentmodel.AgentPageURL) ([]agentmodel.AgentPageCheckResult, error)
	BookByURL(ctx context.Context, u url.URL) (core.BookContainer, error)
	PageBodyByURL(ctx context.Context, u url.URL) (io.Reader, error)
}

type exportUseCases interface {
	ImportArchive(ctx context.Context, body io.Reader, deduplicate, autoVerify bool) (uuid.UUID, error)
}

type metricProvider interface {
	HTTPServerAddHandle(addr, operation string, status bool, d time.Duration)
	HTTPServerIncActive(addr, operation string)
	HTTPServerDecActive(addr, operation string)
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

	parsingUseCases parsingUseCases
	exportUseCases  exportUseCases

	token string
}

type config interface {
	GetAddr() string
	GetToken() string
	GetLogErrorHandler() bool
	GetDebug() bool
}

func New(
	config config,
	startAt time.Time,
	logger *slog.Logger,
	tracer trace.Tracer,
	parsingUseCases parsingUseCases,
	exportUseCases exportUseCases,
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

		parsingUseCases: parsingUseCases,
		exportUseCases:  exportUseCases,
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

func (c *Controller) HandleHeaderAuth(
	ctx context.Context,
	operationName string,
	t agentapi.HeaderAuth,
) (context.Context, error) {
	if c.token == "" {
		return ctx, nil
	}

	if c.token != t.APIKey {
		return ctx, errorAccessForbidden
	}

	return ctx, nil
}
