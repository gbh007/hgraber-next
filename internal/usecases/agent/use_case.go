package agent

import (
	"context"
	"log/slog"

	"github.com/google/uuid"

	"hgnext/internal/entities"
)

type agentSystemAdapter interface {
	SetAgent(agent entities.Agent) error
	DeleteAgent(id uuid.UUID) error
	Status(ctx context.Context, agentID uuid.UUID) (entities.AgentStatus, error)
}

type storage interface {
	Agents(ctx context.Context, canParse, canExport bool) ([]entities.Agent, error)
	NewAgent(ctx context.Context, agent entities.Agent) error
	UpdateAgent(ctx context.Context, agent entities.Agent) error
	DeleteAgent(ctx context.Context, id uuid.UUID) error
}

type UseCase struct {
	logger *slog.Logger

	agentSystemAdapter agentSystemAdapter
	storage            storage
}

func New(
	logger *slog.Logger,
	agentSystemAdapter agentSystemAdapter,
	storage storage,
) *UseCase {
	return &UseCase{
		logger:             logger,
		agentSystemAdapter: agentSystemAdapter,
		storage:            storage,
	}
}
