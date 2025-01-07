package rebuilder

import (
	"context"
	"fmt"
	"maps"
	"slices"
	"time"

	"hgnext/internal/entities"
)

func (uc *UseCase) rebuildBookPages(
	_ context.Context,
	flags entities.RebuildBookRequestFlags,
	selectedPages []int,
	bookToMerge *entities.Book,
	resources rebuildPageResources,
) (rebuildedPagesInfo, error) {
	selectedPages = slices.Compact(selectedPages)
	slices.Sort(selectedPages)

	existsPageHashes := make(map[entities.FileHash]struct{}, len(selectedPages))

	pagesRemap := make(map[int]int, len(selectedPages))
	newPages := make([]entities.Page, 0, len(selectedPages))
	sourcePageNumbers := make([]int, 0, len(selectedPages))

	unusedSourceHashes := maps.Clone(resources.UnusedSourceHashes)
	if unusedSourceHashes == nil {
		unusedSourceHashes = make(map[entities.FileHash]struct{})
	}

	newPageNumberCounter := 0

	for _, oldPageNumber := range selectedPages {
		sourcePage, ok := resources.SourcePagesMap[oldPageNumber]
		if !ok {
			return rebuildedPagesInfo{}, fmt.Errorf("%w (%d)", errMissingSourcePage, oldPageNumber)
		}

		if _, ok := resources.ForbiddenHashes[sourcePage.FileHash]; ok {
			continue
		}

		if flags.OnlyUniquePages || flags.Only1CopyPages {
			if _, ok := existsPageHashes[sourcePage.FileHash]; ok {
				continue
			}

			existsPageHashes[sourcePage.FileHash] = struct{}{}
		}

		newPageNumberCounter++
		newPageNumber := bookToMerge.PageCount + newPageNumberCounter
		pagesRemap[oldPageNumber] = newPageNumber

		sourcePage.BookID = bookToMerge.ID
		sourcePage.PageNumber = newPageNumber
		sourcePage.CreateAt = time.Now()

		newPages = append(newPages, sourcePage.Page)
		sourcePageNumbers = append(sourcePageNumbers, oldPageNumber)

		delete(unusedSourceHashes, sourcePage.FileHash)
	}

	bookToMerge.PageCount += newPageNumberCounter

	return rebuildedPagesInfo{
		PagesRemap:         pagesRemap,
		SourcePageNumbers:  sourcePageNumbers,
		NewPages:           newPages,
		UnusedSourceHashes: unusedSourceHashes,
	}, nil
}
