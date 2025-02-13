package rebuilderusecase

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (uc *UseCase) rebuildBookGetTarget(ctx context.Context, request core.RebuildBookRequest) (
	core.Book,
	map[string][]string,
	map[core.FileHash]struct{},
	error,
) {
	isNewBook := request.MergeWithBook == uuid.Nil

	var (
		err              error
		bookToMerge      core.Book
		attributeToMerge map[string][]string
		targetPageHashes map[core.FileHash]struct{}
	)

	if isNewBook {
		bookToMerge = core.Book{
			ID:        uuid.Must(uuid.NewV7()),
			Name:      request.ModifiedOldBook.Book.Name,
			OriginURL: request.ModifiedOldBook.Book.OriginURL,
			PageCount: 0,
			IsRebuild: true,
			CreateAt:  time.Now().UTC(),
		}

		if request.Flags.AutoVerify {
			bookToMerge.Verified = true
			bookToMerge.VerifiedAt = time.Now().UTC()
		}

		return bookToMerge, nil, nil, nil
	}

	bookToMerge, err = uc.storage.GetBook(ctx, request.MergeWithBook)
	if err != nil {
		return core.Book{}, nil, nil, fmt.Errorf("storage: get book to merge: %w", err)
	}

	if !bookToMerge.IsRebuild {
		return core.Book{}, nil, nil, fmt.Errorf("%w: not rebuilded book", core.ErrRebuildBookForbiddenMerge)
	}

	if bookToMerge.Deleted {
		return core.Book{}, nil, nil, fmt.Errorf("%w: deleted book", core.ErrRebuildBookForbiddenMerge)
	}

	attributeToMerge, err = uc.storage.BookOriginAttributes(ctx, request.MergeWithBook)
	if err != nil {
		return core.Book{}, nil, nil, fmt.Errorf("storage: get attributes to merge: %w", err)
	}

	pages, err := uc.storage.BookPagesWithHash(ctx, request.MergeWithBook)
	if err != nil {
		return core.Book{}, nil, nil, fmt.Errorf("storage: get page to unique: %w", err)
	}

	targetPageHashes = make(map[core.FileHash]struct{}, len(pages)+len(request.SelectedPages))

	for _, page := range pages {
		targetPageHashes[page.FileHash] = struct{}{}
	}

	return bookToMerge, attributeToMerge, targetPageHashes, nil
}
