package exportusecase

import (
	"context"
	"io"
	"log/slog"
	"net/url"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/domain/parsing"
)

type fileStorage interface {
	Get(ctx context.Context, fileID uuid.UUID, fsID *uuid.UUID) (io.Reader, error)
	Create(ctx context.Context, fileID uuid.UUID, body io.Reader, fsID uuid.UUID) error
	FSIDForDownload(ctx context.Context) (uuid.UUID, error)
}

type storage interface {
	NewBook(ctx context.Context, book core.Book) error
	UpdateBookPages(ctx context.Context, id uuid.UUID, pages []core.Page) error
	SetLabel(ctx context.Context, label core.BookLabel) error
	UpdateAttributes(ctx context.Context, bookID uuid.UUID, attributes map[string][]string) error
	UpdateOriginAttributes(ctx context.Context, bookID uuid.UUID, attributes map[string][]string) error
	NewFile(ctx context.Context, file core.File) error

	DeleteBook(ctx context.Context, id uuid.UUID) error

	GetBookIDsByURL(ctx context.Context, u url.URL) ([]uuid.UUID, error)

	MarkBookAsDeleted(ctx context.Context, bookID uuid.UUID) error

	Mirrors(ctx context.Context) ([]parsing.URLMirror, error)

	BookIDs(ctx context.Context, filter core.BookFilter) ([]uuid.UUID, error)

	AttributeRemaps(ctx context.Context) ([]core.AttributeRemap, error)
}

type agentSystem interface {
	ExportArchive(ctx context.Context, agentID uuid.UUID, data agentmodel.AgentExportData) error
}

type tmpStorage interface {
	AddToExport(books []agentmodel.BookToExport)
	ExportList() []agentmodel.BookToExport
}

type bookAdapter interface {
	BookRaw(ctx context.Context, bookID uuid.UUID) (core.BookContainer, error)
}

type UseCase struct {
	logger *slog.Logger

	storage     storage
	fileStorage fileStorage
	agentSystem agentSystem
	tmpStorage  tmpStorage
	bookAdapter bookAdapter

	autoRemap    bool
	remapToLower bool
}

func New(
	logger *slog.Logger,
	storage storage,
	fileStorage fileStorage,
	agentSystem agentSystem,
	tmpStorage tmpStorage,
	bookAdapter bookAdapter,
	autoRemap bool,
	remapToLower bool,
) *UseCase {
	return &UseCase{
		logger:       logger,
		storage:      storage,
		fileStorage:  fileStorage,
		agentSystem:  agentSystem,
		tmpStorage:   tmpStorage,
		bookAdapter:  bookAdapter,
		autoRemap:    autoRemap,
		remapToLower: remapToLower,
	}
}
