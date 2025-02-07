package agenthandlers

import (
	"context"

	"github.com/gbh007/hgraber-next/controllers/apiserver/apiservercore"
	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *AgentHandlersController) APIAgentListPost(ctx context.Context, req *serverapi.APIAgentListPostReq) (serverapi.APIAgentListPostRes, error) {
	agents, err := c.agentUseCases.Agents(ctx, core.AgentFilter{
		CanParse:      req.CanParse.Value,
		CanExport:     req.CanExport.Value,
		CanParseMulti: req.CanParseMulti.Value,
		HasFS:         req.HasFs.Value,
	}, req.IncludeStatus.Value)
	if err != nil {
		return &serverapi.APIAgentListPostInternalServerError{
			InnerCode: apiservercore.AgentUseCaseCode,
			Details:   serverapi.NewOptString(err.Error()),
		}, nil
	}

	responseAgents := pkg.Map(agents, func(aws agentmodel.AgentWithStatus) serverapi.APIAgentListPostOKItem {
		status := serverapi.OptAPIAgentListPostOKItemStatus{}

		switch {
		case aws.StatusError != "":
			status = serverapi.NewOptAPIAgentListPostOKItemStatus(serverapi.APIAgentListPostOKItemStatus{
				CheckStatusError: serverapi.NewOptString(aws.StatusError),
				Status:           serverapi.APIAgentListPostOKItemStatusStatusUnknown,
			})

		case aws.IsOffline:
			status = serverapi.NewOptAPIAgentListPostOKItemStatus(serverapi.APIAgentListPostOKItemStatus{
				Status: serverapi.APIAgentListPostOKItemStatusStatusOffline,
			})

		case !aws.Status.StartAt.IsZero():
			t := serverapi.APIAgentListPostOKItemStatusStatusUnknown

			switch {
			case aws.Status.IsOK:
				t = serverapi.APIAgentListPostOKItemStatusStatusOk
			case aws.Status.IsWarning:
				t = serverapi.APIAgentListPostOKItemStatusStatusWarning
			case aws.Status.IsError:
				t = serverapi.APIAgentListPostOKItemStatusStatusError
			}

			status = serverapi.NewOptAPIAgentListPostOKItemStatus(serverapi.APIAgentListPostOKItemStatus{
				StartAt: serverapi.NewOptDateTime(aws.Status.StartAt),
				Problems: pkg.Map(aws.Status.Problems, func(p agentmodel.AgentStatusProblem) serverapi.APIAgentListPostOKItemStatusProblemsItem {
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
				}),
				Status: t,
			})
		}

		return serverapi.APIAgentListPostOKItem{
			Status: status,
			Info:   apiservercore.ConvertAgentToAPI(aws.Agent),
		}
	})

	res := serverapi.APIAgentListPostOKApplicationJSON(responseAgents)

	return &res, nil
}
