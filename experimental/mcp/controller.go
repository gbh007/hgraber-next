package mcp

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/mark3labs/mcp-go/server"
	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/domain/bff"
	"github.com/gbh007/hgraber-next/domain/core"
)

type bffUseCases interface {
	BookDetails(ctx context.Context, bookID uuid.UUID) (bff.BookDetails, error)
	BookList(ctx context.Context, filter core.BookFilter) (bff.BookList, error)
}

type attrUseCases interface {
	AttributesCount(ctx context.Context) ([]core.AttributeVariant, error)
}

type Controller struct {
	logger       *slog.Logger
	tracer       trace.Tracer
	addr         string
	token        string
	debug        bool
	bffUseCases  bffUseCases
	attrUseCases attrUseCases
}

func New(
	logger *slog.Logger,
	tracer trace.Tracer,
	addr string,
	token string,
	bffUseCases bffUseCases,
	attrUseCases attrUseCases,
	debug bool,
) *Controller {
	return &Controller{
		logger:       logger,
		tracer:       tracer,
		addr:         addr,
		token:        token,
		bffUseCases:  bffUseCases,
		attrUseCases: attrUseCases,
		debug:        debug,
	}
}

func (c *Controller) Start(parentCtx context.Context) (chan struct{}, error) {
	done := make(chan struct{})

	s := server.NewMCPServer(
		"hgraber-next",
		"0.0.1",
	)

	s.AddTools(
		c.bookDetailsTool(),
		c.bookListTool(),
		c.attributesCountTool(),
	)

	httpMux := server.NewStreamableHTTPServer(s)

	server := &http.Server{ //nolint:gosec // будет исправлено позднее
		Handler:  c.logIO(c.authMiddleware(httpMux)),
		Addr:     c.addr,
		ErrorLog: slog.NewLogLogger(c.logger.Handler(), slog.LevelError),
	}

	go func() {
		defer close(done)

		c.logger.InfoContext(parentCtx, "mcp server start")
		defer c.logger.InfoContext(parentCtx, "mcp server stop")

		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			c.logger.ErrorContext(parentCtx, err.Error())
		}
	}()

	go func() {
		<-parentCtx.Done()
		c.logger.InfoContext(parentCtx, "stopping mcp server")

		shutdownCtx, cancel := context.WithTimeout(context.WithoutCancel(parentCtx), time.Second*10) //nolint:mnd,lll,golines // будет исправлено позднее
		defer cancel()

		err := server.Shutdown(shutdownCtx)
		if err != nil {
			c.logger.ErrorContext(parentCtx, err.Error())
		}
	}()

	return done, nil
}

func (c *Controller) Name() string {
	return "mcp server"
}
