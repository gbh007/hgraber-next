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

var errChainAttributeRemap = errors.New("chain attribute remap not supported")

func (uc *UseCase) CreateAttributeRemap(ctx context.Context, ar core.AttributeRemap) error {
	if !ar.IsDelete() && !ar.IsNoRemap() {
		targetAR, err := uc.AttributeRemap(ctx, ar.ToCode, ar.ToValue)
		if err != nil && !errors.Is(err, core.ErrAttributeRemapNotFound) {
			return fmt.Errorf("check existing attribute remap: %w", err)
		}

		if err == nil && !targetAR.IsNoRemap() {
			return errChainAttributeRemap
		}
	}

	ar.CreatedAt = time.Now().UTC()

	err := uc.storage.InsertAttributeRemap(ctx, ar)
	if err != nil {
		return fmt.Errorf("storage: insert attribute remap: %w", err)
	}

	return nil
}

func (uc *UseCase) UpdateAttributeRemap(ctx context.Context, ar core.AttributeRemap) error {
	if !ar.IsDelete() && !ar.IsNoRemap() {
		targetAR, err := uc.AttributeRemap(ctx, ar.ToCode, ar.ToValue)
		if err != nil && !errors.Is(err, core.ErrAttributeRemapNotFound) {
			return fmt.Errorf("check existing attribute remap: %w", err)
		}

		if err == nil && !targetAR.IsNoRemap() {
			return errChainAttributeRemap
		}
	}

	ar.UpdateAt = time.Now().UTC()

	err := uc.storage.UpdateAttributeRemap(ctx, ar)
	if err != nil {
		return fmt.Errorf("storage: update attribute remap: %w", err)
	}

	return nil
}

func (uc *UseCase) DeleteAttributeRemap(ctx context.Context, code, value string) error {
	return uc.storage.DeleteAttributeRemap(ctx, code, value) //nolint:wrapcheck // обвязка не требуется
}

func (uc *UseCase) AttributeRemaps(ctx context.Context) ([]core.AttributeRemap, error) {
	remaps, err := uc.storage.AttributeRemaps(ctx)
	if err != nil {
		return nil, fmt.Errorf("storage: get attribute remaps: %w", err)
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
	return uc.storage.AttributeRemap(ctx, code, value) //nolint:wrapcheck // обвязка не требуется
}
