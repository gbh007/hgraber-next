package bff

import (
	"context"
	"errors"
	"fmt"

	"github.com/gbh007/hgraber-next/domain/bff"
	"github.com/gbh007/hgraber-next/domain/core"
)

func (uc *UseCase) BookList(ctx context.Context, filter core.BookFilter) (bff.BookList, error) {
	count, err := uc.storage.BookCount(ctx, filter)
	if err != nil {
		return bff.BookList{}, fmt.Errorf("storage: get book count: %w", err)
	}

	if count == 0 { // Нет смысла пытаться обрабатывать дальше.
		return bff.BookList{}, nil
	}

	result := bff.BookList{
		Count: count,
		// FIXME: тут костыли, необходимо отказаться от расчета на сервере
		Pages: generatePagination(
			totalToPages(filter.Offset, filter.Limit)+1,
			totalToPages(count, filter.Limit),
		),
	}

	attributesInfo, err := uc.storage.Attributes(ctx)
	if err != nil {
		return bff.BookList{}, fmt.Errorf("storage: get attributes info: %w", err)
	}

	ids, err := uc.storage.BookIDs(ctx, filter)
	if err != nil {
		return bff.BookList{}, fmt.Errorf("get ids :%w", err)
	}

	result.Books = make([]bff.BookShort, 0, len(ids))

	for _, bookID := range ids {
		book, err := uc.storage.GetBook(ctx, bookID)
		if err != nil && !errors.Is(err, core.PageNotFoundError) {
			return bff.BookList{}, fmt.Errorf("storage: get book (%s): %w", bookID.String(), err)
		}

		page, err := uc.storage.BookPageWithHash(ctx, bookID, core.PageNumberForPreview)
		if err != nil && !errors.Is(err, core.PageNotFoundError) {
			return bff.BookList{}, fmt.Errorf("storage: get page (%s): %w", bookID.String(), err)
		}

		attributes, err := uc.storage.BookAttributes(ctx, bookID)
		if err != nil {
			return bff.BookList{}, fmt.Errorf("storage: get attributes (%s): %w", bookID.String(), err)
		}

		bffBook := bff.BookShort{
			Book:        book,
			PreviewPage: bff.PageWithHashToPreview(page),
		}

		for _, attr := range convertBookAttributes(
			convertAttributes(attributesInfo),
			attributes,
		) {
			if attr.Code == "tag" {
				bffBook.Tags = attr.Values
			}
		}

		result.Books = append(result.Books, bffBook)
	}

	return result, nil
}
