package apiserver

import (
	"context"

	"hgnext/internal/controllers/apiserver/internal/server"
	"hgnext/internal/entities"
)

func (c *Controller) APIAgentNewPost(ctx context.Context, req *server.APIAgentNewPostReq) (server.APIAgentNewPostRes, error) {
	err := c.agentUseCases.NewAgent(ctx, entities.Agent{
		Name:      req.Name,
		Addr:      req.Addr,
		Token:     req.Token,
		CanParse:  req.CanParse.Value,
		CanExport: req.CanExport.Value,
		Priority:  req.Priority.Value,
	})
	if err != nil {
		return &server.APIAgentNewPostInternalServerError{
			InnerCode: AgentUseCaseCode,
			Details:   server.NewOptString(err.Error()),
		}, nil
	}

	return &server.APIAgentNewPostNoContent{}, nil
}
