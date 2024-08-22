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

	return &agentAPI.BooksCheckResult{
		Result: pkg.Map(result, func(v entities.AgentBookCheckResult) agentAPI.BooksCheckResultResultItem {
			switch {
			case v.IsPossible:
				return agentAPI.BooksCheckResultResultItem{
					URL:                v.URL,
					Result:             agentAPI.BooksCheckResultResultItemResultOk,
					PossibleDuplicates: v.PossibleDuplicates,
				}

			case v.IsUnsupported:
				return agentAPI.BooksCheckResultResultItem{
					URL:    v.URL,
					Result: agentAPI.BooksCheckResultResultItemResultUnsupported,
				}

			case v.HasError:
				return agentAPI.BooksCheckResultResultItem{
					URL:          v.URL,
					Result:       agentAPI.BooksCheckResultResultItemResultError,
					ErrorDetails: agentAPI.NewOptString(v.ErrorReason),
				}

			default:
				return agentAPI.BooksCheckResultResultItem{
					URL:          v.URL,
					Result:       agentAPI.BooksCheckResultResultItemResultError,
					ErrorDetails: agentAPI.NewOptString("unknown result state"),
				}
			}
		}),
	}, nil
}
