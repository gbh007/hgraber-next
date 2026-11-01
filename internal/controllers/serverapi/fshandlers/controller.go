package fshandlers

import (
	"context"
	"io"
	"log/slog"
	"net/url"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/internal/controllers/serverapi/apiservercore"
	"github.com/gbh007/hgraber-next/internal/domain/fsmodel"
)

type ParseUseCases interface {
	PageBodyByURL(ctx context.Context, u url.URL) (io.Reader, error)
}

type SystemUseCases interface {
	RemoveFilesInFSMismatch(ctx context.Context, fsID uuid.UUID) error
}

type FSUseCases interface {
	FileStoragesWithStatus(
		ctx context.Context,
		includeDBInfo, includeAvailableSizeInfo bool,
	) ([]fsmodel.FSWithStatus, error)
	FileStorage(ctx context.Context, id uuid.UUID) (fsmodel.FileStorageSystem, error)
	NewFileStorage(ctx context.Context, fs fsmodel.FileStorageSystem) (uuid.UUID, error)
	UpdateFileStorage(ctx context.Context, fs fsmodel.FileStorageSystem) error
	DeleteFileStorage(ctx context.Context, id uuid.UUID) error
	ValidateFS(ctx context.Context, fsID uuid.UUID) error
	TransferFSFiles(ctx context.Context, from, to uuid.UUID, onlyPreview bool) error
	TransferFSFilesByBook(ctx context.Context, bookID, to uuid.UUID, pageNumber *int) error

	File(ctx context.Context, fileID uuid.UUID, fsID *uuid.UUID) (io.Reader, error)
	PageBody(ctx context.Context, bookID uuid.UUID, pageNumber int) (io.Reader, error)
}

type FSHandlersController struct {
	logger *slog.Logger
	debug  bool

	apiCore *apiservercore.Controller

	parseUseCases  ParseUseCases
	systemUseCases SystemUseCases
	fsUseCases     FSUseCases
}

func New(
	logger *slog.Logger,
	parseUseCases ParseUseCases,
	systemUseCases SystemUseCases,
	fsUseCases FSUseCases,
	debug bool,
	ac *apiservercore.Controller,
) *FSHandlersController {
	c := &FSHandlersController{
		logger:         logger,
		parseUseCases:  parseUseCases,
		systemUseCases: systemUseCases,
		fsUseCases:     fsUseCases,
		debug:          debug,
		apiCore:        ac,
	}

	return c
}
