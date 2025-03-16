package attributeusecase

import (
	"context"
	"log/slog"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/core"
)

type storage interface {
	AttributesCount(ctx context.Context) ([]core.AttributeVariant, error)
	BookOriginAttributesCount(ctx context.Context) ([]core.AttributeVariant, error)

	Attributes(ctx context.Context) ([]core.Attribute, error)

	InsertAttributeColor(ctx context.Context, color core.AttributeColor) error
	UpdateAttributeColor(ctx context.Context, color core.AttributeColor) error
	DeleteAttributeColor(ctx context.Context, code, value string) error
	AttributeColors(ctx context.Context) ([]core.AttributeColor, error)
	AttributeColor(ctx context.Context, code, value string) (core.AttributeColor, error)

	InsertAttributeRemap(ctx context.Context, ar core.AttributeRemap) error
	UpdateAttributeRemap(ctx context.Context, ar core.AttributeRemap) error
	DeleteAttributeRemap(ctx context.Context, code, value string) error
	AttributeRemaps(ctx context.Context) ([]core.AttributeRemap, error)
	AttributeRemap(ctx context.Context, code, value string) (core.AttributeRemap, error)

	BookIDs(ctx context.Context, filter core.BookFilter) ([]uuid.UUID, error)
	BookOriginAttributes(ctx context.Context, bookID uuid.UUID) (map[string][]string, error)
	UpdateAttributes(ctx context.Context, bookID uuid.UUID, attributes map[string][]string) error
	DeleteBookAttributes(ctx context.Context, bookID uuid.UUID) error
}

type UseCase struct {
	logger *slog.Logger

	storage storage

	remapToLower bool
}

func New(
	logger *slog.Logger,
	storage storage,
	remapToLower bool,
) *UseCase {
	return &UseCase{
		logger:       logger,
		storage:      storage,
		remapToLower: remapToLower,
	}
}
