package agentcacheapi

import (
	"context"

	"github.com/gbh007/hgraber-next/internal/domain/serveragentdto"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
	"github.com/samber/lo"
)

func (c *Controller) APIV2ParsingCheckPost(
	ctx context.Context,
	req *agentapi.APIV2ParsingCheckPostReq,
) (*agentapi.BooksCheckResult, error) {
	result, err := c.parsingUseCases.BooksExists(ctx, req.Urls)
	if err != nil {
		return nil, err //nolint:wrapcheck // будет исправлено позднее
	}

	return &agentapi.BooksCheckResult{
		Result: lo.Map(result, func(v serveragentdto.AgentBookCheckResult, _ int) agentapi.BooksCheckResultResultItem {
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
