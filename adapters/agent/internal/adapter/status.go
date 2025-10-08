package adapter

import (
	"context"
	"fmt"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (a *Adapter) Status(ctx context.Context) (agentmodel.AgentStatus, error) {
	res, err := a.rawClient.APICoreStatusGet(ctx)
	if err != nil {
		return agentmodel.AgentStatus{}, fmt.Errorf("request: %w", err)
	}

	switch typedRes := res.(type) {
	case *agentapi.APICoreStatusGetOK:
		return agentmodel.AgentStatus{
			StartAt:   typedRes.StartAt,
			IsOK:      typedRes.Status == agentapi.APICoreStatusGetOKStatusOk,
			IsWarning: typedRes.Status == agentapi.APICoreStatusGetOKStatusWarning,
			IsError:   typedRes.Status == agentapi.APICoreStatusGetOKStatusError,
			Problems: pkg.Map(
				typedRes.Problems,
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

	case *agentapi.APICoreStatusGetBadRequest:
		return agentmodel.AgentStatus{}, fmt.Errorf("%w: %s", agentmodel.ErrAgentAPIBadRequest, typedRes.Details.Value)

	case *agentapi.APICoreStatusGetUnauthorized:
		return agentmodel.AgentStatus{}, fmt.Errorf(
			"%w: %s",
			agentmodel.ErrAgentAPIUnauthorized,
			typedRes.Details.Value,
		)

	case *agentapi.APICoreStatusGetForbidden:
		return agentmodel.AgentStatus{}, fmt.Errorf("%w: %s", agentmodel.ErrAgentAPIForbidden, typedRes.Details.Value)

	case *agentapi.APICoreStatusGetInternalServerError:
		return agentmodel.AgentStatus{}, fmt.Errorf(
			"%w: %s",
			agentmodel.ErrAgentAPIInternalError,
			typedRes.Details.Value,
		)

	default:
		return agentmodel.AgentStatus{}, agentmodel.ErrAgentAPIUnknownResponse
	}
}
