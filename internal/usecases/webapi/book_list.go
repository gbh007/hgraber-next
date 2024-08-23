package webapi

import (
	"context"
	"fmt"

	"hgnext/internal/entities"
	"hgnext/internal/pkg"
)

func (uc *UseCase) BookList(ctx context.Context, filter entities.BookFilter) ([]entities.BookToWeb, []int, error) {
	total, err := uc.storage.BookCount(ctx, filter)
	if err != nil {
		return nil, nil, fmt.Errorf("get book count from storage: %w", err)
	}

	books, err := uc.bookRequester.Books(ctx, filter)
	if err != nil {
		return nil, nil, fmt.Errorf("get books from requester: %w", err)
	}

	// FIXME: тут костыли, необходимо отказаться от расчета на сервере
	pages := generatePagination(totalToPages(filter.Offset, filter.Limit)+1, totalToPages(total, filter.Limit))

	return pkg.Map(books, uc.bookConvert), pages, nil
}
