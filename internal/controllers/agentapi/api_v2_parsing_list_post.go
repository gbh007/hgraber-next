package agentapi

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/external/parserdto"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
	"github.com/samber/lo"
)

func (c *Controller) APIV2ParsingListPost(
	ctx context.Context,
	req *agentapi.APIV2ParsingListPostReq,
) (*agentapi.APIV2ParsingListPostOK, error) {
	if c.parsingUseCases == nil {
		return nil, apiError{
			Code:    http.StatusBadRequest,
			Details: "unsupported api",
		}
	}

	list, err := c.parsingUseCases.List(ctx, req.URL, req.OnePage.Value)
	if err != nil {
		return nil, err
	}

	var nextPage agentapi.OptURI

	if list.NextURL != nil {
		nextPage = agentapi.NewOptURI(*list.NextURL)
	}

	return &agentapi.APIV2ParsingListPostOK{
		Results: lo.Map(list.Units, func(u parserdto.BookListUnit, _ int) agentapi.APIV2ParsingListPostOKResultsItem {
			var t agentapi.APIV2ParsingListPostOKResultsItemType

			switch u.Type {
			case parserdto.DetailsHProxyListUnitType:
				t = agentapi.APIV2ParsingListPostOKResultsItemTypeDetails
			case parserdto.ListHProxyListUnitType:
				t = agentapi.APIV2ParsingListPostOKResultsItemTypeList
			case parserdto.UnknownHProxyListUnitType:
				// неизвестный тип — оставляем значение по умолчанию
			}

			return agentapi.APIV2ParsingListPostOKResultsItem{
				LinkURL:    u.LinkURL,
				Name:       OptString(u.Name),
				PreviewURL: OptURL(u.PreviewURL),
				Type:       t,
			}
		}),
		NextURL: nextPage,
	}, nil
}
