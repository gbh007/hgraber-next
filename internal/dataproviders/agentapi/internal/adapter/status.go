package adapter

import (
	"context"

	"github.com/gbh007/hgraber-next/internal/domain/serveragentdto"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
	"github.com/samber/lo"
)

func (a *Adapter) Status(ctx context.Context) (serveragentdto.AgentStatus, error) {
	res, err := a.rawClient.APIV2StatusGet(ctx)
	if err != nil {
		return serveragentdto.AgentStatus{}, enrichError(err)
	}

	return serveragentdto.AgentStatus{
		StartAt:   res.StartAt,
		IsOK:      res.Status == agentapi.APIV2StatusGetOKStatusOk,
		IsWarning: res.Status == agentapi.APIV2StatusGetOKStatusWarning,
		IsError:   res.Status == agentapi.APIV2StatusGetOKStatusError,
		Problems: lo.Map(
			res.Problems,
			func(p agentapi.APIV2StatusGetOKProblemsItem, _ int) serveragentdto.AgentStatusProblem {
				return serveragentdto.AgentStatusProblem{
					IsInfo:    p.Type == agentapi.APIV2StatusGetOKProblemsItemTypeInfo,
					IsWarning: p.Type == agentapi.APIV2StatusGetOKProblemsItemTypeWarning,
					IsError:   p.Type == agentapi.APIV2StatusGetOKProblemsItemTypeError,
					Details:   p.Details,
				}
			},
		),
	}, nil
}
