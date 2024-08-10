package apiagent

import (
	"context"

	"hgnext/internal/entities"
	"hgnext/internal/pkg"
	"hgnext/open_api/agentAPI"
)

func (c *Controller) APIParsingBookCheckPost(ctx context.Context, req *agentAPI.APIParsingBookCheckPostReq) (agentAPI.APIParsingBookCheckPostRes, error) {
	result, err := c.parsingUseCases.CheckBooks(ctx, req.Urls)
	if err != nil {
		return &agentAPI.APIParsingBookCheckPostInternalServerError{
			InnerCode: ParseUseCaseCode,
			Details:   agentAPI.NewOptString(err.Error()),
		}, nil
	}

	return &agentAPI.APIParsingBookCheckPostOK{
		Result: pkg.Map(result, func(v entities.AgentBookCheckResult) agentAPI.APIParsingBookCheckPostOKResultItem {
			switch {
			case v.IsPossible:
				return agentAPI.APIParsingBookCheckPostOKResultItem{
					URL:                v.URL,
					Result:             agentAPI.APIParsingBookCheckPostOKResultItemResultOk,
					PossibleDuplicates: v.PossibleDuplicates,
				}

			case v.IsUnsupported:
				return agentAPI.APIParsingBookCheckPostOKResultItem{
					URL:    v.URL,
					Result: agentAPI.APIParsingBookCheckPostOKResultItemResultUnsupported,
				}

			case v.HasError:
				return agentAPI.APIParsingBookCheckPostOKResultItem{
					URL:          v.URL,
					Result:       agentAPI.APIParsingBookCheckPostOKResultItemResultError,
					ErrorDetails: agentAPI.NewOptString(v.ErrorReason),
				}

			default:
				return agentAPI.APIParsingBookCheckPostOKResultItem{
					URL:          v.URL,
					Result:       agentAPI.APIParsingBookCheckPostOKResultItemResultError,
					ErrorDetails: agentAPI.NewOptString("unknown result state"),
				}
			}
		}),
	}, nil
}
