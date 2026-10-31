package agentcacheapi

import (
	"context"

	"github.com/gbh007/hgraber-next/internal/domain/serveragentdto"
	"github.com/gbh007/hgraber-next/internal/pkg"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (c *Controller) APIParsingPageCheckPost(
	ctx context.Context,
	req *agentapi.APIParsingPageCheckPostReq,
) (*agentapi.APIParsingPageCheckPostOK, error) {
	result, err := c.parsingUseCases.PagesExists(
		ctx,
		pkg.Map(req.Urls, func(u agentapi.APIParsingPageCheckPostReqUrlsItem) serveragentdto.AgentPageURL {
			return serveragentdto.AgentPageURL{
				BookURL:  u.BookURL,
				ImageURL: u.ImageURL,
			}
		}),
	)
	if err != nil {
		return nil, err //nolint:wrapcheck // будет исправлено позднее
	}

	return &agentapi.APIParsingPageCheckPostOK{
		Result: pkg.Map(result, func(p serveragentdto.AgentPageCheckResult) agentapi.APIParsingPageCheckPostOKResultItem {
			item := agentapi.APIParsingPageCheckPostOKResultItem{
				BookURL:  p.BookURL,
				ImageURL: p.ImageURL,
			}

			switch {
			case p.HasError:
				item.Result = agentapi.APIParsingPageCheckPostOKResultItemResultError
				item.ErrorDetails = agentapi.NewOptString(p.ErrorReason)

			case p.IsPossible:
				item.Result = agentapi.APIParsingPageCheckPostOKResultItemResultOk

			case p.IsUnsupported:
				item.Result = agentapi.APIParsingPageCheckPostOKResultItemResultUnsupported
			}

			return item
		}),
	}, nil
}
