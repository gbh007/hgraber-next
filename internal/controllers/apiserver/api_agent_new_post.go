package apiserver

import (
	"context"

	"hgnext/internal/entities"
	"hgnext/open_api/serverAPI"
)

func (c *Controller) APIAgentNewPost(ctx context.Context, req *serverAPI.APIAgentNewPostReq) (serverAPI.APIAgentNewPostRes, error) {
	err := c.agentUseCases.NewAgent(ctx, entities.Agent{
		Name:          req.Name,
		Addr:          req.Addr,
		Token:         req.Token,
		CanParse:      req.CanParse.Value,
		CanParseMulti: req.CanParseMulti.Value,
		CanExport:     req.CanExport.Value,
		HasFS:         req.HasFs.Value,
		Priority:      req.Priority.Value,
	})
	if err != nil {
		return &serverAPI.APIAgentNewPostInternalServerError{
			InnerCode: AgentUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIAgentNewPostNoContent{}, nil
}
