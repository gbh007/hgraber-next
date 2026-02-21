package agenthandlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *AgentHandlersController) APIAgentGetPost(
	ctx context.Context,
	req *serverapi.APIAgentGetPostReq,
) (*serverapi.Agent, error) {
	agent, err := c.agentUseCases.Agent(ctx, req.ID)

	if errors.Is(err, core.ErrAgentNotFound) {
		return nil, apiservercore.APIError{
			Code:      http.StatusNotFound,
			InnerCode: apiservercore.AgentUseCaseCode,
			Details:   err.Error(),
		}
	}

	if err != nil {
		return nil, apiservercore.APIError{
			Code:      http.StatusInternalServerError,
			InnerCode: apiservercore.AgentUseCaseCode,
			Details:   err.Error(),
		}
	}

	result := apiservercore.ConvertAgentToAPI(agent)

	return &result, nil
}
