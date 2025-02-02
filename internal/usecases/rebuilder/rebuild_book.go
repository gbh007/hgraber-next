package rebuilder

import (
	"context"
	"fmt"
	"slices"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/internal/entities"
)

type rebuildPageResources struct {
	SourceBook         entities.Book
	SourcePagesMap     map[int]entities.PageWithHash
	ForbiddenHashes    map[entities.FileHash]struct{}
	UnusedSourceHashes map[entities.FileHash]struct{}
}

type rebuildedPagesInfo struct {
	PagesRemap         map[int]int
	SourcePageNumbers  []int
	NewPages           []entities.Page
	UnusedSourceHashes map[entities.FileHash]struct{}
}

func (uc *UseCase) RebuildBook(ctx context.Context, request entities.RebuildBookRequest) (_ uuid.UUID, returnErr error) {
	if len(request.SelectedPages) == 0 {
		return uuid.Nil, fmt.Errorf("%w: %w", entities.ErrRebuildBookIncorrectRequest, entities.ErrRebuildBookEmptyPages)
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
	)
	if err != nil {
		return uuid.Nil, fmt.Errorf("rebuild: pages: %w", err)
	}

	if isNewBook && len(pagesInfo.NewPages) == 0 {
		return uuid.Nil, fmt.Errorf("%w: after deduplicate for new book", entities.ErrRebuildBookEmptyPages)
	}

	bookToMerge.AttributesParsed = true
	newAttributes := entities.MergeAttributeMap(attributeToMerge, request.ModifiedOldBook.Attributes)

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
