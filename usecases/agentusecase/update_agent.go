package agentusecase

import (
	"context"
	"fmt"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (uc *UseCase) UpdateAgent(ctx context.Context, agent core.Agent) error {
	// Установка нового агента идемпотента, поэтому вначале вызываем ее
	err := uc.agentSystemAdapter.SetAgent(agent)
	if err != nil {
		return fmt.Errorf("agent system: setup agent: %w", err)
	}

	err = uc.storage.UpdateAgent(ctx, agent)
	if err != nil {
		return fmt.Errorf("storage: update agent: %w", err)
	}

	return nil
}
