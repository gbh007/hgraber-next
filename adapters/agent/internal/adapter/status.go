package adapter

import (
	"context"
	"fmt"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/open_api/agentAPI"
	"github.com/gbh007/hgraber-next/pkg"
)

func (a *Adapter) Status(ctx context.Context) (core.AgentStatus, error) {
	res, err := a.rawClient.APICoreStatusGet(ctx)
	if err != nil {
		return core.AgentStatus{}, err
	}

	switch typedRes := res.(type) {
	case *agentAPI.APICoreStatusGetOK:
		return core.AgentStatus{
			StartAt:   typedRes.StartAt,
			IsOK:      typedRes.Status == agentAPI.APICoreStatusGetOKStatusOk,
			IsWarning: typedRes.Status == agentAPI.APICoreStatusGetOKStatusWarning,
			IsError:   typedRes.Status == agentAPI.APICoreStatusGetOKStatusError,
			Problems: pkg.Map(typedRes.Problems, func(p agentAPI.APICoreStatusGetOKProblemsItem) core.AgentStatusProblem {
				return core.AgentStatusProblem{
					IsInfo:    p.Type == agentAPI.APICoreStatusGetOKProblemsItemTypeInfo,
					IsWarning: p.Type == agentAPI.APICoreStatusGetOKProblemsItemTypeWarning,
					IsError:   p.Type == agentAPI.APICoreStatusGetOKProblemsItemTypeError,
					Details:   p.Details,
				}
			}),
		}, nil

	case *agentAPI.APICoreStatusGetBadRequest:
		return core.AgentStatus{}, fmt.Errorf("%w: %s", agentmodel.AgentAPIBadRequest, typedRes.Details.Value)

	case *agentAPI.APICoreStatusGetUnauthorized:
		return core.AgentStatus{}, fmt.Errorf("%w: %s", agentmodel.AgentAPIUnauthorized, typedRes.Details.Value)

	case *agentAPI.APICoreStatusGetForbidden:
		return core.AgentStatus{}, fmt.Errorf("%w: %s", agentmodel.AgentAPIForbidden, typedRes.Details.Value)

	case *agentAPI.APICoreStatusGetInternalServerError:
		return core.AgentStatus{}, fmt.Errorf("%w: %s", agentmodel.AgentAPIInternalError, typedRes.Details.Value)

	default:
		return core.AgentStatus{}, agentmodel.AgentAPIUnknownResponse
	}
}
