package agentimportdeduplicatoruc

import (
	"context"
	"io"
	"log/slog"

	"github.com/gbh007/hgraber-next/internal/domain/agentdto"
)

type importFS interface {
	AllZips(ctx context.Context) ([]string, error)
	Get(ctx context.Context, relativePath string) (io.Reader, error)
}

type storage interface {
	CreateImport(ctx context.Context, info agentdto.ImportInfo) error
	CreateMissing(ctx context.Context, path string, maxEntryPercentage float64) error

	ImportedCountByRelativePath(ctx context.Context, path string) (int, error)
	TruncateMissing(ctx context.Context) error
}

type masterAPI interface {
	DeduplicateArchive(ctx context.Context, body io.Reader) ([]agentdto.DeduplicateArchiveResult, error)
}

type UseCase struct {
	logger *slog.Logger

	importFS  importFS
	storage   storage
	masterAPI masterAPI
}

func New(
	logger *slog.Logger,
	importFS importFS,
	storage storage,
	masterAPI masterAPI,
) *UseCase {
	return &UseCase{
		logger:    logger,
		importFS:  importFS,
		storage:   storage,
		masterAPI: masterAPI,
	}
}
