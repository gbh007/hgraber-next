package fshandlers

import (
	"context"
	"io"
	"log/slog"
	"net/url"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
)

type ParseUseCases interface {
	PageBodyByURL(ctx context.Context, u url.URL) (io.Reader, error)
}

type WebAPIUseCases interface {
	File(ctx context.Context, fileID uuid.UUID, fsID *uuid.UUID) (io.Reader, error)
	PageBody(ctx context.Context, bookID uuid.UUID, pageNumber int) (io.Reader, error)
}

type TaskUseCases interface {
	RemoveFilesInFSMismatch(ctx context.Context, fsID uuid.UUID) error
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
}

type FSHandlersController struct {
	logger *slog.Logger
	tracer trace.Tracer
	debug  bool

	apiCore *apiservercore.Controller

	parseUseCases  ParseUseCases
	webAPIUseCases WebAPIUseCases
	taskUseCases   TaskUseCases
	fsUseCases     FSUseCases
}

func New(
	logger *slog.Logger,
	tracer trace.Tracer,
	parseUseCases ParseUseCases,
	webAPIUseCases WebAPIUseCases,
	taskUseCases TaskUseCases,
	fsUseCases FSUseCases,
	debug bool,
	ac *apiservercore.Controller,
) *FSHandlersController {
	c := &FSHandlersController{
		logger:         logger,
		tracer:         tracer,
		parseUseCases:  parseUseCases,
		webAPIUseCases: webAPIUseCases,
		taskUseCases:   taskUseCases,
		fsUseCases:     fsUseCases,
		debug:          debug,
		apiCore:        ac,
	}

	return c
}
