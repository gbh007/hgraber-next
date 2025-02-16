package deduplicatehandlers

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/bff"
)

type BFFUseCases interface {
	BookCompare(ctx context.Context, originID, targetID uuid.UUID) (bff.BookCompareResult, error)
}

type DeduplicateUseCases interface {
	BookByPageEntryPercentage(ctx context.Context, originBookID uuid.UUID) ([]bff.DeduplicateBookResult, error)
	UniquePages(ctx context.Context, originBookID uuid.UUID) ([]bff.PreviewPage, error)
	BooksByPage(ctx context.Context, bookID uuid.UUID, pageNumber int) ([]bff.BookWithPreviewPage, error)

	CreateDeadHashByPage(ctx context.Context, bookID uuid.UUID, pageNumber int) error
	DeleteDeadHashByPage(ctx context.Context, bookID uuid.UUID, pageNumber int) error
	DeleteAllPageByHash(ctx context.Context, bookID uuid.UUID, pageNumber int, setDeadHash bool) error

	MarkBookPagesAsDeadHash(ctx context.Context, bookID uuid.UUID) error
	UnMarkBookPagesAsDeadHash(ctx context.Context, bookID uuid.UUID) error
}

type DeduplicateHandlersController struct {
	logger *slog.Logger
	tracer trace.Tracer
	debug  bool

	apiCore *apiservercore.Controller

	bffUseCases         BFFUseCases
	deduplicateUseCases DeduplicateUseCases
}

func New(
	logger *slog.Logger,
	tracer trace.Tracer,
	bffUseCases BFFUseCases,
	deduplicateUseCases DeduplicateUseCases,
	debug bool,
	ac *apiservercore.Controller,
) *DeduplicateHandlersController {
	c := &DeduplicateHandlersController{
		logger:              logger,
		tracer:              tracer,
		bffUseCases:         bffUseCases,
		deduplicateUseCases: deduplicateUseCases,
		debug:               debug,
		apiCore:             ac,
	}

	return c
}
