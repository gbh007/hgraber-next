package agent

import (
	"context"
	"errors"
	"fmt"

	"github.com/gbh007/hgraber-next/internal/entities"
	"github.com/gbh007/hgraber-next/internal/pkg"
)

func (uc *UseCase) Agents(ctx context.Context, filter entities.AgentFilter, includeStatus bool) ([]entities.AgentWithStatus, error) {
	agents, err := uc.storage.Agents(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("storage get agents: %w", err)
	}

	res := pkg.Map(agents, func(a entities.Agent) entities.AgentWithStatus {
		return entities.AgentWithStatus{
			Agent: a,
		}
	})

	if !includeStatus {
		return res, nil
	}

	for i, a := range res {
		status, err := uc.agentSystemAdapter.Status(ctx, a.Agent.ID)
		if errors.Is(err, entities.AgentAPIOffline) {
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
