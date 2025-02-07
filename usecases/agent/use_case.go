package agent

import (
	"context"
	"log/slog"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/agentmodel"
	"github.com/gbh007/hgraber-next/domain/core"
)

type agentSystemAdapter interface {
	SetAgent(agent core.Agent) error
	DeleteAgent(id uuid.UUID) error
	Status(ctx context.Context, agentID uuid.UUID) (agentmodel.AgentStatus, error)
}

type storage interface {
	Agents(ctx context.Context, filter core.AgentFilter) ([]core.Agent, error)
	Agent(ctx context.Context, id uuid.UUID) (core.Agent, error)
	NewAgent(ctx context.Context, agent core.Agent) error
	UpdateAgent(ctx context.Context, agent core.Agent) error
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
