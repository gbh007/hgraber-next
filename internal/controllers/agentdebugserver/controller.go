package agentdebugserver

import (
	"context"
	"io"
	"log/slog"
	"net/url"

	"github.com/gbh007/hgraber-next/external/parserdto"
	"github.com/gbh007/hgraber-next/internal/application/agent/agnconfig"
)

type ParsingUseCases interface {
	Book(ctx context.Context, u url.URL, pageLimit *int) (parserdto.Book, error)
	Image(ctx context.Context, bookURL, imageURL url.URL) (io.ReadCloser, error)
	List(ctx context.Context, u url.URL, onePage bool) (parserdto.BookList, error)
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
