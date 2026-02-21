package agenthandlers

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *AgentHandlersController) APIAgentNewPost(
	ctx context.Context,
	req *serverapi.APIAgentNewPostReq,
) error {
	err := c.agentUseCases.NewAgent(ctx, core.Agent{
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
	if err != nil {
		return apiservercore.APIError{
			Code:      http.StatusInternalServerError,
			InnerCode: apiservercore.AgentUseCaseCode,
			Details:   err.Error(),
		}
	}

	return nil
}
