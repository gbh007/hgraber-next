package agenthandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *AgentHandlersController) APIAgentUpdatePost(
	ctx context.Context,
	req *serverapi.APIAgentUpdatePostReq,
) (serverapi.APIAgentUpdatePostRes, error) {
	err := c.agentUseCases.UpdateAgent(ctx, core.Agent{
		ID:            req.ID,
		Name:          req.Name,
		Addr:          req.Addr,
		Token:         req.Token,
		CanParse:      req.CanParse,
		CanParseMulti: req.CanParseMulti,
		CanExport:     req.CanExport,
		HasFS:         req.HasFs,
		HasHProxy:     req.HasHproxy,
		Priority:      req.Priority,
	})
	if err != nil {
		return &serverapi.APIAgentUpdatePostInternalServerError{
			InnerCode: apiservercore.AgentUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIAgentUpdatePostNoContent{}, nil
}
