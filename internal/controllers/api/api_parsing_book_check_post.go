package api

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/internal/domain/serveragentdto"
	"github.com/gbh007/hgraber-next/internal/pkg"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (c *Controller) APIParsingBookCheckPost(ctx context.Context, req *agentapi.APIParsingBookCheckPostReq) (*agentapi.BooksCheckResult, error) {
	if c.parsingUseCases == nil {
		return nil, apiError{
			Code:    http.StatusBadRequest,
			Details: "unsupported api",
		}
	}

	result, err := c.parsingUseCases.CheckBooks(ctx, req.Urls)
	if err != nil {
		return nil, err
	}

	return &agentapi.BooksCheckResult{
		Result: convertBooksCheckResultResult(result),
	}, nil
}

func convertBooksCheckResultResult(result []serveragentdto.AgentBookCheckResult) []agentapi.BooksCheckResultResultItem {
	return pkg.Map(result, func(v serveragentdto.AgentBookCheckResult) agentapi.BooksCheckResultResultItem {
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
	})
}
