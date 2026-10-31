package adapter

import (
	"context"

	"github.com/gbh007/hgraber-next/internal/domain/serveragentdto"
	"github.com/gbh007/hgraber-next/internal/pkg"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (a *Adapter) Status(ctx context.Context) (serveragentdto.AgentStatus, error) {
	res, err := a.rawClient.APICoreStatusGet(ctx)
	if err != nil {
		return serveragentdto.AgentStatus{}, enrichError(err)
	}

	return serveragentdto.AgentStatus{
		StartAt:   res.StartAt,
		IsOK:      res.Status == agentapi.APICoreStatusGetOKStatusOk,
		IsWarning: res.Status == agentapi.APICoreStatusGetOKStatusWarning,
		IsError:   res.Status == agentapi.APICoreStatusGetOKStatusError,
		Problems: pkg.Map(
			res.Problems,
			func(p agentapi.APICoreStatusGetOKProblemsItem) serveragentdto.AgentStatusProblem {
				return serveragentdto.AgentStatusProblem{
					IsInfo:    p.Type == agentapi.APICoreStatusGetOKProblemsItemTypeInfo,
					IsWarning: p.Type == agentapi.APICoreStatusGetOKProblemsItemTypeWarning,
					IsError:   p.Type == agentapi.APICoreStatusGetOKProblemsItemTypeError,
					Details:   p.Details,
				}
			},
		),
	}, nil
}
