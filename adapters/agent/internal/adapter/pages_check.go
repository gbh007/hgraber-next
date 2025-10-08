package adapter

import (
	"context"
	"fmt"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (a *Adapter) PagesCheck(
	ctx context.Context,
	urls []agentmodel.AgentPageURL,
) ([]agentmodel.AgentPageCheckResult, error) {
	res, err := a.rawClient.APIParsingPageCheckPost(ctx, &agentapi.APIParsingPageCheckPostReq{
		Urls: pkg.Map(urls, func(u agentmodel.AgentPageURL) agentapi.APIParsingPageCheckPostReqUrlsItem {
			return agentapi.APIParsingPageCheckPostReqUrlsItem{
				BookURL:  u.BookURL,
				ImageURL: u.ImageURL,
			}
		}),
	})
	if err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	var result []agentmodel.AgentPageCheckResult

	switch typedRes := res.(type) {
	case *agentapi.APIParsingPageCheckPostOK:
		result = pkg.Map(
			typedRes.Result,
			func(v agentapi.APIParsingPageCheckPostOKResultItem) agentmodel.AgentPageCheckResult {
				return agentmodel.AgentPageCheckResult{
					BookURL:       v.BookURL,
					ImageURL:      v.ImageURL,
					IsUnsupported: v.Result == agentapi.APIParsingPageCheckPostOKResultItemResultUnsupported,
					IsPossible:    v.Result == agentapi.APIParsingPageCheckPostOKResultItemResultOk,
					HasError:      v.Result == agentapi.APIParsingPageCheckPostOKResultItemResultError,
					ErrorReason:   v.ErrorDetails.Value,
				}
			},
		)

	case *agentapi.APIParsingPageCheckPostBadRequest:
		return nil, fmt.Errorf("%w: %s", agentmodel.ErrAgentAPIBadRequest, typedRes.Details.Value)

	case *agentapi.APIParsingPageCheckPostUnauthorized:
		return nil, fmt.Errorf("%w: %s", agentmodel.ErrAgentAPIUnauthorized, typedRes.Details.Value)

	case *agentapi.APIParsingPageCheckPostForbidden:
		return nil, fmt.Errorf("%w: %s", agentmodel.ErrAgentAPIForbidden, typedRes.Details.Value)

	case *agentapi.APIParsingPageCheckPostInternalServerError:
		return nil, fmt.Errorf("%w: %s", agentmodel.ErrAgentAPIInternalError, typedRes.Details.Value)

	default:
		return nil, agentmodel.ErrAgentAPIUnknownResponse
	}

	return result, nil
}
