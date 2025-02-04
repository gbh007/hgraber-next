package deduplicator

import (
	"context"
	"errors"
	"fmt"
	"slices"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/entities"
)

// TODO: по факту для этого метода не нужны превью, подумать над выделением в BFF.
func (uc *UseCase) BooksByPage(ctx context.Context, bookID uuid.UUID, pageNumber int) ([]entities.BookWithPreviewPage, error) {
	originPage, err := uc.storage.BookPageWithHash(ctx, bookID, pageNumber)
	if err != nil {
		return nil, fmt.Errorf("get origin page: %w", err)
	}

	pages, err := uc.storage.BookPagesWithHashByHash(ctx, originPage.FileHash)
	if err != nil {
		return nil, fmt.Errorf("get pages by hash: %w", err)
	}

	books := make([]entities.BookWithPreviewPage, 0, min(len(pages), 10))
	booksHandled := make(map[uuid.UUID]struct{}, min(len(pages), 10))

	for _, page := range pages {
		if _, ok := booksHandled[page.BookID]; ok {
			continue
		}

		book, err := uc.storage.GetBook(ctx, page.BookID)
		if err != nil {
			return nil, fmt.Errorf("get book %s: %w", page.BookID.String(), err)
		}

		previewPage, err := uc.storage.BookPageWithHash(ctx, book.ID, entities.PageNumberForPreview)
		if err != nil && !errors.Is(err, entities.PageNotFoundError) { // Отсутствие превью это нормально
			return nil, fmt.Errorf("get book %s preview page: %w", page.BookID.String(), err)
		}

		booksHandled[page.BookID] = struct{}{}

		books = append(books, entities.BookWithPreviewPage{
			Book:        book,
			PreviewPage: previewPage.ToPreview(),
		})
	}

	slices.SortFunc(books, func(a, b entities.BookWithPreviewPage) int {
		return -1 * a.CreateAt.Compare(b.CreateAt) // Вначале идут самые новые книги
	})

	return books, nil
}
