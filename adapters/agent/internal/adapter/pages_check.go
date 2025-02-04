package adapter

import (
	"context"
	"fmt"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/open_api/agentAPI"
	"github.com/gbh007/hgraber-next/pkg"
)

func (a *Adapter) PagesCheck(ctx context.Context, urls []agentmodel.AgentPageURL) ([]agentmodel.AgentPageCheckResult, error) {
	res, err := a.rawClient.APIParsingPageCheckPost(ctx, &agentAPI.APIParsingPageCheckPostReq{
		Urls: pkg.Map(urls, func(u agentmodel.AgentPageURL) agentAPI.APIParsingPageCheckPostReqUrlsItem {
			return agentAPI.APIParsingPageCheckPostReqUrlsItem{
				BookURL:  u.BookURL,
				ImageURL: u.ImageURL,
			}
		}),
	})
	if err != nil {
		return nil, err
	}

	var result []agentmodel.AgentPageCheckResult

	switch typedRes := res.(type) {
	case *agentAPI.APIParsingPageCheckPostOK:
		result = pkg.Map(typedRes.Result, func(v agentAPI.APIParsingPageCheckPostOKResultItem) agentmodel.AgentPageCheckResult {
			return agentmodel.AgentPageCheckResult{
				BookURL:       v.BookURL,
				ImageURL:      v.ImageURL,
				IsUnsupported: v.Result == agentAPI.APIParsingPageCheckPostOKResultItemResultUnsupported,
				IsPossible:    v.Result == agentAPI.APIParsingPageCheckPostOKResultItemResultOk,
				HasError:      v.Result == agentAPI.APIParsingPageCheckPostOKResultItemResultError,
				ErrorReason:   v.ErrorDetails.Value,
			}
		})

	case *agentAPI.APIParsingPageCheckPostBadRequest:
		return nil, fmt.Errorf("%w: %s", agentmodel.AgentAPIBadRequest, typedRes.Details.Value)

	case *agentAPI.APIParsingPageCheckPostUnauthorized:
		return nil, fmt.Errorf("%w: %s", agentmodel.AgentAPIUnauthorized, typedRes.Details.Value)

	case *agentAPI.APIParsingPageCheckPostForbidden:
		return nil, fmt.Errorf("%w: %s", agentmodel.AgentAPIForbidden, typedRes.Details.Value)

	case *agentAPI.APIParsingPageCheckPostInternalServerError:
		return nil, fmt.Errorf("%w: %s", agentmodel.AgentAPIInternalError, typedRes.Details.Value)

	default:
		return nil, agentmodel.AgentAPIUnknownResponse
	}

	return result, nil
}
