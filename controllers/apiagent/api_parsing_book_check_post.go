package apiagent

import (
	"context"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *Controller) APIParsingBookCheckPost(ctx context.Context, req *agentapi.APIParsingBookCheckPostReq) (agentapi.APIParsingBookCheckPostRes, error) {
	result, err := c.parsingUseCases.CheckBooks(ctx, req.Urls)
	if err != nil {
		return &agentapi.APIParsingBookCheckPostInternalServerError{
			InnerCode: ParseUseCaseCode,
			Details:   agentapi.NewOptString(err.Error()),
		}, nil
	}

	return &agentapi.BooksCheckResult{
		Result: pkg.Map(result, func(v agentmodel.AgentBookCheckResult) agentapi.BooksCheckResultResultItem {
			switch {
			case v.IsPossible:
				return agentapi.BooksCheckResultResultItem{
					URL:    v.URL,
					Result: agentapi.BooksCheckResultResultItemResultOk,
				}

			case v.IsUnsupported:
				return agentapi.BooksCheckResultResultItem{
					URL:    v.URL,
					Result: agentapi.BooksCheckResultResultItemResultUnsupported,
				}

			case v.HasError:
				return agentapi.BooksCheckResultResultItem{
					URL:          v.URL,
					Result:       agentapi.BooksCheckResultResultItemResultError,
					ErrorDetails: agentapi.NewOptString(v.ErrorReason),
				}

			default:
				return agentapi.BooksCheckResultResultItem{
					URL:          v.URL,
					Result:       agentapi.BooksCheckResultResultItemResultError,
					ErrorDetails: agentapi.NewOptString("unknown result state"),
				}
			}
		}),
	}, nil
}
