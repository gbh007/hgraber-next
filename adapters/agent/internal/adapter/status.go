package adapter

import (
	"context"
	"fmt"

	"github.com/gbh007/hgraber-next/entities"
	"github.com/gbh007/hgraber-next/open_api/agentAPI"
	"github.com/gbh007/hgraber-next/pkg"
)

func (a *Adapter) Status(ctx context.Context) (entities.AgentStatus, error) {
	res, err := a.rawClient.APICoreStatusGet(ctx)
	if err != nil {
		return entities.AgentStatus{}, err
	}

	switch typedRes := res.(type) {
	case *agentAPI.APICoreStatusGetOK:
		return entities.AgentStatus{
			StartAt:   typedRes.StartAt,
			IsOK:      typedRes.Status == agentAPI.APICoreStatusGetOKStatusOk,
			IsWarning: typedRes.Status == agentAPI.APICoreStatusGetOKStatusWarning,
			IsError:   typedRes.Status == agentAPI.APICoreStatusGetOKStatusError,
			Problems: pkg.Map(typedRes.Problems, func(p agentAPI.APICoreStatusGetOKProblemsItem) entities.AgentStatusProblem {
				return entities.AgentStatusProblem{
					IsInfo:    p.Type == agentAPI.APICoreStatusGetOKProblemsItemTypeInfo,
					IsWarning: p.Type == agentAPI.APICoreStatusGetOKProblemsItemTypeWarning,
					IsError:   p.Type == agentAPI.APICoreStatusGetOKProblemsItemTypeError,
					Details:   p.Details,
				}
			}),
		}, nil

	case *agentAPI.APICoreStatusGetBadRequest:
		return entities.AgentStatus{}, fmt.Errorf("%w: %s", entities.AgentAPIBadRequest, typedRes.Details.Value)

	case *agentAPI.APICoreStatusGetUnauthorized:
		return entities.AgentStatus{}, fmt.Errorf("%w: %s", entities.AgentAPIUnauthorized, typedRes.Details.Value)

	case *agentAPI.APICoreStatusGetForbidden:
		return entities.AgentStatus{}, fmt.Errorf("%w: %s", entities.AgentAPIForbidden, typedRes.Details.Value)

	case *agentAPI.APICoreStatusGetInternalServerError:
		return entities.AgentStatus{}, fmt.Errorf("%w: %s", entities.AgentAPIInternalError, typedRes.Details.Value)

	default:
		return entities.AgentStatus{}, entities.AgentAPIUnknownResponse
	}
}
