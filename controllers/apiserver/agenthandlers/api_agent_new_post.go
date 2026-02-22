package agenthandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *AgentHandlersController) APIAgentNewPost(
	ctx context.Context,
	req *serverapi.APIAgentNewPostReq,
) error {
	return c.agentUseCases.NewAgent(ctx, core.Agent{
		Name:          req.Name,
		Addr:          req.Addr,
		Token:         req.Token,
		CanParse:      req.CanParse.Value,
		CanParseMulti: req.CanParseMulti.Value,
		CanExport:     req.CanExport.Value,
		HasFS:         req.HasFs.Value,
		HasHProxy:     req.HasHproxy.Value,
		Priority:      req.Priority.Value,
	})
}
