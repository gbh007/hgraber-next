package agentapi

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (c *Controller) APIV2HighwayTokenPost(ctx context.Context) (*agentapi.APIV2HighwayTokenPostOK, error) {
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

	return &agentapi.APIV2HighwayTokenPostOK{
		ValidUntil: vu,
		Token:      token,
	}, nil
}
