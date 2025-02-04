package webapi

import (
	"context"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (uc *UseCase) AttributesCount(ctx context.Context) ([]core.AttributeVariant, error) {
	res, err := uc.storage.AttributesCount(ctx)
	if err != nil {
		return nil, fmt.Errorf("storage: %w", err)
	}

	slices.SortFunc(res, func(a, b core.AttributeVariant) int {
		return b.Count - a.Count
	})

	return res, nil
}

func (uc *UseCase) CreateAttributeColor(ctx context.Context, color core.AttributeColor) error {
	color.CreatedAt = time.Now().UTC()

	return uc.storage.InsertAttributeColor(ctx, color)
}

func (uc *UseCase) UpdateAttributeColor(ctx context.Context, color core.AttributeColor) error {
	return uc.storage.UpdateAttributeColor(ctx, color)
}

func (uc *UseCase) DeleteAttributeColor(ctx context.Context, code, value string) error {
	return uc.storage.DeleteAttributeColor(ctx, code, value)
}

func (uc *UseCase) AttributeColors(ctx context.Context) ([]core.AttributeColor, error) {
	colors, err := uc.storage.AttributeColors(ctx)
	if err != nil {
		return nil, err
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
	return uc.storage.AttributeColor(ctx, code, value)
}
