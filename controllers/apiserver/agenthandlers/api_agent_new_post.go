package agenthandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *AgentHandlersController) APIAgentNewPost(ctx context.Context, req *serverapi.APIAgentNewPostReq) (serverapi.APIAgentNewPostRes, error) {
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
		return &serverapi.APIAgentNewPostInternalServerError{
			InnerCode: apiservercore.AgentUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIAgentNewPostNoContent{}, nil
}
