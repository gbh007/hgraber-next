package agenthandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *AgentHandlersController) APIAgentUpdatePost(ctx context.Context, req *serverAPI.Agent) (serverAPI.APIAgentUpdatePostRes, error) {
	err := c.agentUseCases.UpdateAgent(ctx, core.Agent{
		ID:            req.ID,
		Name:          req.Name,
		Addr:          req.Addr,
		Token:         req.Token,
		CanParse:      req.CanParse,
		CanParseMulti: req.CanParseMulti,
		CanExport:     req.CanExport,
		HasFS:         req.HasFs,
		Priority:      req.Priority,
		CreateAt:      req.CreatedAt,
	})
	if err != nil {
		return &serverAPI.APIAgentUpdatePostInternalServerError{
			InnerCode: apiservercore.AgentUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIAgentUpdatePostNoContent{}, nil
}
