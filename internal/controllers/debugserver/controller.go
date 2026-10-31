package debugserver

import (
	"context"
	"io"
	"log/slog"
	"net/url"

	"github.com/gbh007/hgraber-next/internal/agent/config"
	"github.com/gbh007/hgraber-next/internal/domain/entities"
)

type ParsingUseCases interface {
	ParseBook(ctx context.Context, u url.URL) (entities.AgentBookDetails, error)
	DownloadPage(ctx context.Context, bookURL, imageURL url.URL) (io.Reader, error)
	MultiHandle(ctx context.Context, multiUrl url.URL) ([]entities.AgentBookCheckResult, error)
}

type Controller[T any] struct {
	config          config.AgentConfig[T]
	logger          *slog.Logger
	parsing         ParsingUseCases
	addr            string
	debug           bool
	logErrorHandler bool
}

func New[T any](
	config config.AgentConfig[T],
	logger *slog.Logger,
	parsing ParsingUseCases,
) *Controller[T] {
	c := &Controller[T]{
		config:          config,
		logger:          logger,
		parsing:         parsing,
		addr:            config.DebugServer.Addr,
		debug:           config.DebugServer.Debug,
		logErrorHandler: config.DebugServer.LogErrorHandler,
	}

	return c
}
