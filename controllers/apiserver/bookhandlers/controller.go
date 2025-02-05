package bookhandlers

import (
	"context"
	"io"
	"log/slog"
	"net/url"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/bff"
	"github.com/gbh007/hgraber-next/domain/core"
)

type ParseUseCases interface {
	BookByURL(ctx context.Context, u url.URL) (core.BookContainer, error)
}

type WebAPIUseCases interface {
	BookRaw(ctx context.Context, bookID uuid.UUID) (core.BookContainer, error)
	VerifyBook(ctx context.Context, bookID uuid.UUID, verified bool) error
	DeleteBook(ctx context.Context, bookID uuid.UUID) error
}

type ExportUseCases interface {
	ExportBook(ctx context.Context, bookID uuid.UUID) (io.Reader, core.BookContainer, error)
}

type ReBuilderUseCases interface {
	UpdateBook(ctx context.Context, book core.BookContainer) error
	RebuildBook(ctx context.Context, request core.RebuildBookRequest) (uuid.UUID, error)
	RestoreBook(ctx context.Context, bookID uuid.UUID, onlyPages bool) error
}

type BFFUseCases interface {
	BookDetails(ctx context.Context, bookID uuid.UUID) (bff.BookDetails, error)
	BookList(ctx context.Context, filter core.BookFilter) (bff.BookList, error)
}

type BookHandlersController struct {
	logger *slog.Logger
	tracer trace.Tracer
	debug  bool

	apiCore *apiservercore.Controller

	parseUseCases     ParseUseCases
	webAPIUseCases    WebAPIUseCases
	exportUseCases    ExportUseCases
	rebuilderUseCases ReBuilderUseCases
	bffUseCases       BFFUseCases
}

func New(
	logger *slog.Logger,
	tracer trace.Tracer,
	parseUseCases ParseUseCases,
	webAPIUseCases WebAPIUseCases,
	exportUseCases ExportUseCases,
	rebuilderUseCases ReBuilderUseCases,
	bffUseCases BFFUseCases,
	debug bool,
	ac *apiservercore.Controller,
) *BookHandlersController {
	c := &BookHandlersController{
		logger:            logger,
		tracer:            tracer,
		parseUseCases:     parseUseCases,
		webAPIUseCases:    webAPIUseCases,
		exportUseCases:    exportUseCases,
		rebuilderUseCases: rebuilderUseCases,
		bffUseCases:       bffUseCases,
		debug:             debug,
		apiCore:           ac,
	}

	return c
}
