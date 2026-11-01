package hproxyhandlers

import (
	"context"
	"io"
	"log/slog"
	"net/url"

	"github.com/gbh007/hgraber-next/internal/controllers/serverapi/apiservercore"
	"github.com/gbh007/hgraber-next/internal/domain/hproxymodel"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

type HProxyUseCases interface {
	List(ctx context.Context, u url.URL) (hproxymodel.List, error)
	Book(ctx context.Context, u url.URL, pageLimit *int) (hproxymodel.Book, error)
	Image(ctx context.Context, bookURL, imageURL url.URL) (io.Reader, error)
}

type HProxyHandlersController struct {
	logger *slog.Logger
	debug  bool

	apiCore *apiservercore.Controller

	hProxyUseCases HProxyUseCases
}

func New(
	logger *slog.Logger,
	hProxyUseCases HProxyUseCases,
	debug bool,
	ac *apiservercore.Controller,
) *HProxyHandlersController {
	c := &HProxyHandlersController{
		logger:         logger,
		hProxyUseCases: hProxyUseCases,
		debug:          debug,
		apiCore:        ac,
	}

	return c
}

func (c *HProxyHandlersController) filePreview(ctx context.Context, bookURL url.URL, imageURL *url.URL) serverapi.OptURI {
	if imageURL == nil {
		return serverapi.OptURI{}
	}

	return serverapi.NewOptURI(c.apiCore.GetHProxyFileURL(ctx, bookURL, *imageURL))
}
