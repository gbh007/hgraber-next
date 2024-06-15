package adapter

import (
	"context"
	"fmt"

	"hgnext/internal/adapters/agent/internal/client"
	"hgnext/internal/entities"
	"hgnext/internal/pkg"
)

func (a *Adapter) Status(ctx context.Context) (entities.AgentStatus, error) {
	res, err := a.rawClient.APICoreStatusGet(ctx)
	if err != nil {
		return entities.AgentStatus{}, err
	}

	switch typedRes := res.(type) {
	case *client.APICoreStatusGetOK:
		return entities.AgentStatus{
			StartAt:   typedRes.StartAt,
			IsOK:      typedRes.Status == client.APICoreStatusGetOKStatusOk,
			IsWarning: typedRes.Status == client.APICoreStatusGetOKStatusWarning,
			IsError:   typedRes.Status == client.APICoreStatusGetOKStatusError,
			Problems: pkg.Map(typedRes.Problems, func(p client.APICoreStatusGetOKProblemsItem) entities.AgentStatusProblem {
				return entities.AgentStatusProblem{
					IsInfo:    p.Type == client.APICoreStatusGetOKProblemsItemTypeInfo,
					IsWarning: p.Type == client.APICoreStatusGetOKProblemsItemTypeWarning,
					IsError:   p.Type == client.APICoreStatusGetOKProblemsItemTypeError,
				}
			}),
		}, nil

	case *client.APICoreStatusGetBadRequest:
		return entities.AgentStatus{}, fmt.Errorf("%w: %s", entities.AgentAPIBadRequest, typedRes.Details.Value)

	case *client.APICoreStatusGetUnauthorized:
		return entities.AgentStatus{}, fmt.Errorf("%w: %s", entities.AgentAPIUnauthorized, typedRes.Details.Value)

	case *client.APICoreStatusGetForbidden:
		return entities.AgentStatus{}, fmt.Errorf("%w: %s", entities.AgentAPIForbidden, typedRes.Details.Value)

	case *client.APICoreStatusGetInternalServerError:
		return entities.AgentStatus{}, fmt.Errorf("%w: %s", entities.AgentAPIInternalError, typedRes.Details.Value)

	default:
		return entities.AgentStatus{}, entities.AgentAPIUnknownResponse
	}
}
