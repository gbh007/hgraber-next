package api

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (c *Controller) APIHighwayTokenCreatePost(ctx context.Context) (*agentapi.APIHighwayTokenCreatePostOK, error) {
	if c.highwayUseCase == nil {
		return nil, apiError{
			Code:    http.StatusBadRequest,
			Details: "unsupported api",
		}
	}

	token, vu, err := c.highwayUseCase.NewToken(ctx)
	if err != nil {
		return nil, err
	}

	return &agentapi.APIHighwayTokenCreatePostOK{
		ValidUntil: vu,
		Token:      token,
	}, nil
}
