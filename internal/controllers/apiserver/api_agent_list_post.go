package apiserver

import (
	"context"

	"hgnext/internal/entities"
	"hgnext/internal/pkg"
	"hgnext/open_api/serverAPI"
)

func (c *Controller) APIAgentListPost(ctx context.Context, req *serverAPI.APIAgentListPostReq) (serverAPI.APIAgentListPostRes, error) {
	agents, err := c.agentUseCases.Agents(ctx, entities.AgentFilter{
		CanParse:      req.CanParse.Value,
		CanExport:     req.CanExport.Value,
		CanParseMulti: req.CanParseMulti.Value,
		HasFS:         req.HasFs.Value,
	}, req.IncludeStatus.Value)
	if err != nil {
		return &serverAPI.APIAgentListPostInternalServerError{
			InnerCode: AgentUseCaseCode,
			Details:   serverAPI.NewOptString(err.Error()),
		}, nil
	}

	responseAgents := pkg.Map(agents, func(aws entities.AgentWithStatus) serverAPI.APIAgentListPostOKItem {
		status := serverAPI.OptAPIAgentListPostOKItemStatus{}

		switch {
		case aws.StatusError != "":
			status = serverAPI.NewOptAPIAgentListPostOKItemStatus(serverAPI.APIAgentListPostOKItemStatus{
				CheckStatusError: serverAPI.NewOptString(aws.StatusError),
				Status:           serverAPI.APIAgentListPostOKItemStatusStatusUnknown,
			})

		case aws.IsOffline:
			status = serverAPI.NewOptAPIAgentListPostOKItemStatus(serverAPI.APIAgentListPostOKItemStatus{
				Status: serverAPI.APIAgentListPostOKItemStatusStatusOffline,
			})

		case !aws.Status.StartAt.IsZero():
			t := serverAPI.APIAgentListPostOKItemStatusStatusUnknown

			switch {
			case aws.Status.IsOK:
				t = serverAPI.APIAgentListPostOKItemStatusStatusOk
			case aws.Status.IsWarning:
				t = serverAPI.APIAgentListPostOKItemStatusStatusWarning
			case aws.Status.IsError:
				t = serverAPI.APIAgentListPostOKItemStatusStatusError
			}

			status = serverAPI.NewOptAPIAgentListPostOKItemStatus(serverAPI.APIAgentListPostOKItemStatus{
				StartAt: serverAPI.NewOptDateTime(aws.Status.StartAt),
				Problems: pkg.Map(aws.Status.Problems, func(p entities.AgentStatusProblem) serverAPI.APIAgentListPostOKItemStatusProblemsItem {
					t := serverAPI.APIAgentListPostOKItemStatusProblemsItemTypeError

					switch {
					case p.IsInfo:
						t = serverAPI.APIAgentListPostOKItemStatusProblemsItemTypeInfo
					case p.IsWarning:
						t = serverAPI.APIAgentListPostOKItemStatusProblemsItemTypeWarning
					}

					return serverAPI.APIAgentListPostOKItemStatusProblemsItem{
						Type:    t,
						Details: p.Details,
					}
				}),
				Status: t,
			})
		}

		return serverAPI.APIAgentListPostOKItem{
			Status: status,
			Info:   convertAgentToAPI(aws.Agent),
		}
	})

	res := serverAPI.APIAgentListPostOKApplicationJSON(responseAgents)

	return &res, nil
}
