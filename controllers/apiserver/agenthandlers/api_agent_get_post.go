package agenthandlers

import (
	"context"
	"errors"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *AgentHandlersController) APIAgentGetPost(ctx context.Context, req *serverapi.APIAgentGetPostReq) (serverapi.APIAgentGetPostRes, error) {
	agent, err := c.agentUseCases.Agent(ctx, req.ID)

	if errors.Is(err, core.AgentNotFoundError) {
		return &serverapi.APIAgentGetPostNotFound{
			InnerCode: apiservercore.AgentUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &serverapi.APIAgentGetPostInternalServerError{
			InnerCode: apiservercore.AgentUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	result := apiservercore.ConvertAgentToAPI(agent)

	return &result, nil
}
