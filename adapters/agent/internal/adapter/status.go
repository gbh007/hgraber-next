package adapter

import (
	"context"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (a *Adapter) Status(ctx context.Context) (agentmodel.AgentStatus, error) {
	res, err := a.rawClient.APICoreStatusGet(ctx)
	if err != nil {
		return agentmodel.AgentStatus{}, enrichError(err)
	}

	return agentmodel.AgentStatus{
		StartAt:   res.StartAt,
		IsOK:      res.Status == agentapi.APICoreStatusGetOKStatusOk,
		IsWarning: res.Status == agentapi.APICoreStatusGetOKStatusWarning,
		IsError:   res.Status == agentapi.APICoreStatusGetOKStatusError,
		Problems: pkg.Map(
			res.Problems,
			func(p agentapi.APICoreStatusGetOKProblemsItem) agentmodel.AgentStatusProblem {
				return agentmodel.AgentStatusProblem{
					IsInfo:    p.Type == agentapi.APICoreStatusGetOKProblemsItemTypeInfo,
					IsWarning: p.Type == agentapi.APICoreStatusGetOKProblemsItemTypeWarning,
					IsError:   p.Type == agentapi.APICoreStatusGetOKProblemsItemTypeError,
					Details:   p.Details,
				}
			},
		),
	}, nil
}
