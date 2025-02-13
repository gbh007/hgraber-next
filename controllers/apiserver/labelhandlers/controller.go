package labelhandlers

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
)

type LabelUseCases interface {
	SetLabel(ctx context.Context, label core.BookLabel) error
	DeleteLabel(ctx context.Context, label core.BookLabel) error
	Labels(ctx context.Context, bookID uuid.UUID) ([]core.BookLabel, error)
	CreateLabelPreset(ctx context.Context, preset core.BookLabelPreset) error
	UpdateLabelPreset(ctx context.Context, preset core.BookLabelPreset) error
	DeleteLabelPreset(ctx context.Context, name string) error
	LabelPresets(ctx context.Context) ([]core.BookLabelPreset, error)
	LabelPreset(ctx context.Context, name string) (core.BookLabelPreset, error)
}

type LabelHandlersController struct {
	logger *slog.Logger
	tracer trace.Tracer
	debug  bool

	apiCore *apiservercore.Controller

	labelUseCases LabelUseCases
}

func New(
	logger *slog.Logger,
	tracer trace.Tracer,
	labelUseCases LabelUseCases,
	debug bool,
	ac *apiservercore.Controller,
) *LabelHandlersController {
	c := &LabelHandlersController{
		logger:        logger,
		tracer:        tracer,
		labelUseCases: labelUseCases,
		debug:         debug,
		apiCore:       ac,
	}

	return c
}
