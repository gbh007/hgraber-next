package agenthandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *AgentHandlersController) APIAgentListPost(
	ctx context.Context,
	req *serverapi.APIAgentListPostReq,
) (serverapi.APIAgentListPostRes, error) {
	agents, err := c.agentUseCases.Agents(ctx, core.AgentFilter{
		CanParse:      req.CanParse.Value,
		CanExport:     req.CanExport.Value,
		CanParseMulti: req.CanParseMulti.Value,
		HasFS:         req.HasFs.Value,
		HasHProxy:     req.HasHproxy.Value,
	}, req.IncludeStatus.Value)
	if err != nil {
		return &serverapi.APIAgentListPostInternalServerError{
			InnerCode: apiservercore.AgentUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	responseAgents := pkg.Map(agents, func(agent agentmodel.AgentWithStatus) serverapi.APIAgentListPostOKItem {
		status := serverapi.OptAPIAgentListPostOKItemStatus{}

		switch {
		case agent.StatusError != "":
			status = serverapi.NewOptAPIAgentListPostOKItemStatus(serverapi.APIAgentListPostOKItemStatus{
				CheckStatusError: serverapi.NewOptString(agent.StatusError),
				Status:           serverapi.APIAgentListPostOKItemStatusStatusUnknown,
			})

		case agent.IsOffline:
			status = serverapi.NewOptAPIAgentListPostOKItemStatus(serverapi.APIAgentListPostOKItemStatus{
				Status: serverapi.APIAgentListPostOKItemStatusStatusOffline,
			})

		case !agent.Status.StartAt.IsZero():
			t := serverapi.APIAgentListPostOKItemStatusStatusUnknown

			switch {
			case agent.Status.IsOK:
				t = serverapi.APIAgentListPostOKItemStatusStatusOk
			case agent.Status.IsWarning:
				t = serverapi.APIAgentListPostOKItemStatusStatusWarning
			case agent.Status.IsError:
				t = serverapi.APIAgentListPostOKItemStatusStatusError
			}

			status = serverapi.NewOptAPIAgentListPostOKItemStatus(serverapi.APIAgentListPostOKItemStatus{
				StartAt: serverapi.NewOptDateTime(agent.Status.StartAt),
				Problems: pkg.Map(
					agent.Status.Problems,
					func(p agentmodel.AgentStatusProblem) serverapi.APIAgentListPostOKItemStatusProblemsItem {
						t := serverapi.APIAgentListPostOKItemStatusProblemsItemTypeError

						switch {
						case p.IsInfo:
							t = serverapi.APIAgentListPostOKItemStatusProblemsItemTypeInfo
						case p.IsWarning:
							t = serverapi.APIAgentListPostOKItemStatusProblemsItemTypeWarning
						}

						return serverapi.APIAgentListPostOKItemStatusProblemsItem{
							Type:    t,
							Details: p.Details,
						}
					},
				),
				Status: t,
			})
		}

		return serverapi.APIAgentListPostOKItem{
			Status: status,
			Info:   apiservercore.ConvertAgentToAPI(agent.Agent),
		}
	})

	res := serverapi.APIAgentListPostOKApplicationJSON(responseAgents)

	return &res, nil
}
