package agenthandlers

import (
	"context"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *AgentHandlersController) APIAgentUpdatePost(
	ctx context.Context,
	req *serverapi.APIAgentUpdatePostReq,
) error {
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
		return apiservercore.APIError{
			Code:      http.StatusInternalServerError,
			InnerCode: apiservercore.AgentUseCaseCode,
			Details:   err.Error(),
		}
	}

	return nil
}
