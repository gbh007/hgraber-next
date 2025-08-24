package massloadhandlers

import (
	"context"
	"log/slog"
	"net/url"

	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/massloadmodel"
)

type MassloadUseCases interface {
	CreateMassload(ctx context.Context, ml massloadmodel.Massload) (int, error)
	UpdateMassload(ctx context.Context, ml massloadmodel.Massload) error
	DeleteMassload(ctx context.Context, id int) error
	Massload(ctx context.Context, id int) (massloadmodel.Massload, error)
	Massloads(ctx context.Context, filter massloadmodel.Filter) ([]massloadmodel.Massload, error)

	MassloadFlags(ctx context.Context) ([]massloadmodel.Flag, error)

	CreateMassloadAttribute(ctx context.Context, massloadID int, code, value string) error
	DeleteMassloadAttribute(ctx context.Context, massloadID int, code, value string) error

	CreateMassloadExternalLink(ctx context.Context, massloadID int, url url.URL) error
	DeleteMassloadExternalLink(ctx context.Context, massloadID int, url url.URL) error
}

type MassloadController struct {
	logger *slog.Logger
	tracer trace.Tracer
	debug  bool

	apiCore *apiservercore.Controller

	massloadUseCases MassloadUseCases
}

func New(
	logger *slog.Logger,
	tracer trace.Tracer,
	debug bool,
	ac *apiservercore.Controller,
	massloadUseCases MassloadUseCases,
) *MassloadController {
	c := &MassloadController{
		logger:           logger,
		tracer:           tracer,
		debug:            debug,
		apiCore:          ac,
		massloadUseCases: massloadUseCases,
	}

	return c
}
