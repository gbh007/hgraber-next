package attributehandlers

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
)

type AttributeUseCases interface {
	AttributesCount(ctx context.Context) ([]core.AttributeVariant, error)
	CreateAttributeColor(ctx context.Context, color core.AttributeColor) error
	UpdateAttributeColor(ctx context.Context, color core.AttributeColor) error
	DeleteAttributeColor(ctx context.Context, code, value string) error
	AttributeColors(ctx context.Context) ([]core.AttributeColor, error)
	AttributeColor(ctx context.Context, code, value string) (core.AttributeColor, error)

	CreateAttributeRemap(ctx context.Context, ar core.AttributeRemap) error
	UpdateAttributeRemap(ctx context.Context, ar core.AttributeRemap) error
	DeleteAttributeRemap(ctx context.Context, code, value string) error
	AttributeRemaps(ctx context.Context) ([]core.AttributeRemap, error)
	AttributeRemap(ctx context.Context, code, value string) (core.AttributeRemap, error)
}

type AttributeHandlersController struct {
	logger *slog.Logger
	tracer trace.Tracer
	debug  bool

	apiCore *apiservercore.Controller

	attributeUseCases AttributeUseCases
}

func New(
	logger *slog.Logger,
	tracer trace.Tracer,
	attributeUseCases AttributeUseCases,
	debug bool,
	ac *apiservercore.Controller,
) *AttributeHandlersController {
	c := &AttributeHandlersController{
		logger:            logger,
		tracer:            tracer,
		attributeUseCases: attributeUseCases,
		debug:             debug,
		apiCore:           ac,
	}

	return c
}
