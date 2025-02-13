package labelhandler

import (
	"context"
	"log/slog"
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/domain/systemmodel"
)

type storage interface {
	SystemSize(ctx context.Context) (systemmodel.SystemSizeInfo, error)
	GetPage(ctx context.Context, id uuid.UUID, pageNumber int) (core.Page, error)

	VerifyBook(ctx context.Context, bookID uuid.UUID, verified bool, verifiedAt time.Time) error
	MarkBookAsDeleted(ctx context.Context, bookID uuid.UUID) error

	SetLabel(ctx context.Context, label core.BookLabel) error
	DeleteLabel(ctx context.Context, label core.BookLabel) error
	Labels(ctx context.Context, bookID uuid.UUID) ([]core.BookLabel, error)
	InsertLabelPreset(ctx context.Context, preset core.BookLabelPreset) error
	UpdateLabelPreset(ctx context.Context, preset core.BookLabelPreset) error
	DeleteLabelPreset(ctx context.Context, name string) error
	LabelPresets(ctx context.Context) ([]core.BookLabelPreset, error)
	LabelPreset(ctx context.Context, name string) (core.BookLabelPreset, error)
}

type UseCase struct {
	logger *slog.Logger

	storage storage
}

func New(
	logger *slog.Logger,
	storage storage,
) *UseCase {
	return &UseCase{
		logger:  logger,
		storage: storage,
	}
}
