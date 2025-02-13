package systemhandlers

import (
	"context"
	"io"
	"log/slog"
	"net/url"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/domain/parsing"
	"github.com/gbh007/hgraber-next/domain/systemmodel"
)

type ParseUseCases interface {
	NewBooks(ctx context.Context, urls []url.URL, flags parsing.ParseFlags) (parsing.FirstHandleMultipleResult, error)
	NewBooksMulti(ctx context.Context, urls []url.URL, flags parsing.ParseFlags) (parsing.MultiHandleMultipleResult, error)

	NewMirror(ctx context.Context, mirror parsing.URLMirror) error
	UpdateMirror(ctx context.Context, mirror parsing.URLMirror) error
	DeleteMirror(ctx context.Context, id uuid.UUID) error
	Mirrors(ctx context.Context) ([]parsing.URLMirror, error)
	Mirror(ctx context.Context, id uuid.UUID) (parsing.URLMirror, error)
}

type ExportUseCases interface {
	ImportArchive(ctx context.Context, body io.Reader, deduplicate bool, autoVerify bool) (uuid.UUID, error)
}

type DeduplicateUseCases interface {
	ArchiveEntryPercentage(ctx context.Context, archiveBody io.Reader) ([]core.DeduplicateArchiveResult, error)
}

type SystemUseCases interface {
	RunTask(ctx context.Context, code systemmodel.TaskCode) error
	TaskResults(ctx context.Context) ([]*systemmodel.TaskResult, error)
	SystemSize(ctx context.Context) (systemmodel.SystemSizeInfo, error)
	WorkersInfo(ctx context.Context) []systemmodel.SystemWorkerStat
	SetWorkerConfig(ctx context.Context, counts map[string]int)
}

type SystemHandlersController struct {
	logger *slog.Logger
	tracer trace.Tracer
	debug  bool

	apiCore *apiservercore.Controller

	parseUseCases       ParseUseCases
	exportUseCases      ExportUseCases
	deduplicateUseCases DeduplicateUseCases
	systemUseCases      SystemUseCases
}

func New(
	logger *slog.Logger,
	tracer trace.Tracer,
	parseUseCases ParseUseCases,
	exportUseCases ExportUseCases,
	deduplicateUseCases DeduplicateUseCases,
	systemUseCases SystemUseCases,
	debug bool,
	ac *apiservercore.Controller,
) *SystemHandlersController {
	c := &SystemHandlersController{
		logger:              logger,
		tracer:              tracer,
		parseUseCases:       parseUseCases,
		exportUseCases:      exportUseCases,
		deduplicateUseCases: deduplicateUseCases,
		systemUseCases:      systemUseCases,
		debug:               debug,
		apiCore:             ac,
	}

	return c
}
