package apiagent

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"net/url"
	"time"

	"hgnext/internal/entities"
	"hgnext/open_api/agentAPI"
)

type parsingUseCases interface {
	CheckBooks(ctx context.Context, urls []url.URL) ([]entities.AgentBookCheckResult, error)
	ParseBook(ctx context.Context, u url.URL) (entities.AgentBookDetails, error)
	DownloadPage(ctx context.Context, bookURL, imageURL url.URL) (io.Reader, error)
	CheckPages(ctx context.Context, pages []entities.AgentPageURL) ([]entities.AgentPageCheckResult, error)
}

type Controller struct {
	startAt time.Time
	logger  *slog.Logger
	addr    string
	debug   bool

	ogenServer *agentAPI.Server

	parsingUseCases parsingUseCases

	token string
}

func New(
	startAt time.Time,
	logger *slog.Logger,
	parsingUseCases parsingUseCases,
	addr string,
	debug bool,
	token string,
) (*Controller, error) {
	c := &Controller{
		startAt: startAt,
		logger:  logger,
		addr:    addr,
		debug:   debug,
		token:   token,

		parsingUseCases: parsingUseCases,
	}

	ogenServer, err := agentAPI.NewServer(
		c, c,
		// agentAPI.WithErrorHandler(), // FIXME: реализовать
		// agentAPI.WithMethodNotAllowed(), // FIXME: реализовать
		// agentAPI.WithNotFound(), // FIXME: реализовать
	)
	if err != nil {
		return nil, err
	}

	c.ogenServer = ogenServer

	return c, nil
}

var errorAccessForbidden = errors.New("access forbidden")

func (c *Controller) HandleHeaderAuth(ctx context.Context, operationName string, t agentAPI.HeaderAuth) (context.Context, error) {
	if c.token == "" {
		return ctx, nil
	}

	if c.token != t.APIKey {
		return ctx, errorAccessForbidden
	}

	return ctx, nil
}
