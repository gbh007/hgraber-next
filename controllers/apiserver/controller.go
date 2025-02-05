package apiserver

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/url"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/controllers/apiserver/agenthandlers"
	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/domain/bff"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

type ParseUseCases interface {
	NewBooks(ctx context.Context, urls []url.URL, autoVerify bool) (core.FirstHandleMultipleResult, error)

	BooksExists(ctx context.Context, urls []url.URL) ([]agentmodel.AgentBookCheckResult, error)
	PagesExists(ctx context.Context, urls []agentmodel.AgentPageURL) ([]agentmodel.AgentPageCheckResult, error)
	BookByURL(ctx context.Context, u url.URL) (core.BookContainer, error)
	PageBodyByURL(ctx context.Context, u url.URL) (io.Reader, error)

	NewBooksMulti(ctx context.Context, urls []url.URL, autoVerify bool) (core.MultiHandleMultipleResult, error)
}

type WebAPIUseCases interface {
	SystemSize(ctx context.Context) (core.SystemSizeInfo, error)
	WorkersInfo(ctx context.Context) []core.SystemWorkerStat

	File(ctx context.Context, fileID uuid.UUID, fsID *uuid.UUID) (io.Reader, error)
	PageBody(ctx context.Context, bookID uuid.UUID, pageNumber int) (io.Reader, error)

	BookRaw(ctx context.Context, bookID uuid.UUID) (core.BookContainer, error)

	VerifyBook(ctx context.Context, bookID uuid.UUID, verified bool) error
	DeleteBook(ctx context.Context, bookID uuid.UUID) error

	SetWorkerConfig(ctx context.Context, counts map[string]int)

	AttributesCount(ctx context.Context) ([]core.AttributeVariant, error)
	CreateAttributeColor(ctx context.Context, color core.AttributeColor) error
	UpdateAttributeColor(ctx context.Context, color core.AttributeColor) error
	DeleteAttributeColor(ctx context.Context, code, value string) error
	AttributeColors(ctx context.Context) ([]core.AttributeColor, error)
	AttributeColor(ctx context.Context, code, value string) (core.AttributeColor, error)

	SetLabel(ctx context.Context, label core.BookLabel) error
	DeleteLabel(ctx context.Context, label core.BookLabel) error
	Labels(ctx context.Context, bookID uuid.UUID) ([]core.BookLabel, error)
	CreateLabelPreset(ctx context.Context, preset core.BookLabelPreset) error
	UpdateLabelPreset(ctx context.Context, preset core.BookLabelPreset) error
	DeleteLabelPreset(ctx context.Context, name string) error
	LabelPresets(ctx context.Context) ([]core.BookLabelPreset, error)
	LabelPreset(ctx context.Context, name string) (core.BookLabelPreset, error)

	BookCompare(ctx context.Context, originID, targetID uuid.UUID) (bff.BookCompareResult, error)
}

type AgentUseCases interface {
	NewAgent(ctx context.Context, agent core.Agent) error
	DeleteAgent(ctx context.Context, id uuid.UUID) error
	Agents(ctx context.Context, filter core.AgentFilter, includeStatus bool) ([]core.AgentWithStatus, error)
	UpdateAgent(ctx context.Context, agent core.Agent) error
	Agent(ctx context.Context, id uuid.UUID) (core.Agent, error)
}

type ExportUseCases interface {
	Export(ctx context.Context, agentID uuid.UUID, filter core.BookFilter, deleteAfter bool) error
	ExportBook(ctx context.Context, bookID uuid.UUID) (io.Reader, core.BookContainer, error)
	ImportArchive(ctx context.Context, body io.Reader, deduplicate bool, autoVerify bool) (uuid.UUID, error)
}

type DeduplicateUseCases interface {
	ArchiveEntryPercentage(ctx context.Context, archiveBody io.Reader) ([]core.DeduplicateArchiveResult, error)
	BookByPageEntryPercentage(ctx context.Context, originBookID uuid.UUID) ([]bff.DeduplicateBookResult, error)
	UniquePages(ctx context.Context, originBookID uuid.UUID) ([]bff.PreviewPage, error)
	BooksByPage(ctx context.Context, bookID uuid.UUID, pageNumber int) ([]bff.BookWithPreviewPage, error)

	CreateDeadHashByPage(ctx context.Context, bookID uuid.UUID, pageNumber int) error
	DeleteDeadHashByPage(ctx context.Context, bookID uuid.UUID, pageNumber int) error
	DeleteAllPageByHash(ctx context.Context, bookID uuid.UUID, pageNumber int, setDeadHash bool) error

	MarkBookPagesAsDeadHash(ctx context.Context, bookID uuid.UUID) error
	UnMarkBookPagesAsDeadHash(ctx context.Context, bookID uuid.UUID) error
	RemoveBookPagesWithDeadHash(ctx context.Context, bookID uuid.UUID, deleteEmptyBook bool) error
	DeleteBookDeadHashedPages(ctx context.Context, bookID uuid.UUID) error
}

type TaskUseCases interface {
	RunTask(ctx context.Context, code core.TaskCode) error
	TaskResults(ctx context.Context) ([]*core.TaskResult, error)
	RemoveFilesInFSMismatch(ctx context.Context, fsID uuid.UUID) error
}

type ReBuilderUseCases interface {
	UpdateBook(ctx context.Context, book core.BookContainer) error
	RebuildBook(ctx context.Context, request core.RebuildBookRequest) (uuid.UUID, error)
	RestoreBook(ctx context.Context, bookID uuid.UUID, onlyPages bool) error
}

type FSUseCases interface {
	FileStoragesWithStatus(ctx context.Context, includeDBInfo, includeAvailableSizeInfo bool) ([]core.FSWithStatus, error)
	FileStorage(ctx context.Context, id uuid.UUID) (core.FileStorageSystem, error)
	NewFileStorage(ctx context.Context, fs core.FileStorageSystem) (uuid.UUID, error)
	UpdateFileStorage(ctx context.Context, fs core.FileStorageSystem) error
	DeleteFileStorage(ctx context.Context, id uuid.UUID) error
	ValidateFS(ctx context.Context, fsID uuid.UUID) error
	TransferFSFiles(ctx context.Context, from, to uuid.UUID, onlyPreview bool) error
	TransferFSFilesByBook(ctx context.Context, bookID, to uuid.UUID, pageNumber *int) error

	HighwayFileURL(ctx context.Context, fileID uuid.UUID, ext string, fsID uuid.UUID) (url.URL, bool, error)
}

type BFFUseCases interface {
	BookDetails(ctx context.Context, bookID uuid.UUID) (bff.BookDetails, error)
	BookList(ctx context.Context, filter core.BookFilter) (bff.BookList, error)
}

type config interface {
	GetAddr() string
	GetExternalAddr() string
	GetStaticDir() string
	GetToken() string
}

type Controller struct {
	*agenthandlers.AgentHandlersController

	logger    *slog.Logger
	tracer    trace.Tracer
	debug     bool
	staticDir string

	apiCore *apiservercore.Controller

	parseUseCases       ParseUseCases
	webAPIUseCases      WebAPIUseCases
	agentUseCases       AgentUseCases
	exportUseCases      ExportUseCases
	deduplicateUseCases DeduplicateUseCases
	taskUseCases        TaskUseCases
	rebuilderUseCases   ReBuilderUseCases
	fsUseCases          FSUseCases
	bffUseCases         BFFUseCases

	ogenServer *serverAPI.Server

	serverAddr string

	token string
}

func New(
	logger *slog.Logger,
	tracer trace.Tracer,
	config config,
	parseUseCases ParseUseCases,
	webAPIUseCases WebAPIUseCases,
	agentUseCases AgentUseCases,
	exportUseCases ExportUseCases,
	deduplicateUseCases DeduplicateUseCases,
	taskUseCases TaskUseCases,
	rebuilderUseCases ReBuilderUseCases,
	fsUseCases FSUseCases,
	bffUseCases BFFUseCases,
	debug bool,
) (*Controller, error) {
	ac, err := apiservercore.New(
		logger,
		tracer,
		config,
		fsUseCases,
		debug,
	)
	if err != nil {
		return nil, fmt.Errorf("init core handlers: %w", err)
	}

	c := &Controller{
		AgentHandlersController: agenthandlers.New(
			logger,
			tracer,
			agentUseCases,
			exportUseCases,
			debug,
			ac,
		),

		logger:              logger,
		tracer:              tracer,
		serverAddr:          config.GetAddr(),
		parseUseCases:       parseUseCases,
		webAPIUseCases:      webAPIUseCases,
		agentUseCases:       agentUseCases,
		exportUseCases:      exportUseCases,
		deduplicateUseCases: deduplicateUseCases,
		taskUseCases:        taskUseCases,
		rebuilderUseCases:   rebuilderUseCases,
		fsUseCases:          fsUseCases,
		bffUseCases:         bffUseCases,
		debug:               debug,
		staticDir:           config.GetStaticDir(),
		token:               config.GetToken(),

		apiCore: ac,
	}

	ogenServer, err := serverAPI.NewServer(
		c, c,
		serverAPI.WithErrorHandler(methodErrorHandler),
		serverAPI.WithMethodNotAllowed(methodNotAllowed),
		serverAPI.WithNotFound(methodNotFound),
		serverAPI.WithMiddleware(c.simplePanicRecover),
	)
	if err != nil {
		return nil, fmt.Errorf("create ogen server: %w", err)
	}

	c.ogenServer = ogenServer

	return c, nil
}
