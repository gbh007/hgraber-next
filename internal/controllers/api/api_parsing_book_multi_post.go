package api

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (c *Controller) APIParsingBookMultiPost(ctx context.Context, req *agentapi.APIParsingBookMultiPostReq) (*agentapi.BooksCheckResult, error) {
	if c.parsingUseCases == nil {
		return nil, apiError{
			Code:    http.StatusBadRequest,
			Details: "unsupported api",
		}
	}

	result, err := c.parsingUseCases.MultiHandle(ctx, req.URL)
	if err != nil {
		return nil, err
	}

	return &agentapi.BooksCheckResult{
		Result: convertBooksCheckResultResult(result),
	}, nil
}
