package rebuilder

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/entities"
)

func (uc *UseCase) rebuildBookPrepareResources(
	ctx context.Context,
	flags entities.RebuildBookRequestFlags,
	oldBookID uuid.UUID,
	targetPageHashes map[entities.FileHash]struct{},
) (
	rebuildPageResources,
	error,
) {
	sourceBook, err := uc.storage.GetBook(ctx, oldBookID)
	if err != nil {
		return rebuildPageResources{}, fmt.Errorf("storage: get source book: %w", err)
	}

	sourcePages, err := uc.storage.BookPagesWithHash(ctx, oldBookID)
	if err != nil {
		return rebuildPageResources{}, fmt.Errorf("storage: get source pages: %w", err)
	}

	sourcePagesMap := make(map[int]entities.PageWithHash, len(sourcePages))
	sourcePagesHashes := make(map[entities.FileHash]struct{}, len(sourcePages))
	sourcePagesMD5Sums := make([]string, 0, len(sourcePages))
	unusedSourceHashes := make(map[entities.FileHash]struct{}, len(sourcePages))

	for _, page := range sourcePages {
		sourcePagesMap[page.PageNumber] = page
		sourcePagesMD5Sums = append(sourcePagesMD5Sums, page.Md5Sum)
		sourcePagesHashes[page.FileHash] = struct{}{}

		if flags.MarkUnusedPagesAsDeadHash || flags.MarkUnusedPagesAsDeleted {
			unusedSourceHashes[page.FileHash] = struct{}{}
		}
	}

	forbiddenHashes := make(map[entities.FileHash]struct{}, len(targetPageHashes)+len(sourcePagesHashes))

	if flags.ExcludeDeadHashPages {
		deadHashes, err := uc.storage.DeadHashesByMD5Sums(ctx, sourcePagesMD5Sums)
		if err != nil {
			return rebuildPageResources{}, fmt.Errorf("storage: get dead hashes: %w", err)
		}

		for _, hash := range deadHashes {
			forbiddenHashes[hash.FileHash] = struct{}{}
		}
	}

	if flags.Only1CopyPages {
		pages, err := uc.storage.BookPagesWithHashByMD5Sums(ctx, sourcePagesMD5Sums)
		if err != nil {
			return rebuildPageResources{}, fmt.Errorf("storage: get pages by source hashes: %w", err)
		}

		for _, page := range pages {
			// Отсекаем не совпадения из-за неполного ограничения в БД
			if _, ok := sourcePagesHashes[page.FileHash]; !ok {
				continue
			}

			// Это текущая книга, ее исключать не надо
			if page.BookID == oldBookID {
				continue
			}

			forbiddenHashes[page.FileHash] = struct{}{}
		}
	}

	if flags.OnlyUniquePages || flags.Only1CopyPages {
		for hash := range targetPageHashes {
			forbiddenHashes[hash] = struct{}{}
		}
	}

	if flags.MarkUnusedPagesAsDeadHash || flags.MarkUnusedPagesAsDeleted {
		for hash := range targetPageHashes {
			delete(unusedSourceHashes, hash)
		}
	}

	return rebuildPageResources{
		SourceBook:         sourceBook,
		SourcePagesMap:     sourcePagesMap,
		ForbiddenHashes:    forbiddenHashes,
		UnusedSourceHashes: unusedSourceHashes,
	}, nil
}
