package webapi

import (
	"context"
	"fmt"

	"hgnext/internal/entities"
	"hgnext/internal/pkg"
)

func (uc *UseCase) BookList(ctx context.Context, filter entities.BookFilter) (entities.BookListToWeb, error) {
	total, err := uc.storage.BookCount(ctx, filter)
	if err != nil {
		return entities.BookListToWeb{}, fmt.Errorf("get book count from storage: %w", err)
	}

	books, err := uc.bookRequester.Books(ctx, filter)
	if err != nil {
		return entities.BookListToWeb{}, fmt.Errorf("get books from requester: %w", err)
	}

	attributesInfo, err := uc.storage.Attributes(ctx)
	if err != nil {
		return entities.BookListToWeb{}, fmt.Errorf("storage: get attributes info: %w", err)
	}

	attributesInfoMap := convertAttributes(attributesInfo)

	// FIXME: тут костыли, необходимо отказаться от расчета на сервере
	pages := generatePagination(totalToPages(filter.Offset, filter.Limit)+1, totalToPages(total, filter.Limit))

	return entities.BookListToWeb{
		Books: pkg.Map(books, func(book entities.BookFull) entities.BookToWeb {
			return uc.bookConvert(book, attributesInfoMap)
		}),
		Pages: pages,
		Count: total,
	}, nil
}
