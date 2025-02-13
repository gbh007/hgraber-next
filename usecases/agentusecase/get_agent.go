package agentusecase

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/core"
)

func (uc *UseCase) Agent(ctx context.Context, id uuid.UUID) (core.Agent, error) {
	agent, err := uc.storage.Agent(ctx, id)
	if err != nil {
		return core.Agent{}, fmt.Errorf("storage: get agent: %w", err)
	}

	return agent, nil
}
