package deduplicator

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/internal/entities"
)

func (uc *UseCase) BookAttributesCompare(ctx context.Context, originID, targetID uuid.UUID, useOrigin bool) (entities.BookAttributesCompareResult, error) {
	var (
		originBookAttributes map[string][]string
		targetBookAttributes map[string][]string
	)

	allAttributesInfo, err := uc.storage.Attributes(ctx)
	if err != nil {
		return entities.BookAttributesCompareResult{}, fmt.Errorf("get attributes info from storage: %w", err)
	}

	if useOrigin {
		originBookAttributes, err = uc.storage.BookOriginAttributes(ctx, originID)
		if err != nil {
			return entities.BookAttributesCompareResult{}, fmt.Errorf("get origin attributes (%s) from storage: %w", originID.String(), err)
		}

		targetBookAttributes, err = uc.storage.BookOriginAttributes(ctx, targetID)
		if err != nil {
			return entities.BookAttributesCompareResult{}, fmt.Errorf("get origin attributes (%s) from storage: %w", targetID.String(), err)
		}
	} else {
		originBookAttributes, err = uc.storage.BookAttributes(ctx, originID)
		if err != nil {
			return entities.BookAttributesCompareResult{}, fmt.Errorf("get attributes (%s) from storage: %w", originID.String(), err)
		}

		targetBookAttributes, err = uc.storage.BookAttributes(ctx, targetID)
		if err != nil {
			return entities.BookAttributesCompareResult{}, fmt.Errorf("get attributes (%s) from storage: %w", targetID.String(), err)
		}
	}

	result := entities.BookAttributesCompareResult{
		OriginAttributes: make(map[string][]string, entities.PossibleAttributeCount),
		BothAttributes:   make(map[string][]string, entities.PossibleAttributeCount),
		TargetAttributes: make(map[string][]string, entities.PossibleAttributeCount),
	}

	for _, attr := range allAttributesInfo {
		originValues, bothValues, targetValues := entities.AttributesValuesDiff(
			originBookAttributes[attr.Code],
			targetBookAttributes[attr.Code],
		)

		if len(originValues) > 0 {
			result.OriginAttributes[attr.Code] = originValues
		}

		if len(bothValues) > 0 {
			result.BothAttributes[attr.Code] = bothValues
		}

		if len(targetValues) > 0 {
			result.TargetAttributes[attr.Code] = targetValues
		}
	}

	return result, nil
}
