package apiserver

import (
	"context"
	"errors"

	"github.com/gbh007/hgraber-next/entities"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *Controller) APIAgentDeletePost(ctx context.Context, req *serverAPI.APIAgentDeletePostReq) (serverAPI.APIAgentDeletePostRes, error) {
	err := c.agentUseCases.DeleteAgent(ctx, req.ID)

	if errors.Is(err, entities.AgentNotFoundError) {
		return &serverAPI.APIAgentDeletePostNotFound{
			InnerCode: AgentUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &serverAPI.APIAgentDeletePostInternalServerError{
			InnerCode: AgentUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIAgentDeletePostNoContent{}, nil
}
