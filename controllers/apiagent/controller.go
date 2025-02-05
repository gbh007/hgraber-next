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
	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

type parsingUseCases interface {
	CheckBooks(ctx context.Context, urls []url.URL) ([]agentmodel.AgentBookCheckResult, error)
	ParseBook(ctx context.Context, u url.URL) (agentmodel.AgentBookDetails, error)
	DownloadPage(ctx context.Context, bookURL, imageURL url.URL) (io.Reader, error)
	CheckPages(ctx context.Context, pages []agentmodel.AgentPageURL) ([]agentmodel.AgentPageCheckResult, error)
}

type exportUseCases interface {
	ImportArchive(ctx context.Context, body io.Reader, deduplicate bool, autoVerify bool) (uuid.UUID, error)
}

type Controller struct {
	startAt time.Time
	logger  *slog.Logger
	tracer  trace.Tracer
	addr    string
	debug   bool

	ogenServer *agentapi.Server

	parsingUseCases parsingUseCases
	exportUseCases  exportUseCases

	token string
}

func New(
	startAt time.Time,
	logger *slog.Logger,
	tracer trace.Tracer,
	parsingUseCases parsingUseCases,
	exportUseCases exportUseCases,
	addr string,
	debug bool,
	token string,
) (*Controller, error) {
	c := &Controller{
		startAt: startAt,
		logger:  logger,
		tracer:  tracer,
		addr:    addr,
		debug:   debug,
		token:   token,

		parsingUseCases: parsingUseCases,
		exportUseCases:  exportUseCases,
	}

	ogenServer, err := agentapi.NewServer(
		c, c,
		agentapi.WithErrorHandler(methodErrorHandler),
		agentapi.WithMethodNotAllowed(methodNotAllowed),
		agentapi.WithNotFound(methodNotFound),
		agentapi.WithMiddleware(c.simplePanicRecover),
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
