package agenthandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/open_api/serverAPI"
)

func (c *AgentHandlersController) APIAgentNewPost(ctx context.Context, req *serverAPI.APIAgentNewPostReq) (serverAPI.APIAgentNewPostRes, error) {
	err := c.agentUseCases.NewAgent(ctx, core.Agent{
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
			InnerCode: apiservercore.AgentUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	return &serverAPI.APIAgentNewPostNoContent{}, nil
}
