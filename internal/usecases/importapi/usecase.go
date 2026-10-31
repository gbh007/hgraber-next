package importapi

import (
	"context"
	"log/slog"
	"net/url"

	"github.com/gbh007/hgraber-next/internal/domain/entities"
	"github.com/google/uuid"
)

type storage interface {
	CreateImport(ctx context.Context, info entities.ImportInfo) error
	ImportedCountByID(ctx context.Context, bookID uuid.UUID) (int, error)
	ImportedCountByURL(ctx context.Context, u url.URL) (int, error)
}

type importFS interface {
	CreateImport(ctx context.Context, data entities.ImportData) (string, error)
}

type UseCase struct {
	logger *slog.Logger

	storage  storage
	importFS importFS
}

func New(
	logger *slog.Logger,
	storage storage,
	importFS importFS,
) *UseCase {
	return &UseCase{
		logger:   logger,
		storage:  storage,
		importFS: importFS,
	}
}
