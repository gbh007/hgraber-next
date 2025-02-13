package attributeusecase

import (
	"context"
	"log/slog"

	"github.com/gbh007/hgraber-next/domain/core"
)

type storage interface {
	AttributesCount(ctx context.Context) ([]core.AttributeVariant, error)
	Attributes(ctx context.Context) ([]core.Attribute, error)

	InsertAttributeColor(ctx context.Context, color core.AttributeColor) error
	UpdateAttributeColor(ctx context.Context, color core.AttributeColor) error
	DeleteAttributeColor(ctx context.Context, code, value string) error
	AttributeColors(ctx context.Context) ([]core.AttributeColor, error)
	AttributeColor(ctx context.Context, code, value string) (core.AttributeColor, error)
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
