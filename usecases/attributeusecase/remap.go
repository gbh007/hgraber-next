package attributeusecase

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (uc *UseCase) CreateAttributeRemap(ctx context.Context, ar core.AttributeRemap) error {
	if !(ar.IsDelete() || ar.IsNoRemap()) {
		targetAR, err := uc.AttributeRemap(ctx, ar.ToCode, ar.ToValue)
		if err != nil && !errors.Is(err, core.AttributeRemapNotFoundError) {
			return fmt.Errorf("check existing attribute remap: %w", err)
		}

		if err == nil && !targetAR.IsNoRemap() {
			return fmt.Errorf("chain attribute remap not supported")
		}
	}

	ar.CreatedAt = time.Now().UTC()

	return uc.storage.InsertAttributeRemap(ctx, ar)
}

func (uc *UseCase) UpdateAttributeRemap(ctx context.Context, ar core.AttributeRemap) error {
	if !(ar.IsDelete() || ar.IsNoRemap()) {
		targetAR, err := uc.AttributeRemap(ctx, ar.ToCode, ar.ToValue)
		if err != nil && !errors.Is(err, core.AttributeRemapNotFoundError) {
			return fmt.Errorf("check existing attribute remap: %w", err)
		}

		if err == nil && !targetAR.IsNoRemap() {
			return fmt.Errorf("chain attribute remap not supported")
		}
	}

	ar.UpdateAt = time.Now().UTC()

	return uc.storage.UpdateAttributeRemap(ctx, ar)
}

func (uc *UseCase) DeleteAttributeRemap(ctx context.Context, code, value string) error {
	return uc.storage.DeleteAttributeRemap(ctx, code, value)
}

func (uc *UseCase) AttributeRemaps(ctx context.Context) ([]core.AttributeRemap, error) {
	remaps, err := uc.storage.AttributeRemaps(ctx)
	if err != nil {
		return nil, err
	}

	slices.SortStableFunc(remaps, func(a, b core.AttributeRemap) int {
		if a.Code != b.Code {
			return strings.Compare(a.Code, b.Code)
		}

		if a.Value != b.Value {
			return strings.Compare(a.Value, b.Value)
		}

		return a.CreatedAt.Compare(b.CreatedAt)
	})

	return remaps, nil
}

func (uc *UseCase) AttributeRemap(ctx context.Context, code, value string) (core.AttributeRemap, error) {
	return uc.storage.AttributeRemap(ctx, code, value)
}
