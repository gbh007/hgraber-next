package attributeusecase

import (
	"context"
	"slices"
	"strings"
	"time"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (uc *UseCase) CreateAttributeRemap(ctx context.Context, color core.AttributeRemap) error {
	color.CreatedAt = time.Now().UTC()

	return uc.storage.InsertAttributeRemap(ctx, color)
}

func (uc *UseCase) UpdateAttributeRemap(ctx context.Context, color core.AttributeRemap) error {
	color.UpdateAt = time.Now().UTC()

	return uc.storage.UpdateAttributeRemap(ctx, color)
}

func (uc *UseCase) DeleteAttributeRemap(ctx context.Context, code, value string) error {
	return uc.storage.DeleteAttributeRemap(ctx, code, value)
}

func (uc *UseCase) AttributeRemaps(ctx context.Context) ([]core.AttributeRemap, error) {
	colors, err := uc.storage.AttributeRemaps(ctx)
	if err != nil {
		return nil, err
	}

	slices.SortStableFunc(colors, func(a, b core.AttributeRemap) int {
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

func (uc *UseCase) AttributeRemap(ctx context.Context, code, value string) (core.AttributeRemap, error) {
	return uc.storage.AttributeRemap(ctx, code, value)
}
