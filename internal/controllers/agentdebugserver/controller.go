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

type Controller struct {
	config          agnconfig.AgentConfig
	logger          *slog.Logger
	parsing         ParsingUseCases
	addr            string
	debug           bool
	logErrorHandler bool
}

func New(
	config agnconfig.AgentConfig,
	logger *slog.Logger,
	parsing ParsingUseCases,
) *Controller {
	c := &Controller{
		config:          config,
		logger:          logger,
		parsing:         parsing,
		addr:            config.DebugServer.Addr,
		debug:           config.DebugServer.Debug,
		logErrorHandler: config.DebugServer.LogErrorHandler,
	}

	return c
}
