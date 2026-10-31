package agent

import (
	"context"
	"fmt"
	"net/url"

	"github.com/gbh007/hgraber-next/internal/domain/entities"
	"github.com/gbh007/hgraber-next/internal/pkg"
)

func (uc *UseCase) HProxyBook(ctx context.Context, u url.URL, pageLimit *int) (entities.HProxyBookDetails, error) {
	parser, err := uc.loader.Load(ctx, u)
	if err != nil {
		return entities.HProxyBookDetails{}, fmt.Errorf("load parser: %w", err)
	}

	details, err := parser.HProxyBookDetails(ctx, pageLimit)
	if err != nil {
		return entities.HProxyBookDetails{}, fmt.Errorf("parse: %w", err)
	}

	return details, nil
}

func (uc *UseCase) HProxyList(ctx context.Context, u url.URL) (entities.HProxyList, error) {
	list, err := uc.loader.HProxyList(ctx, u)
	if err != nil {
		return entities.HProxyList{}, fmt.Errorf("loader: %w", err)
	}

	list.Units = pkg.SliceFilter(list.Units, func(u entities.HProxyListUnit) bool {
		return u.Type != entities.UnknownHProxyListUnitType
	})

	return list, nil
}
