package agent

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"hgnext/internal/entities"
)

func (uc *UseCase) NewAgent(ctx context.Context, agent entities.Agent) error {
	agent.ID = uuid.Must(uuid.NewV7())
	agent.CreateAt = time.Now()

	// Установка нового агента идемпотента, поэтому вначале вызываем ее
	err := uc.agentSystemAdapter.SetAgent(agent)
	if err != nil {
		return fmt.Errorf("agent system setup agent: %w", err)
	}

	err = uc.storage.NewAgent(ctx, agent)
	if err != nil {
		return fmt.Errorf("storage create agent: %w", err)
	}

	return nil
}