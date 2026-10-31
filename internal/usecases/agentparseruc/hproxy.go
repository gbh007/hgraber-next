package agentparseruc

import (
	"context"
	"fmt"
	"net/url"

	"github.com/gbh007/hgraber-next/external/parserdto"
	"github.com/samber/lo"
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

func (uc *UseCase) HProxyList(ctx context.Context, u url.URL, onePage bool) (parserdto.BookList, error) {
	nextUrl := &u

	list := parserdto.BookList{}

	for nextUrl != nil {
		units, err := uc.loader.HProxyList(ctx, u)
		if err != nil {
			return parserdto.BookList{}, fmt.Errorf("loader: %w", err)
		}

		list.Units = append(list.Units, units.Units...)
		list.NextURL = units.NextURL

		if onePage {
			nextUrl = nil
		} else {
			nextUrl = units.NextURL
		}
	}

	list.Units = lo.Filter(list.Units, func(u parserdto.BookListUnit, _ int) bool {
		return u.Type != parserdto.UnknownHProxyListUnitType
	})

	return list, nil
}
