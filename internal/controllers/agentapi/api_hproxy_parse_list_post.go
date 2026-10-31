package agentapi

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/external/parserdto"
	"github.com/gbh007/hgraber-next/internal/pkg"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (c *Controller) APIHproxyParseListPost(ctx context.Context, req *agentapi.APIHproxyParseListPostReq) (*agentapi.APIHproxyParseListPostOK, error) {
	if c.parsingUseCases == nil {
		return nil, apiError{
			Code:    http.StatusBadRequest,
			Details: "unsupported api",
		}
	}

	list, err := c.parsingUseCases.HProxyList(ctx, req.URL)
	if err != nil {
		return nil, err
	}

	var nextPage agentapi.OptURI

	if list.NextURL != nil {
		nextPage = agentapi.NewOptURI(*list.NextURL)
	}

	return &agentapi.APIHproxyParseListPostOK{
		Results: pkg.Map(list.Units, func(u parserdto.BookListUnit) agentapi.APIHproxyParseListPostOKResultsItem {
			var t agentapi.APIHproxyParseListPostOKResultsItemType

			switch u.Type {
			case parserdto.DetailsHProxyListUnitType:
				t = agentapi.APIHproxyParseListPostOKResultsItemTypeDetails
			case parserdto.ListHProxyListUnitType:
				t = agentapi.APIHproxyParseListPostOKResultsItemTypeList
			}

			return agentapi.APIHproxyParseListPostOKResultsItem{
				LinkURL:    u.LinkURL,
				Name:       OptString(u.Name),
				PreviewURL: OptURL(u.PreviewURL),
				Type:       t,
			}
		}),
		NextURL: nextPage,
	}, nil
}
