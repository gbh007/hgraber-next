package rebuilder

import (
	"context"
	"fmt"
	"maps"
	"slices"
	"time"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (uc *UseCase) rebuildBookPages(
	_ context.Context,
	flags core.RebuildBookRequestFlags,
	selectedPages []int,
	bookToMerge *core.Book,
	resources rebuildPageResources,
	newPageOrder map[int]int,
) (rebuildedPagesInfo, error) {
	selectedPages = slices.Compact(selectedPages)

	if len(newPageOrder) > 0 {
		slices.SortFunc(selectedPages, func(a, b int) int {
			return newPageOrder[a] - newPageOrder[b]
		})
	} else {
		slices.Sort(selectedPages)
	}

	existsPageHashes := make(map[core.FileHash]struct{}, len(selectedPages))

	pagesRemap := make(map[int]int, len(selectedPages))
	newPages := make([]core.Page, 0, len(selectedPages))
	sourcePageNumbers := make([]int, 0, len(selectedPages))

	unusedSourceHashes := maps.Clone(resources.UnusedSourceHashes)
	if unusedSourceHashes == nil {
		unusedSourceHashes = make(map[core.FileHash]struct{})
	}

	newPageNumberCounter := 0

	for _, oldPageNumber := range selectedPages {
		sourcePage, ok := resources.SourcePagesMap[oldPageNumber]
		if !ok {
			return rebuildedPagesInfo{}, fmt.Errorf("%w (%d)", core.ErrRebuildBookMissingSourcePage, oldPageNumber)
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
