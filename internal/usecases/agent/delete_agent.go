package agent

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"

	"hgnext/internal/entities"
)

func (uc *UseCase) DeleteAgent(ctx context.Context, id uuid.UUID) error {
	err := uc.agentSystemAdapter.DeleteAgent(id)
	if err != nil && !errors.Is(err, entities.AgentNotFoundError) {
		return fmt.Errorf("agent system delete agent: %w", err)
	}

	err = uc.storage.DeleteAgent(ctx, id)
	if err != nil {
		return fmt.Errorf("storage delete agent: %w", err)
	}

	return nil
}
