package agenthandlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *AgentHandlersController) APIAgentDeletePost(
	ctx context.Context,
	req *serverapi.APIAgentDeletePostReq,
) error {
	err := c.agentUseCases.DeleteAgent(ctx, req.ID)

	if errors.Is(err, core.ErrAgentNotFound) {
		return apiservercore.APIError{
			Code:      http.StatusNotFound,
			InnerCode: apiservercore.AgentUseCaseCode,
			Details:   err.Error(),
		}
	}

	if err != nil {
		return apiservercore.APIError{
			Code:      http.StatusInternalServerError,
			InnerCode: apiservercore.AgentUseCaseCode,
			Details:   err.Error(),
		}
	}

	return nil
}
