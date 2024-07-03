package apiserver

import (
	"context"

	"hgnext/internal/controllers/apiserver/internal/server"
	"hgnext/internal/entities"
	"hgnext/internal/pkg"
)

func (c *Controller) APIAgentListPost(ctx context.Context, req *server.APIAgentListPostReq) (server.APIAgentListPostRes, error) {
	agents, err := c.agentUseCases.Agents(ctx, req.CanParse.Value, req.CanExport.Value, req.IncludeStatus.Value)
	if err != nil {
		return &server.APIAgentListPostInternalServerError{
			InnerCode: AgentUseCaseCode,
			Details:   server.NewOptString(err.Error()),
		}, nil
	}

	responseAgents := pkg.Map(agents, func(aws entities.AgentWithStatus) server.APIAgentListPostOKItem {
		status := server.OptAPIAgentListPostOKItemStatus{}

		switch {
		case aws.StatusError != "":
			status = server.NewOptAPIAgentListPostOKItemStatus(server.APIAgentListPostOKItemStatus{
				CheckStatusError: server.NewOptString(aws.StatusError),
				Status:           server.APIAgentListPostOKItemStatusStatusUnknown,
			})

		case aws.IsOffline:
			status = server.NewOptAPIAgentListPostOKItemStatus(server.APIAgentListPostOKItemStatus{
				Status: server.APIAgentListPostOKItemStatusStatusOffline,
			})

		case !aws.Status.StartAt.IsZero():
			t := server.APIAgentListPostOKItemStatusStatusUnknown

			switch {
			case aws.Status.IsOK:
				t = server.APIAgentListPostOKItemStatusStatusOk
			case aws.Status.IsWarning:
				t = server.APIAgentListPostOKItemStatusStatusWarning
			case aws.Status.IsError:
				t = server.APIAgentListPostOKItemStatusStatusError
			}

			status = server.NewOptAPIAgentListPostOKItemStatus(server.APIAgentListPostOKItemStatus{
				StartAt: server.NewOptDateTime(aws.Status.StartAt),
				Problems: pkg.Map(aws.Status.Problems, func(p entities.AgentStatusProblem) server.APIAgentListPostOKItemStatusProblemsItem {
					t := server.APIAgentListPostOKItemStatusProblemsItemTypeError

					switch {
					case p.IsInfo:
						t = server.APIAgentListPostOKItemStatusProblemsItemTypeInfo
					case p.IsWarning:
						t = server.APIAgentListPostOKItemStatusProblemsItemTypeWarning
					}

					return server.APIAgentListPostOKItemStatusProblemsItem{
						Type:    t,
						Details: p.Details,
					}
				}),
				Status: t,
			})
		}

		return server.APIAgentListPostOKItem{
			Status:    status,
			ID:        aws.Agent.ID,
			Name:      aws.Agent.Name,
			Addr:      aws.Agent.Addr,
			CanParse:  aws.Agent.CanParse,
			CanExport: aws.Agent.CanExport,
			Priority:  aws.Agent.Priority,
			CreateAt:  aws.Agent.CreateAt,
		}
	})

	res := server.APIAgentListPostOKApplicationJSON(responseAgents)

	return &res, nil
}
