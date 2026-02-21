package apiagent

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *Controller) APIParsingBookCheckPost(
	ctx context.Context,
	req *agentapi.APIParsingBookCheckPostReq,
) (*agentapi.BooksCheckResult, error) {
	result, err := c.parsingUseCases.BooksExists(ctx, req.Urls)
	if err != nil {
		return nil, apiError{
			Code:      http.StatusInternalServerError,
			InnerCode: ParseUseCaseCode,
			Details:   err.Error(),
		}
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
