package agentusecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/domain/core"
	"github.com/gbh007/hgraber-next/pkg"
)

func (uc *UseCase) Agents(
	ctx context.Context,
	filter core.AgentFilter,
	includeStatus bool,
) ([]agentmodel.AgentWithStatus, error) {
	agents, err := uc.storage.Agents(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("storage get agents: %w", err)
	}

	res := pkg.Map(agents, func(a core.Agent) agentmodel.AgentWithStatus {
		return agentmodel.AgentWithStatus{
			Agent: a,
		}
	})

	if !includeStatus {
		return res, nil
	}

	for i, a := range res {
		status, err := uc.agentSystemAdapter.Status(ctx, a.Agent.ID)
		if errors.Is(err, agentmodel.AgentAPIOffline) {
			res[i].IsOffline = true

			continue
		}

		if err != nil {
			res[i].StatusError = err.Error()

			continue
		}

		res[i].Status = status
	}

	return res, nil
}
