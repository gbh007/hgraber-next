package attributeusecase

import (
	"context"
	"fmt"
	"slices"

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

func (uc *UseCase) BookOriginAttributesCount(ctx context.Context) ([]core.AttributeVariant, error) {
	res, err := uc.storage.BookOriginAttributesCount(ctx)
	if err != nil {
		return nil, fmt.Errorf("storage: %w", err)
	}

	slices.SortFunc(res, func(a, b core.AttributeVariant) int {
		return b.Count - a.Count
	})

	return res, nil
}
