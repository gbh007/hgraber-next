package systemhandlers

import (
	"context"
	"io"
	"log/slog"
	"net/url"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/domain/parsing"
	"github.com/gbh007/hgraber-next/domain/systemmodel"
)

type ParseUseCases interface {
	NewBooks(ctx context.Context, urls []url.URL, flags parsing.ParseFlags) (parsing.FirstHandleMultipleResult, error)
	NewBooksMulti(ctx context.Context, urls []url.URL, flags parsing.ParseFlags) (parsing.MultiHandleMultipleResult, error)

	BooksExists(ctx context.Context, urls []url.URL) ([]agentmodel.AgentBookCheckResult, error)
	PagesExists(ctx context.Context, urls []agentmodel.AgentPageURL) ([]agentmodel.AgentPageCheckResult, error)

	NewMirror(ctx context.Context, mirror parsing.URLMirror) error
	UpdateMirror(ctx context.Context, mirror parsing.URLMirror) error
	DeleteMirror(ctx context.Context, id uuid.UUID) error
	Mirrors(ctx context.Context) ([]parsing.URLMirror, error)
	Mirror(ctx context.Context, id uuid.UUID) (parsing.URLMirror, error)
}

type WebAPIUseCases interface {
	SystemSize(ctx context.Context) (systemmodel.SystemSizeInfo, error)
	WorkersInfo(ctx context.Context) []systemmodel.SystemWorkerStat
	SetWorkerConfig(ctx context.Context, counts map[string]int)
}

type ExportUseCases interface {
	ImportArchive(ctx context.Context, body io.Reader, deduplicate bool, autoVerify bool) (uuid.UUID, error)
}

type DeduplicateUseCases interface {
	ArchiveEntryPercentage(ctx context.Context, archiveBody io.Reader) ([]core.DeduplicateArchiveResult, error)
}

type TaskUseCases interface {
	RunTask(ctx context.Context, code systemmodel.TaskCode) error
	TaskResults(ctx context.Context) ([]*systemmodel.TaskResult, error)
}

type SystemHandlersController struct {
	logger *slog.Logger
	tracer trace.Tracer
	debug  bool

	apiCore *apiservercore.Controller

	parseUseCases       ParseUseCases
	webAPIUseCases      WebAPIUseCases
	exportUseCases      ExportUseCases
	deduplicateUseCases DeduplicateUseCases
	taskUseCases        TaskUseCases
}

func New(
	logger *slog.Logger,
	tracer trace.Tracer,
	parseUseCases ParseUseCases,
	webAPIUseCases WebAPIUseCases,
	exportUseCases ExportUseCases,
	deduplicateUseCases DeduplicateUseCases,
	taskUseCases TaskUseCases,
	debug bool,
	ac *apiservercore.Controller,
) *SystemHandlersController {
	c := &SystemHandlersController{
		logger:              logger,
		tracer:              tracer,
		parseUseCases:       parseUseCases,
		webAPIUseCases:      webAPIUseCases,
		exportUseCases:      exportUseCases,
		deduplicateUseCases: deduplicateUseCases,
		taskUseCases:        taskUseCases,
		debug:               debug,
		apiCore:             ac,
	}

	return c
}
