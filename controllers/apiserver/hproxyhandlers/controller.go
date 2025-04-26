package hproxyhandlers

import (
	"context"
	"io"
	"log/slog"
	"net/url"

	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/hproxymodel"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

type HProxyUseCases interface {
	List(ctx context.Context, u url.URL) (hproxymodel.List, error)
	Book(ctx context.Context, u url.URL) (hproxymodel.Book, error)
	Image(ctx context.Context, bookURL, imageURL url.URL) (io.Reader, error)
}

type HProxyHandlersController struct {
	logger *slog.Logger
	tracer trace.Tracer
	debug  bool

	apiCore *apiservercore.Controller

	hProxyUseCases HProxyUseCases
}

func New(
	logger *slog.Logger,
	tracer trace.Tracer,
	hProxyUseCases HProxyUseCases,
	debug bool,
	ac *apiservercore.Controller,
) *HProxyHandlersController {
	c := &HProxyHandlersController{
		logger:         logger,
		tracer:         tracer,
		hProxyUseCases: hProxyUseCases,
		debug:          debug,
		apiCore:        ac,
	}

	return c
}

func (c *HProxyHandlersController) filePreview(bookURL url.URL, imageURL *url.URL) serverapi.OptURI {
	if imageURL == nil {
		return serverapi.OptURI{}
	}

	return serverapi.NewOptURI(c.apiCore.GetHProxyFileURL(bookURL, *imageURL))
}
