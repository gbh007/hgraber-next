package agent

import (
	"context"
	"fmt"
	"net/url"

	"github.com/gbh007/hgraber-next/internal/domain/agentdto"
	"github.com/gbh007/hgraber-next/internal/pkg"
)

func (uc *UseCase) HProxyBook(ctx context.Context, u url.URL, pageLimit *int) (agentdto.HProxyBookDetails, error) {
	parser, err := uc.loader.Load(ctx, u)
	if err != nil {
		return agentdto.HProxyBookDetails{}, fmt.Errorf("load parser: %w", err)
	}

	details, err := parser.HProxyBookDetails(ctx, pageLimit)
	if err != nil {
		return agentdto.HProxyBookDetails{}, fmt.Errorf("parse: %w", err)
	}

	return details, nil
}

func (uc *UseCase) HProxyList(ctx context.Context, u url.URL) (agentdto.HProxyList, error) {
	list, err := uc.loader.HProxyList(ctx, u)
	if err != nil {
		return agentdto.HProxyList{}, fmt.Errorf("loader: %w", err)
	}

	list.Units = pkg.SliceFilter(list.Units, func(u agentdto.HProxyListUnit) bool {
		return u.Type != agentdto.UnknownHProxyListUnitType
	})

	return list, nil
}
