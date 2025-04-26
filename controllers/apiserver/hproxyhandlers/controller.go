package hproxyhandlers

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

type HProxyUseCases interface{}

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

func (c *HProxyHandlersController) APIHproxyBookPost(ctx context.Context, req *serverapi.APIHproxyBookPostReq) (serverapi.APIHproxyBookPostRes, error) {
	return &serverapi.APIHproxyBookPostInternalServerError{
		InnerCode: apiservercore.HProxyUseCaseCode,
		Details:   serverapi.NewOptString("unimplemented"),
	}, nil
}

func (c *HProxyHandlersController) APIHproxyFileGet(ctx context.Context, params serverapi.APIHproxyFileGetParams) (serverapi.APIHproxyFileGetRes, error) {
	return &serverapi.APIHproxyFileGetInternalServerError{
		InnerCode: apiservercore.HProxyUseCaseCode,
		Details:   serverapi.NewOptString("unimplemented"),
	}, nil
}

func (c *HProxyHandlersController) APIHproxyListPost(ctx context.Context, req *serverapi.APIHproxyListPostReq) (serverapi.APIHproxyListPostRes, error) {
	return &serverapi.APIHproxyListPostInternalServerError{
		InnerCode: apiservercore.HProxyUseCaseCode,
		Details:   serverapi.NewOptString("unimplemented"),
	}, nil
}
