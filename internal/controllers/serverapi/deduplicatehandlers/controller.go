package deduplicatehandlers

import (
	"context"
	"io"
	"log/slog"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/internal/controllers/serverapi/apiservercore"
	"github.com/gbh007/hgraber-next/internal/domain/bff"
	"github.com/gbh007/hgraber-next/internal/domain/core"
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

	MarkBookPagesAsDeadHash(ctx context.Context, bookID uuid.UUID) error
	UnMarkBookPagesAsDeadHash(ctx context.Context, bookID uuid.UUID) error

	ArchiveEntryPercentage(ctx context.Context, archiveBody io.Reader) ([]core.DeduplicateArchiveResult, error)
}

type DeduplicateHandlersController struct {
	logger *slog.Logger
	debug  bool

	apiCore *apiservercore.Controller

	bffUseCases         BFFUseCases
	deduplicateUseCases DeduplicateUseCases
}

func New(
	logger *slog.Logger,
	bffUseCases BFFUseCases,
	deduplicateUseCases DeduplicateUseCases,
	debug bool,
	ac *apiservercore.Controller,
) *DeduplicateHandlersController {
	c := &DeduplicateHandlersController{
		logger:              logger,
		bffUseCases:         bffUseCases,
		deduplicateUseCases: deduplicateUseCases,
		debug:               debug,
		apiCore:             ac,
	}

	return c
}
