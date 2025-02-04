package webapi

import (
	"context"
	"io"
	"log/slog"
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/entities"
)

type storage interface {
	SystemSize(ctx context.Context) (entities.SystemSizeInfo, error)
	GetPage(ctx context.Context, id uuid.UUID, pageNumber int) (entities.Page, error)

	VerifyBook(ctx context.Context, bookID uuid.UUID, verified bool, verifiedAt time.Time) error
	MarkBookAsDeleted(ctx context.Context, bookID uuid.UUID) error

	AttributesCount(ctx context.Context) ([]entities.AttributeVariant, error)
	Attributes(ctx context.Context) ([]entities.Attribute, error)

	SetLabel(ctx context.Context, label entities.BookLabel) error
	DeleteLabel(ctx context.Context, label entities.BookLabel) error
	Labels(ctx context.Context, bookID uuid.UUID) ([]entities.BookLabel, error)
	InsertLabelPreset(ctx context.Context, preset entities.BookLabelPreset) error
	UpdateLabelPreset(ctx context.Context, preset entities.BookLabelPreset) error
	DeleteLabelPreset(ctx context.Context, name string) error
	LabelPresets(ctx context.Context) ([]entities.BookLabelPreset, error)
	LabelPreset(ctx context.Context, name string) (entities.BookLabelPreset, error)

	InsertAttributeColor(ctx context.Context, color entities.AttributeColor) error
	UpdateAttributeColor(ctx context.Context, color entities.AttributeColor) error
	DeleteAttributeColor(ctx context.Context, code, value string) error
	AttributeColors(ctx context.Context) ([]entities.AttributeColor, error)
	AttributeColor(ctx context.Context, code, value string) (entities.AttributeColor, error)
}

type bookRequester interface {
	BookOriginFull(ctx context.Context, bookID uuid.UUID) (entities.BookContainer, error)
}

type workerManager interface {
	Info() []entities.SystemWorkerStat

	SetRunnerCount(ctx context.Context, counts map[string]int)
}

type fileStorage interface {
	Get(ctx context.Context, fileID uuid.UUID, fsID *uuid.UUID) (io.Reader, error)
}

type deduplicator interface {
	BookAttributesCompare(ctx context.Context, originID, targetID uuid.UUID, useOrigin bool) (entities.BookAttributesCompareResult, error)
	BookPagesCompare(ctx context.Context, originID, targetID uuid.UUID) (entities.BookPagesCompareResult, error)
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
