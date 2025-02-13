package rebuilderusecase

import (
	"context"
	"fmt"
	"slices"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/core"
)

type rebuildPageResources struct {
	SourceBook         core.Book
	SourcePagesMap     map[int]core.PageWithHash
	ForbiddenHashes    map[core.FileHash]struct{}
	UnusedSourceHashes map[core.FileHash]struct{}
}

type rebuildedPagesInfo struct {
	PagesRemap         map[int]int
	SourcePageNumbers  []int
	NewPages           []core.Page
	UnusedSourceHashes map[core.FileHash]struct{}
}

func (uc *UseCase) RebuildBook(ctx context.Context, request core.RebuildBookRequest) (_ uuid.UUID, returnErr error) {
	if len(request.SelectedPages) == 0 {
		return uuid.Nil, fmt.Errorf("%w: %w", core.ErrRebuildBookIncorrectRequest, core.ErrRebuildBookEmptyPages)
	}

	newPageOrder := make(map[int]int, len(request.PageOrder))

	if request.Flags.PageReOrder {
		for i, pageNumber := range request.PageOrder {
			newPageOrder[pageNumber] = i + 1
		}

		for _, pageNumber := range request.SelectedPages {
			if _, ok := newPageOrder[pageNumber]; !ok {
				return uuid.Nil, fmt.Errorf("%w: missing page %d in page order", core.ErrRebuildBookIncorrectRequest, pageNumber)
			}
		}
	}

	isNewBook := request.MergeWithBook == uuid.Nil

	bookToMerge, attributeToMerge, targetPageHashes, err := uc.rebuildBookGetTarget(ctx, request)
	if err != nil {
		return uuid.Nil, fmt.Errorf("rebuild: get target: %w", err)
	}

	resources, err := uc.rebuildBookPrepareResources(
		ctx,
		request.Flags,
		request.ModifiedOldBook.Book.ID,
		targetPageHashes,
	)
	if err != nil {
		return uuid.Nil, fmt.Errorf("rebuild: prepare resources: %w", err)
	}

	pagesInfo, err := uc.rebuildBookPages(
		ctx,
		request.Flags,
		slices.Clone(request.SelectedPages),
		&bookToMerge,
		resources,
		newPageOrder,
	)
	if err != nil {
		return uuid.Nil, fmt.Errorf("rebuild: pages: %w", err)
	}

	if isNewBook && len(pagesInfo.NewPages) == 0 {
		return uuid.Nil, fmt.Errorf("%w: after deduplicate for new book", core.ErrRebuildBookEmptyPages)
	}

	bookToMerge.AttributesParsed = true
	newAttributes := core.MergeAttributeMap(attributeToMerge, request.ModifiedOldBook.Attributes)

	newLabels, err := uc.rebuildBookLabels(
		ctx,
		bookToMerge,
		resources.SourceBook,
		request.Flags,
		request.ModifiedOldBook.Labels,
		pagesInfo,
	)
	if err != nil {
		return uuid.Nil, fmt.Errorf("rebuild: labels: %w", err)
	}

	err = uc.rebuildBookSave(
		ctx,
		isNewBook,
		bookToMerge,
		newAttributes,
		pagesInfo.NewPages,
		newLabels,
	)
	if err != nil {
		return uuid.Nil, fmt.Errorf("rebuild: save: %w", err)
	}

	err = uc.rebuildBookCleanSource(
		ctx,
		request.Flags,
		resources.SourceBook.ID,
		pagesInfo.UnusedSourceHashes,
		pagesInfo.SourcePageNumbers,
	)
	if err != nil {
		return uuid.Nil, fmt.Errorf("rebuild: clean source: %w", err)
	}

	return bookToMerge.ID, nil
}
