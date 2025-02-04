package webapi

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/bff"
)

func (uc *UseCase) BookCompare(ctx context.Context, originID, targetID uuid.UUID) (bff.BookCompareResult, error) {
	pageCompare, err := uc.deduplicator.BookPagesCompare(ctx, originID, targetID)
	if err != nil {
		return bff.BookCompareResult{}, fmt.Errorf("deduplicator: page compare: %w", err)
	}

	attributeCompare, err := uc.deduplicator.BookAttributesCompare(ctx, originID, targetID, true) // Пока сравниваем на оригинальных атрибутах
	if err != nil {
		return bff.BookCompareResult{}, fmt.Errorf("deduplicator: attribute compare: %w", err)
	}

	attributesInfo, err := uc.storage.Attributes(ctx)
	if err != nil {
		return bff.BookCompareResult{}, fmt.Errorf("storage: get attributes info: %w", err)
	}

	attributesInfoMap := convertAttributes(attributesInfo)

	result := bff.BookCompareResult{
		BookPagesCompareResult: pageCompare,
		OriginAttributes:       convertBookAttributes(attributesInfoMap, attributeCompare.OriginAttributes),
		BothAttributes:         convertBookAttributes(attributesInfoMap, attributeCompare.BothAttributes),
		TargetAttributes:       convertBookAttributes(attributesInfoMap, attributeCompare.TargetAttributes),
	}

	return result, nil
}
