package webapi

import (
	"context"
	"io"
	"log/slog"
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/bff"
	"github.com/gbh007/hgraber-next/domain/core"
)

type storage interface {
	SystemSize(ctx context.Context) (core.SystemSizeInfo, error)
	GetPage(ctx context.Context, id uuid.UUID, pageNumber int) (core.Page, error)

	VerifyBook(ctx context.Context, bookID uuid.UUID, verified bool, verifiedAt time.Time) error
	MarkBookAsDeleted(ctx context.Context, bookID uuid.UUID) error

	AttributesCount(ctx context.Context) ([]core.AttributeVariant, error)
	Attributes(ctx context.Context) ([]core.Attribute, error)

	SetLabel(ctx context.Context, label core.BookLabel) error
	DeleteLabel(ctx context.Context, label core.BookLabel) error
	Labels(ctx context.Context, bookID uuid.UUID) ([]core.BookLabel, error)
	InsertLabelPreset(ctx context.Context, preset core.BookLabelPreset) error
	UpdateLabelPreset(ctx context.Context, preset core.BookLabelPreset) error
	DeleteLabelPreset(ctx context.Context, name string) error
	LabelPresets(ctx context.Context) ([]core.BookLabelPreset, error)
	LabelPreset(ctx context.Context, name string) (core.BookLabelPreset, error)

	InsertAttributeColor(ctx context.Context, color core.AttributeColor) error
	UpdateAttributeColor(ctx context.Context, color core.AttributeColor) error
	DeleteAttributeColor(ctx context.Context, code, value string) error
	AttributeColors(ctx context.Context) ([]core.AttributeColor, error)
	AttributeColor(ctx context.Context, code, value string) (core.AttributeColor, error)
}

type bookRequester interface {
	BookOriginFull(ctx context.Context, bookID uuid.UUID) (core.BookContainer, error)
}

type workerManager interface {
	Info() []core.SystemWorkerStat

	SetRunnerCount(ctx context.Context, counts map[string]int)
}

type fileStorage interface {
	Get(ctx context.Context, fileID uuid.UUID, fsID *uuid.UUID) (io.Reader, error)
}

type deduplicator interface {
	BookAttributesCompare(ctx context.Context, originID, targetID uuid.UUID, useOrigin bool) (bff.BookAttributesCompareResult, error)
	BookPagesCompare(ctx context.Context, originID, targetID uuid.UUID) (bff.BookPagesCompareResult, error)
}

type UseCase struct {
	logger *slog.Logger

	workerManager workerManager
	storage       storage
	fileStorage   fileStorage
	bookRequester bookRequester
	deduplicator  deduplicator
}

func New(
	logger *slog.Logger,
	workerManager workerManager,
	storage storage,
	fileStorage fileStorage,
	bookRequester bookRequester,
	deduplicator deduplicator,
) *UseCase {
	return &UseCase{
		logger:        logger,
		workerManager: workerManager,
		storage:       storage,
		fileStorage:   fileStorage,
		bookRequester: bookRequester,
		deduplicator:  deduplicator,
	}
}
