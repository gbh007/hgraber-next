package agent

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"hgnext/internal/entities"
)

func (uc *UseCase) Agent(ctx context.Context, id uuid.UUID) (entities.Agent, error) {
	agent, err := uc.storage.Agent(ctx, id)
	if err != nil {
		return entities.Agent{}, fmt.Errorf("storage: get agent: %w", err)
	}

	return agent, nil
}
