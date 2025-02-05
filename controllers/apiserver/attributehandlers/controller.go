package attributehandlers

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
)

type WebAPIUseCases interface {
	AttributesCount(ctx context.Context) ([]core.AttributeVariant, error)
	CreateAttributeColor(ctx context.Context, color core.AttributeColor) error
	UpdateAttributeColor(ctx context.Context, color core.AttributeColor) error
	DeleteAttributeColor(ctx context.Context, code, value string) error
	AttributeColors(ctx context.Context) ([]core.AttributeColor, error)
	AttributeColor(ctx context.Context, code, value string) (core.AttributeColor, error)
}

type AttributeHandlersController struct {
	logger *slog.Logger
	tracer trace.Tracer
	debug  bool

	apiCore *apiservercore.Controller

	webAPIUseCases WebAPIUseCases
}

func New(
	logger *slog.Logger,
	tracer trace.Tracer,
	webAPIUseCases WebAPIUseCases,
	debug bool,
	ac *apiservercore.Controller,
) *AttributeHandlersController {
	c := &AttributeHandlersController{
		logger:         logger,
		tracer:         tracer,
		webAPIUseCases: webAPIUseCases,
		debug:          debug,
		apiCore:        ac,
	}

	return c
}
