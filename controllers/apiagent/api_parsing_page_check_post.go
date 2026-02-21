package apiagent

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *Controller) APIParsingPageCheckPost(
	ctx context.Context,
	req *agentapi.APIParsingPageCheckPostReq,
) (*agentapi.APIParsingPageCheckPostOK, error) {
	result, err := c.parsingUseCases.PagesExists(
		ctx,
		pkg.Map(req.Urls, func(u agentapi.APIParsingPageCheckPostReqUrlsItem) agentmodel.AgentPageURL {
			return agentmodel.AgentPageURL{
				BookURL:  u.BookURL,
				ImageURL: u.ImageURL,
			}
		}),
	)
	if err != nil {
		return nil, apiError{
			Code:      http.StatusInternalServerError,
			InnerCode: ParseUseCaseCode,
			Details:   err.Error(),
		}
	}

	return &agentapi.APIParsingPageCheckPostOK{
		Result: pkg.Map(result, func(p agentmodel.AgentPageCheckResult) agentapi.APIParsingPageCheckPostOKResultItem {
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
