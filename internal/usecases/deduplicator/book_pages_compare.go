package deduplicator

import (
	"context"
	"fmt"
	"slices"

	"github.com/google/uuid"

	"hgnext/internal/entities"
)

func (uc *UseCase) BookPagesCompare(ctx context.Context, originID, targetID uuid.UUID) (entities.BookPagesCompareResult, error) {
	originBook, err := uc.storage.GetBook(ctx, originID)
	if err != nil {
		return entities.BookPagesCompareResult{}, fmt.Errorf("get book (%s) from storage: %w", originID.String(), err)
	}

	targetBook, err := uc.storage.GetBook(ctx, targetID)
	if err != nil {
		return entities.BookPagesCompareResult{}, fmt.Errorf("get book (%s) from storage: %w", targetID.String(), err)
	}

	originPages, err := uc.storage.BookPagesWithHash(ctx, originID)
	if err != nil {
		return entities.BookPagesCompareResult{}, fmt.Errorf("get pages (%s) from storage: %w", originID.String(), err)
	}

	targetPages, err := uc.storage.BookPagesWithHash(ctx, targetID)
	if err != nil {
		return entities.BookPagesCompareResult{}, fmt.Errorf("get pages (%s) from storage: %w", targetID.String(), err)
	}

	result := entities.BookPagesCompareResult{
		OriginBook:  originBook,
		TargetBook:  targetBook,
		OriginPages: make([]entities.Page, 0, len(originPages)),
		BothPages:   make([]entities.Page, 0, max(len(originPages), len(targetPages))),
		TargetPages: make([]entities.Page, 0, len(targetPages)),
	}

	hashes := make(map[entities.FileHash]int, len(originPages))

	for _, page := range originPages {
		if page.PageNumber == 1 && page.Downloaded {
			result.OriginPreviewPage = page.Page()
		}

		hashes[page.Hash()] = 1 // Специальная логика, т.к. в книге могут быть дубликаты страниц
	}

	for _, page := range targetPages {
		if page.PageNumber == 1 && page.Downloaded {
			result.TargetPreviewPage = page.Page()
		}

		if hashes[page.Hash()] == 1 { // Специальная логика, т.к. в книге могут быть дубликаты страниц
			hashes[page.Hash()] = 2
		}
	}

	for _, page := range originPages {
		if hashes[page.Hash()] == 1 {
			result.OriginPages = append(result.OriginPages, page.Page())
		} else {
			result.BothPages = append(result.BothPages, page.Page()) // Приоритет отдаем оригинальной книге
		}
	}

	for _, page := range targetPages {
		if hashes[page.Hash()] == 1 {
			result.TargetPages = append(result.TargetPages, page.Page())
		}
	}

	slices.SortStableFunc(result.OriginPages, func(a, b entities.Page) int {
		return a.PageNumber - b.PageNumber
	})
	slices.SortStableFunc(result.BothPages, func(a, b entities.Page) int {
		return a.PageNumber - b.PageNumber
	})
	slices.SortStableFunc(result.TargetPages, func(a, b entities.Page) int {
		return a.PageNumber - b.PageNumber
	})

	return result, nil
}
