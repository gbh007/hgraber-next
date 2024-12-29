package webapi

import (
	"context"
	"fmt"
	"slices"

	"hgnext/internal/entities"
)

func (uc *UseCase) AttributesCount(ctx context.Context) ([]entities.AttributeVariant, error) {
	res, err := uc.storage.AttributesCount(ctx)
	if err != nil {
		return nil, fmt.Errorf("storage: %w", err)
	}

	slices.SortFunc(res, func(a, b entities.AttributeVariant) int {
		return b.Count - a.Count
	})

	return res, nil
}
