package agentparseruc

import (
	"context"
	"fmt"
	"net/url"

	"github.com/gbh007/hgraber-next/external/parserdto"
	"github.com/gbh007/hgraber-next/internal/pkg"
)

func (uc *UseCase) HProxyBook(ctx context.Context, u url.URL, pageLimit *int) (parserdto.Book, error) {
	parser, err := uc.loader.Load(ctx, u)
	if err != nil {
		return parserdto.Book{}, fmt.Errorf("load parser: %w", err)
	}

	details, err := parser.HProxyBookDetails(ctx, pageLimit)
	if err != nil {
		return parserdto.Book{}, fmt.Errorf("parse: %w", err)
	}

	return details, nil
}

func (uc *UseCase) HProxyList(ctx context.Context, u url.URL) (parserdto.BookList, error) {
	list, err := uc.loader.HProxyList(ctx, u)
	if err != nil {
		return parserdto.BookList{}, fmt.Errorf("loader: %w", err)
	}

	list.Units = pkg.SliceFilter(list.Units, func(u parserdto.BookListUnit) bool {
		return u.Type != parserdto.UnknownHProxyListUnitType
	})

	return list, nil
}
