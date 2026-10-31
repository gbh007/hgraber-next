package agentdebugserver

import (
	"context"
	"io"
	"log/slog"
	"net/url"

	"github.com/gbh007/hgraber-next/internal/application/agent/agnconfig"
	"github.com/gbh007/hgraber-next/internal/domain/serveragentdto"
)

type ParsingUseCases interface {
	ParseBook(ctx context.Context, u url.URL) (serveragentdto.AgentBookDetails, error)
	DownloadPage(ctx context.Context, bookURL, imageURL url.URL) (io.Reader, error)
	MultiHandle(ctx context.Context, multiUrl url.URL) ([]serveragentdto.AgentBookCheckResult, error)
}

type Controller[T any] struct {
	config          agnconfig.AgentConfig[T]
	logger          *slog.Logger
	parsing         ParsingUseCases
	addr            string
	debug           bool
	logErrorHandler bool
}

func New[T any](
	config agnconfig.AgentConfig[T],
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
