package adapter

import (
	"context"
	"fmt"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (a *Adapter) Status(ctx context.Context) (core.AgentStatus, error) {
	res, err := a.rawClient.APICoreStatusGet(ctx)
	if err != nil {
		return core.AgentStatus{}, err
	}

	switch typedRes := res.(type) {
	case *agentapi.APICoreStatusGetOK:
		return core.AgentStatus{
			StartAt:   typedRes.StartAt,
			IsOK:      typedRes.Status == agentapi.APICoreStatusGetOKStatusOk,
			IsWarning: typedRes.Status == agentapi.APICoreStatusGetOKStatusWarning,
			IsError:   typedRes.Status == agentapi.APICoreStatusGetOKStatusError,
			Problems: pkg.Map(typedRes.Problems, func(p agentapi.APICoreStatusGetOKProblemsItem) core.AgentStatusProblem {
				return core.AgentStatusProblem{
					IsInfo:    p.Type == agentapi.APICoreStatusGetOKProblemsItemTypeInfo,
					IsWarning: p.Type == agentapi.APICoreStatusGetOKProblemsItemTypeWarning,
					IsError:   p.Type == agentapi.APICoreStatusGetOKProblemsItemTypeError,
					Details:   p.Details,
				}
			}),
		}, nil

	case *agentapi.APICoreStatusGetBadRequest:
		return core.AgentStatus{}, fmt.Errorf("%w: %s", agentmodel.AgentAPIBadRequest, typedRes.Details.Value)

	case *agentapi.APICoreStatusGetUnauthorized:
		return core.AgentStatus{}, fmt.Errorf("%w: %s", agentmodel.AgentAPIUnauthorized, typedRes.Details.Value)

	case *agentapi.APICoreStatusGetForbidden:
		return core.AgentStatus{}, fmt.Errorf("%w: %s", agentmodel.AgentAPIForbidden, typedRes.Details.Value)

	case *agentapi.APICoreStatusGetInternalServerError:
		return core.AgentStatus{}, fmt.Errorf("%w: %s", agentmodel.AgentAPIInternalError, typedRes.Details.Value)

	default:
		return core.AgentStatus{}, agentmodel.AgentAPIUnknownResponse
	}
}
