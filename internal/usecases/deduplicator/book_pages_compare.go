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
		OriginPages: make([]entities.PageWithDeadHash, 0, len(originPages)),
		BothPages:   make([]entities.PageWithDeadHash, 0, max(len(originPages), len(targetPages))),
		TargetPages: make([]entities.PageWithDeadHash, 0, len(targetPages)),
	}

	hashes := make(map[entities.FileHash]int, len(originPages))

	md5Sums := make([]string, 0, len(originPages)+len(targetPages))

	for _, page := range originPages {
		hashes[page.FileHash] = 1 // Специальная логика, т.к. в книге могут быть дубликаты страниц

		md5Sums = append(md5Sums, page.Md5Sum)
	}

	for _, page := range targetPages {
		if hashes[page.FileHash] == 1 { // Специальная логика, т.к. в книге могут быть дубликаты страниц
			hashes[page.FileHash] = 2
		}

		md5Sums = append(md5Sums, page.Md5Sum)
	}

	md5Sums = slices.Compact(md5Sums)

	deadHashes, err := uc.storage.DeadHashesByMD5Sums(ctx, md5Sums)
	if err != nil {
		return entities.BookPagesCompareResult{}, fmt.Errorf("storage: get dead hashes: %w", err)
	}

	existsDeadHashes := make(map[entities.FileHash]struct{}, len(deadHashes))

	for _, hash := range deadHashes {
		existsDeadHashes[hash.FileHash] = struct{}{}
	}

	result.EntryPercentage = entities.EntryPercentageForPages(originPages, targetPages, nil)
	result.ReverseEntryPercentage = entities.EntryPercentageForPages(targetPages, originPages, nil)

	result.EntryPercentageWithoutDeadHashes = entities.EntryPercentageForPages(originPages, targetPages, existsDeadHashes)
	result.ReverseEntryPercentageWithoutDeadHashes = entities.EntryPercentageForPages(targetPages, originPages, existsDeadHashes)

	for _, page := range originPages {
		_, hasDeadHash := existsDeadHashes[page.FileHash]

		if page.PageNumber == entities.PageNumberForPreview {
			result.OriginPreviewPage = page.ToPreview()
			result.OriginPreviewPage.HasDeadHash = &hasDeadHash
		}

		if hashes[page.FileHash] == 1 {
			result.OriginPages = append(result.OriginPages, entities.PageWithDeadHash{
				Page:        page.Page,
				FSID:        &page.FSID,
				HasDeadHash: hasDeadHash,
			})
		} else {
			result.BothPages = append(result.BothPages, entities.PageWithDeadHash{
				Page:        page.Page,
				FSID:        &page.FSID,
				HasDeadHash: hasDeadHash,
			}) // Приоритет отдаем оригинальной книге
		}
	}

	for _, page := range targetPages {
		_, hasDeadHash := existsDeadHashes[page.FileHash]

		if page.PageNumber == entities.PageNumberForPreview {
			result.TargetPreviewPage = page.ToPreview()
			result.TargetPreviewPage.HasDeadHash = &hasDeadHash
		}

		if hashes[page.FileHash] == 0 {
			result.TargetPages = append(result.TargetPages, entities.PageWithDeadHash{
				Page:        page.Page,
				FSID:        &page.FSID,
				HasDeadHash: hasDeadHash,
			})
		}
	}

	slices.SortStableFunc(result.OriginPages, func(a, b entities.PageWithDeadHash) int {
		return a.PageNumber - b.PageNumber
	})
	slices.SortStableFunc(result.BothPages, func(a, b entities.PageWithDeadHash) int {
		return a.PageNumber - b.PageNumber
	})
	slices.SortStableFunc(result.TargetPages, func(a, b entities.PageWithDeadHash) int {
		return a.PageNumber - b.PageNumber
	})

	return result, nil
}
