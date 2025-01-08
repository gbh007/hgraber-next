package apiserver

import (
	"context"
	"errors"

	"hgnext/internal/entities"
	"hgnext/open_api/serverAPI"
)

func (c *Controller) APIAgentGetPost(ctx context.Context, req *serverAPI.APIAgentGetPostReq) (serverAPI.APIAgentGetPostRes, error) {
	agent, err := c.agentUseCases.Agent(ctx, req.ID)

	if errors.Is(err, entities.AgentNotFoundError) {
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
