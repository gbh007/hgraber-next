package agenthandlers

import (
	"context"
	"errors"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
)

func (c *AgentHandlersController) APIAgentDeletePost(ctx context.Context, req *serverapi.APIAgentDeletePostReq) (serverapi.APIAgentDeletePostRes, error) {
	err := c.agentUseCases.DeleteAgent(ctx, req.ID)

	if errors.Is(err, core.AgentNotFoundError) {
		return &serverapi.APIAgentDeletePostNotFound{
			InnerCode: apiservercore.AgentUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &serverapi.APIAgentDeletePostInternalServerError{
			InnerCode: apiservercore.AgentUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	return &serverapi.APIAgentDeletePostNoContent{}, nil
}
