package apiserver

import (
	"context"

	"hgnext/internal/entities"
	"hgnext/open_api/serverAPI"
)

func (c *Controller) APIAgentUpdatePost(ctx context.Context, req *serverAPI.Agent) (serverAPI.APIAgentUpdatePostRes, error) {
	err := c.agentUseCases.UpdateAgent(ctx, entities.Agent{
		ID:            req.ID,
		Name:          req.Name,
		Addr:          req.Addr,
		Token:         req.Token,
		CanParse:      req.CanParse,
		CanParseMulti: req.CanParseMulti,
		CanExport:     req.CanExport,
		Priority:      req.Priority,
		CreateAt:      req.CreatedAt,
	})
	if err != nil {
		return &serverAPI.APIAgentUpdatePostInternalServerError{
			InnerCode: AgentUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIAgentUpdatePostNoContent{}, nil
}
