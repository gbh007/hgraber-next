package apiserver

import (
	"context"
	"errors"

	"hgnext/internal/controllers/apiserver/internal/server"
	"hgnext/internal/entities"
)

func (c *Controller) APIAgentDeletePost(ctx context.Context, req *server.APIAgentDeletePostReq) (server.APIAgentDeletePostRes, error) {
	err := c.agentUseCases.DeleteAgent(ctx, req.ID)

	if errors.Is(err, entities.AgentNotFoundError) {
		return &server.APIAgentDeletePostNotFound{
			InnerCode: AgentUseCaseCode,
			Details:   server.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &server.APIAgentDeletePostInternalServerError{
			InnerCode: AgentUseCaseCode,
			Details:   server.NewOptString(err.Error()),
		}, nil
	}

	return &server.APIAgentDeletePostNoContent{}, nil
}
