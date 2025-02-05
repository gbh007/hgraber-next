package agenthandlers

import (
	"context"
	"errors"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *AgentHandlersController) APIAgentGetPost(ctx context.Context, req *serverAPI.APIAgentGetPostReq) (serverAPI.APIAgentGetPostRes, error) {
	agent, err := c.agentUseCases.Agent(ctx, req.ID)

	if errors.Is(err, core.AgentNotFoundError) {
		return &serverAPI.APIAgentGetPostNotFound{
			InnerCode: apiservercore.AgentUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &serverAPI.APIAgentGetPostInternalServerError{
			InnerCode: apiservercore.AgentUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	result := apiservercore.ConvertAgentToAPI(agent)

	return &result, nil
}
