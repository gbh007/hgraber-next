package apiserver

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/url"

	"github.com/google/uuid"

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

type Controller struct {
	logger    *slog.Logger
	debug     bool
	staticDir string

	parseUseCases  parseUseCases
	webAPIUseCases webAPIUseCases

	ogenServer *server.Server

	serverAddr string

	externalServerScheme       string
	externalServerHostWithPort string
	token                      string
}

func New(
	logger *slog.Logger,
	serverAddr string,
	externalServerAddr string,
	parseUseCases parseUseCases,
	webAPIUseCases webAPIUseCases,
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
		serverAddr:                 serverAddr,
		externalServerScheme:       u.Scheme,
		externalServerHostWithPort: u.Host,
		parseUseCases:              parseUseCases,
		webAPIUseCases:             webAPIUseCases,
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
