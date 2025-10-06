package attributeusecase

import (
	"context"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (uc *UseCase) CreateAttributeColor(ctx context.Context, color core.AttributeColor) error {
	color.CreatedAt = time.Now().UTC()

	return uc.storage.InsertAttributeColor(ctx, color) //nolint:wrapcheck // обвязка не требуется
}

func (uc *UseCase) UpdateAttributeColor(ctx context.Context, color core.AttributeColor) error {
	return uc.storage.UpdateAttributeColor(ctx, color) //nolint:wrapcheck // обвязка не требуется
}

func (uc *UseCase) DeleteAttributeColor(ctx context.Context, code, value string) error {
	return uc.storage.DeleteAttributeColor(ctx, code, value) //nolint:wrapcheck // обвязка не требуется
}

func (uc *UseCase) AttributeColors(ctx context.Context) ([]core.AttributeColor, error) {
	colors, err := uc.storage.AttributeColors(ctx)
	if err != nil {
		return nil, fmt.Errorf("storage: get attribute colors: %w", err)
	}

	slices.SortStableFunc(colors, func(a, b core.AttributeColor) int {
		if a.Code != b.Code {
			return strings.Compare(a.Code, b.Code)
		}

		if a.Value != b.Value {
			return strings.Compare(a.Value, b.Value)
		}

		return a.CreatedAt.Compare(b.CreatedAt)
	})

	return colors, nil
}

func (uc *UseCase) AttributeColor(ctx context.Context, code, value string) (core.AttributeColor, error) {
	return uc.storage.AttributeColor(ctx, code, value) //nolint:wrapcheck // обвязка не требуется
}
