package apiserver

import (
	"context"
	"errors"

	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *Controller) APIAgentGetPost(ctx context.Context, req *serverAPI.APIAgentGetPostReq) (serverAPI.APIAgentGetPostRes, error) {
	agent, err := c.agentUseCases.Agent(ctx, req.ID)

	if errors.Is(err, core.AgentNotFoundError) {
		return &serverAPI.APIAgentGetPostNotFound{
			InnerCode: AgentUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &serverAPI.APIAgentGetPostInternalServerError{
			InnerCode: AgentUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	result := convertAgentToAPI(agent)

	return &result, nil
}
